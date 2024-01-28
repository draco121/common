package models

import (
	"github.com/draco121/common/constants"
	"time"
)

type AuthorizationLog struct {
	UserId      string             `json:"userId"`
	Role        constants.Role     `json:"role"`
	Actions     []constants.Action `json:"action"`
	RequestedAt time.Time          `json:"requestedAt"`
	Grant       constants.Grant    `json:"grant"`
	Reason      string             `json:"reason"`
}
