package models

import "time"

// TODO: create DTOs?
type Vocation string

const (
	Knight   Vocation = "Knight"
	Paladin  Vocation = "Paladin"
	Druid    Vocation = "Druid"
	Sorcerer Vocation = "Sorcerer"
)

type World struct {
	Name string `json:"name" gorm:"primaryKey;unique;not null"`
}

type Boss struct {
	Name string `json:"name" gorm:"primaryKey;unique;not null"`
}

type Bosshunt struct {
	ID    string    `json:"id" gorm:"primary_key"`
	Boss  string    `json:"boss"`
	World string    `json:"world"`
	When  time.Time `json:"when"`
	// DiscordRef  string    `json:"discord_ref"`
	// DiscordOnly bool      `json:"discord_only"`
	// Participants []string  `json:"participants"`
}

type Participant struct {
	ID         string   `json:"id" gorm:"primary_key"`
	BosshuntID string   `json:"bosshunt_id"`
	Name       string   `json:"name"`
	Vocation   Vocation `json:"vocation" gorm:"type:ENUM('Knight', 'Paladin', 'Druid', 'Sorcerer')"`
	// UserID     string `json:"user_id"`
}
