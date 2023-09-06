package entity

import "time"

type User struct {
	tableName    struct{}  `pg:"users,alias:users"`
	Id           string    `json:"id"`
	FullName     string    `json:"full_name"`
	EmployeeCode int       `json:"employee_code"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
