package moviesession

import (
	"cinema-admin/models"
	"time"
)

//Session ...
type Session struct {
	TimeStart *time.Time `json:"time_start,omitempty"`
	TimeEnd   *time.Time `json:"time_end,omitempty"`
}

// Display ...
type Display struct {
	models.Movie
	Sessions []Session `json:"sessions,omitempty"`
}
