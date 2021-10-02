package tables

import "time"

// OrmUserGroup 数据库用户分组表
type OrmUserGroup struct {
	ID         int32     // 用户分组主键
	Name       string    // 用户分组名称
	IsSuper    bool      // 是否是超级分组（超级分组不允许删除）
	CreateTime time.Time // 用户分组创建时间
	UpdateTime time.Time // 用户分组更新时间
}

// TableName 数据表映射
func (OrmUserGroup) TableName() string {
	return "user_group"
}
