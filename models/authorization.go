package models

import (
	"github.com/draco121/common/constants"
	"time"
)

type AuthorizationLog struct {
	UserId      string             `json:"userId"`
	Role        constants.Role     `json:"role"`
	Actions     []constants.Action `json:"actions"`
	RequestedAt time.Time          `json:"requestedAt"`
	Grant       constants.Grant    `json:"grant"`
	Reason      string             `json:"reason"`
}

type AuthorizationInput struct {
	Token   string             `json:"token"`
	Actions []constants.Action `json:"actions"`
}

type AuthorizationOutput struct {
	Grant  constants.Grant `json:"grant"`
	UserId string          `json:"userId"`
}
