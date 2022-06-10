package models_pg

import (
	"fmt"
	"github.com/feriyusuf/go-sign/app/helpers"
	"github.com/feriyusuf/go-sign/app/models_pg/commons"
	"github.com/feriyusuf/go-sign/app/script/migration"
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

	// Populate default menus
	tx := PGDB.Exec(migration.SqlDefaultMenusInit)

	if tx.Error != nil {
		panic("Something went wrong while creating menu(s)")
	}

	if tx.Error == nil {
		AssignMenus(0, superAdminRole.ID)
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

func AssignMenus(parent uint, roleId uint) {
	query := `parent_id is null`

	if parent != 0 {
		query = fmt.Sprintf("parent_id = %v", parent)
	}

	rows, err := PGDB.Debug().Model(&commons.ComMenu{}).Where(query).Rows()
	defer rows.Close()

	if err != nil {
		panic("Something went wrong while assign menu(s)")
	}

	for rows.Next() {
		var menu commons.ComMenu
		PGDB.ScanRows(rows, &menu)

		result := PGDB.Debug().Create(&commons.ComMenuRole{
			RoleId:   roleId,
			ComMenu:  menu,
			IsActive: true,
			Read:     true,
			Write:    true,
		})

		if result.Error != nil {
			panic("Something went wrong while assign menu(s)")
		}

		AssignMenus(menu.ID, roleId)
	}
}
