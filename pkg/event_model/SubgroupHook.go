package event_model

import (
	"time"
)

type CreateOrRemoveSubgroupEvent struct {
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	EventName      string    `json:"event_name"`
	Name           string    `json:"name"`
	Path           string    `json:"path"`
	FullPath       string    `json:"full_path"`
	GroupID        int       `json:"group_id"`
	ParentGroupID  int       `json:"parent_group_id"`
	ParentName     string    `json:"parent_name"`
	ParentPath     string    `json:"parent_path"`
	ParentFullPath string    `json:"parent_full_path"`
}