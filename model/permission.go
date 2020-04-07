package model

/*
 * 管理员权限表格
 */
type Permission struct {
	PermissionId int64  `xorm:"pk autoincr" json:"id"`
	Level        string `xorm:"varchar(32)" json:"level"`
	Name         string `xorm:"varchar(32)" json:"name"`
}
