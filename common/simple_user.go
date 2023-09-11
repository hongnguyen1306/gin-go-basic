package common

import "time"

type SimpleUser struct {
	tableName    struct{}  `pg:"users"`
	Id           string    `json:"id"`
	FullName     string    `json:"fullName"`
	EmployeeCode int       `json:"employeeCode"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
