package models_pg

import (
	"github.com/feriyusuf/go-sign/app/helpers"
	"github.com/feriyusuf/go-sign/app/models_pg/commons"
)

const Superadmin = "SUPERADMIN"
const SuperAdminPassword = "pass123"
const SuperAdminUsername = "superadmin@go-sign.dev"
const PublicUser = "PUBLIC_USER"

func PostgresAutoPopulate() {
	var superAdminRole commons.ComRole
	var superAdminUser commons.ComUser
	var publicUserRole commons.ComRole
	var superAdminRoleUser commons.ComRoleUser

	// Create Role as Superadmin if it doesn't exist
	PGDB.Where("name = ?", Superadmin).Find(&superAdminRole)
	if superAdminRole.ID == 0 {
		superAdminRole = commons.ComRole{
			Name:      Superadmin,
			CreatedBy: 1,
		}

		PGDB.Create(&superAdminRole)
	}

	// Assign all access to super admin (if any)

	// Create User as Superadmin
	PGDB.Where("username = ?", SuperAdminUsername).Find(&superAdminUser)
	if superAdminUser.ID == 0 {
		superAdminUser = commons.ComUser{
			Name:      Superadmin,
			Username:  SuperAdminUsername,
			Password:  helpers.GenerateHashPassword([]byte(SuperAdminPassword)),
			CreatedBy: 1,
		}

		PGDB.Create(&superAdminUser)
	}

	// Assign Role to Superadmin Role
	PGDB.Where("user_id = ? AND role_id = ?", superAdminUser.ID, superAdminRole.ID).Find(&superAdminRoleUser)
	if superAdminRoleUser.RoleId == 0 {
		superAdminRoleUser = commons.ComRoleUser{
			UserId:    superAdminUser.ID,
			RoleId:    superAdminRole.ID,
			CreatedBy: superAdminUser.ID,
		}

		PGDB.Create(&superAdminRoleUser)
	}

	// Create Role as PUBLIC USER
	PGDB.Where("name = ?", PublicUser).Find(&publicUserRole)
	if publicUserRole.ID == 0 {
		publicUserRole = commons.ComRole{
			Name:      PublicUser,
			CreatedBy: superAdminUser.ID,
		}

		PGDB.Create(&publicUserRole)
	}
}
