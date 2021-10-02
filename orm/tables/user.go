package tables

// OrmUser 数据库用户表
type OrmUser struct {
	ID int32 // 用户表主键
}

// TableName 数据表映射
func (OrmUser) TableName() string {
	return "user"
}
