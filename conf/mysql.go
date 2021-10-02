package conf

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// DB 全局MySQL操作对象
var DB *gorm.DB

// initMySQL 初始化MySQL链接
func initMySQL() {
	Log.Info("mysql database is connecting···")
	// 配置连接参数
	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		Config.MySQL.User,
		Config.MySQL.Passwd,
		Config.MySQL.Host,
		Config.MySQL.Port,
		Config.MySQL.Name,
		Config.MySQL.CharSet,
	)
	// 预定义错误值
	var err error
	// 执行数据库连接
	DB, err = gorm.Open(
		mysql.Open(DSN),
		&gorm.Config{
			Logger: newDbLog(
				logger.Info,
				1,
			), // 自定义数据库日志记录器
		},
	)
	// 判断是否连接成功
	if err != nil {
		Log.Error("mysql database connect error: %s", err.Error())
		panic(err.Error())
	}
	// 数据库连接成功
	Log.Info("mysql database connect success")
}

// dbLogger 自定义数据库日志管理器
type dbLogger struct {
	LogLevel      logger.LogLevel // 日志级别
	SlowThreshold int             // 慢日志阈值
}

// newDbLog 创建数据库日志记录器
func newDbLog(level logger.LogLevel, slow int) logger.Interface {
	return &dbLogger{
		LogLevel:      level,
		SlowThreshold: slow,
	}
}

// LogMode 数据库日志级别配置
func (l *dbLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info 数据库信息级别日志
func (l dbLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		format := fmt.Sprintf("[exec:%s] %s", utils.FileWithLineNum(), msg)
		Log.Info(format, data...)
	}
}

// Warn 数据库警告级别日志
func (l dbLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		format := fmt.Sprintf("[exec:%s] %s", utils.FileWithLineNum(), msg)
		Log.Info(format, data...)
	}
}

// Error 数据库错误级别日志
func (l dbLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		format := fmt.Sprintf("[exec:%s] %s", utils.FileWithLineNum(), msg)
		Log.Info(format, data...)
	}
}

// Trace 数据库查询日志
func (l dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// 级别为1时不记录日志
	if l.LogLevel <= logger.Silent {
		return
	}
	// 计算SQL执行耗时
	elapsed := time.Since(begin)
	// 区别打印
	switch {
	// 有错误，当前记录器级别大于等于错误级别，不是查询为空的记录，
	case err != nil && l.LogLevel >= logger.Error && !errors.Is(err, logger.ErrRecordNotFound):
		// 获取SQL及影响行数
		sql, rows := fc()
		Log.Error(
			"[exec:%s] [%.3fms] [rows:%d] [err: %s] SQL: %s",
			utils.FileWithLineNum(),
			float64(elapsed.Nanoseconds())/1e6,
			rows,
			err.Error(),
			sql,
		)
	// 慢日志打印
	case int(elapsed.Seconds()) > l.SlowThreshold && l.SlowThreshold >= 0 && l.LogLevel >= logger.Warn:
		// 获取SQL及影响行数
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		Log.Warn(
			"[exec:%s] [%.3fms] [rows:%d] [warn: %s] SQL: %s",
			utils.FileWithLineNum(),
			float64(elapsed.Nanoseconds())/1e6,
			rows,
			slowLog,
			sql,
		)
	// 普通级别日志打印
	case l.LogLevel == logger.Info:
		// 获取SQL及影响行数
		sql, rows := fc()
		Log.Info(
			"[exec:%s] [%.3fms] [rows:%d] SQL: %s",
			utils.FileWithLineNum(),
			float64(elapsed.Nanoseconds())/1e6,
			rows,
			sql,
		)
	}
}
