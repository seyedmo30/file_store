package entity

import "time"

type User struct {
	ID               uint      `gorm:"primaryKey"`
	Username         string    `gorm:"unique;not null"`
	PasswordHash     string    `gorm:"not null"`
	RegistrationDate time.Time `gorm:"default:current_timestamp"`
}
