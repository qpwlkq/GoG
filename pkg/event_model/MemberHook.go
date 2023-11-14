package event_model

import (
	"time"
)

type GroupNumberEvent struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	GroupName    string    `json:"group_name"`
	GroupPath    string    `json:"group_path"`
	GroupID      int       `json:"group_id"`
	UserUsername string    `json:"user_username"`
	UserName     string    `json:"user_name"`
	UserEmail    string    `json:"user_email"`
	UserID       int       `json:"user_id"`
	GroupAccess  string    `json:"group_access"`
	GroupPlan    any       `json:"group_plan"`
	ExpiresAt    time.Time `json:"expires_at"`
	EventName    string    `json:"event_name"`
}