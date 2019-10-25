// Code generated by sqlc. DO NOT EDIT.

package ondeck

import (
	"database/sql"
	"time"
)

type Status string

const (
	StatusOpen   Status = "open"
	StatusClosed Status = "closed"
)

type City struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Venue struct {
	ID              int32          `json:"id"`
	Status          Status         `json:"status"`
	Slug            string         `json:"slug"`
	Name            string         `json:"name"`
	City            string         `json:"city"`
	SpotifyPlaylist string         `json:"spotify_playlist"`
	SongkickID      sql.NullString `json:"songkick_id"`
	Tags            []string       `json:"tags"`
	CreatedAt       time.Time      `json:"created_at"`
}
