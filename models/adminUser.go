package models

import (
	"cinema-admin/db"
	"cinema-admin/utils"
	"errors"
	"fmt"
	"time"
)

// AdminUser defines how an admin user is represented in database
type AdminUser struct {
	ID        string `json:"id,omitempty"`
	Email     string `gorm:"not null;unique" json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string
	Username  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	LastLogin *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type enumRole struct {
	ROOT            string
	COMMUNITY       string
	BUSINESSMANAGER string
}

// Role ...
var Role = enumRole{
	ROOT:            "root",
	COMMUNITY:       "community",
	BUSINESSMANAGER: "business_manager",
}

// BeforeSave ...
func (adminUser *AdminUser) BeforeSave() error {
	if adminUser.Email == "" || len(adminUser.Password) <= 6 {
		return errors.New("Email or password is empty")
	}

	if len(adminUser.Password) < 35 {
		var err error
		adminUser.Password, err = utils.Generate(adminUser.Password)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateAdminUser ...
func CreateAdminUser(email string, password string) (AdminUser, error) {
	if email == "" || len(password) < 6 {
		fmt.Println(email)
		fmt.Println(password)
		return AdminUser{}, errors.New("Data or data type is invalid")
	}

	var err error
	var adminUser AdminUser
	adminUser.Email = email
	adminUser.Password, err = utils.Generate(password)
	adminUser.Username = "Admin"
	if err != nil {
		return AdminUser{}, err
	}
	err = adminUser.create()
	if err != nil {
		return AdminUser{}, errors.New("Can't create new resource to database")
	}
	return adminUser, nil
}

// GetAdminUserByEmail ...
func GetAdminUserByEmail(email string) (AdminUser, error) {
	dbConn := db.GetDB()

	adminUser := AdminUser{}
	res := dbConn.Where("email = ?", email).First(&adminUser)
	if res.Error != nil {
		return adminUser, errors.New("Data or data type is invalid")
	}
	return adminUser, nil
}

// GetAdminUserByID ...
func GetAdminUserByID(adminUserID string) (AdminUser, error) {
	dbConn := db.GetDB()

	adminUser := AdminUser{}
	res := dbConn.Where("id = ?", adminUserID).Find(&adminUser)
	if res.Error != nil {
		return adminUser, errors.New("Data or data type is invalid")
	}
	return adminUser, nil
}

func (adminUser *AdminUser) create() error {
	dbConn := db.GetDB()
	return dbConn.Create(&adminUser).Error
}

// DisplayName ...
func (adminUser AdminUser) DisplayName() string {
	return adminUser.Username
}

// GetAdminByEmail ...
func GetAdminByEmail(email string) (AdminUser, error) {
	dbConn := db.GetDB()

	admin := AdminUser{}
	res := dbConn.Where("email = ?", email).First(&admin)
	if res.Error != nil {
		return admin, errors.New("Data or data type is invalid")
	}
	return admin, nil
}

// GetAdminByID ...
func GetAdminByID(adminID string) (AdminUser, error) {
	dbConn := db.GetDB()

	admin := AdminUser{}
	res := dbConn.Where("id = ?", adminID).Find(&admin)
	if res.Error != nil {
		return admin, errors.New("Data or data type is invalid")
	}
	return admin, nil
}
