package jobs

import (
	"github.com/spurtcms/auth"
	role "github.com/spurtcms/team-roles"
	"gorm.io/gorm"
)

type Type string

const ( //for permission check
	Postgres Type = "postgres"
	Mysql    Type = "mysql"
)

type Config struct {
	DB               *gorm.DB
	AuthEnable       bool
	PermissionEnable bool
	DataBaseType     Type
	Auth             *auth.Auth
	Permissions      *role.PermissionConfig
}

type Jobs struct {
	AuthEnable       bool
	PermissionEnable bool
	AuthFlg          bool
	PermissionFlg    bool
	DB               *gorm.DB
	Auth             *auth.Auth
	Permissions      *role.PermissionConfig
}
