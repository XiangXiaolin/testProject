package model

/*
 * 管理员的权限明细表
 */
type AdminPermission struct {
	Admin      *Admin      `xorm:"extends"`
	Permission *Permission `xorm:"extends"`
}
