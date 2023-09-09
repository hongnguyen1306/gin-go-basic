package common

const CurrentUser = "user"

type Requester interface {
	GetUserId() string
	GetEmail() string
	GetRole() string
}
