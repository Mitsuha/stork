package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID                   int            `json:"id" gorm:"column:id;primary_key"`
	Name                 string         `json:"name" gorm:"column:name"`
	Email                string         `json:"email" gorm:"column:email"`
	Password             string         `json:"-" gorm:"column:password"`
	IsAdmin              int            `json:"is_admin" gorm:"column:is_admin"`
	Preferences          sql.NullString `json:"preferences" gorm:"column:preferences"`
	RememberToken        sql.NullString `json:"remember_token" gorm:"column:remember_token"`
	CreatedAt            time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt            time.Time      `json:"updated_at" gorm:"column:updated_at"`
	InvitationToken      sql.NullString `json:"invitation_token" gorm:"column:invitation_token"`
	InvitedAt            time.Time      `json:"invited_at" gorm:"column:invited_at"`
	InvitationAcceptedAt time.Time      `json:"invitation_accepted_at" gorm:"column:invitation_accepted_at"`
	InvitedByID          sql.NullInt64  `json:"invited_by_id" gorm:"column:invited_by_id"`
}
