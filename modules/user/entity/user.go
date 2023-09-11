package entity

import (
	"app/common"
	"app/component/tokenprovider"
	"errors"
	"time"
)

type User struct {
	Id           string    `json:"id"`
	FullName     string    `json:"fullName"`
	EmployeeCode int       `json:"employeeCode"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	Password     string    `json:"password"`
	Salt         string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Account struct {
	AccressToken *tokenprovider.Token `json:"accessToken"`
	RefreshToken *tokenprovider.Token `json:"refreshToken"`
}

func NewAccount(accessToken, refreshToken *tokenprovider.Token) *Account {
	return &Account{
		AccressToken: accessToken,
		RefreshToken: refreshToken,
	}
}

type UserCreate struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	EmployeeCode int    `json:"employee_code"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	Password     string `json:"password"`
	Salt         string `json:"-"`
}

type SimpleUser struct {
	tableName    struct{}  `pg:"users"`
	Id           string    `json:"id"`
	FullName     string    `json:"full_name"`
	EmployeeCode int       `json:"employee_code"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"Email has already existed",
		"ErrEmailExisted",
	)
)

func (u *User) GetUserId() string {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}
