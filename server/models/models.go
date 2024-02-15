package models

import (
	"time"

	"github.com/google/uuid"
)

type Vocation string

const (
	Knight   Vocation = "knight"
	Paladin  Vocation = "paladin"
	Druid    Vocation = "druid"
	Sorcerer Vocation = "sorcerer"
)

type World struct {
	Name string `json:"name" gorm:"primaryKey;unique;not null"`
}

type Boss struct {
	Name string `json:"name" gorm:"primaryKey;unique;not null"`
}

// TODO: implement handling templates for participants (e.g. 1ek1ed3any)
type Bosshunt struct {
	ID           uuid.UUID     `json:"id" gorm:"type:char(36);primary_key"`
	Boss         string        `json:"boss"`
	World        string        `json:"world"`
	When         time.Time     `json:"when"`
	Participants []Participant `json:"participants"`
	// DiscordRef  string    `json:"discord_ref"`
	// DiscordOnly bool      `json:"discord_only"`
}

// TODO: populate the db with some participants
type Participant struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	BosshuntID string    `json:"bosshunt_id"`
	Name       string    `json:"name"`
	Vocation   Vocation  `json:"vocation" gorm:"type:ENUM('knight', 'paladin', 'druid', 'sorcerer')"`
	UserID     string    `json:"user_id"`
	User       User      `json:"-"`
}

// TODO: implement signup/login with discord
type User struct {
	ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
}
