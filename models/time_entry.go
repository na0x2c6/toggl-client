package models

import (
	"time"
)

type TimeEntry struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	WorkspaceID int        `json:"workspace_id"`
	ProjectID   *int       `json:"project_id"`
	TaskID      *int       `json:"task_id"`
	Billable    bool       `json:"billable"`
	Start       time.Time  `json:"start"`
	Stop        *time.Time `json:"stop"`
	Duration    int        `json:"duration"`
	CreatedWith string     `json:"created_with"`
	Tags        []string   `json:"tags"`
	TagIDs      []int      `json:"tag_ids"`
	DurOnly     bool       `json:"duronly"`
	At          time.Time  `json:"at"`
}

type TimeEntriesResponse struct {
	Items []TimeEntry `json:"items"`
}
