package entity

import (
	"app/common"
	"app/component/tokenprovider"
	"errors"
	"time"
)

type User struct {
	tableName    struct{}  `pg:"users"`
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
	tableName struct{} `pg:"users"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
}

type UserUpdate struct {
	tableName struct{}  `pg:"users"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
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
	FullName     string `json:"fullName"`
	EmployeeCode int    `json:"employeeCode"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	Password     string `json:"password"`
	Salt         string `json:"-"`
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

const (
	Member     string = "MEMBER"
	SuperAdmin string = "SUPER_ADMIN"
)
