package entityuser

import "time"

type UserEntity struct {
	Id             uint      `db:"id" gorm:"primaryKey" json:"id"`
	Email          string    `db:"email" gorm:"uniqueIndex" json:"email"`
	FirstName      string    `db:"first_name" json:"first_name"`
	LastName       string    `db:"last_name" json:"last_name"`
	ConstituencyId uint      `db:"constituency_id" gorm:"index;not null" json:"constituency_id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

type UserTokenEntity struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type UserTokenClaimsEntity struct {
	Id uint `json:"id"`
}
