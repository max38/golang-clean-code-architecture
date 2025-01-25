package entitycandidate

import (
	"time"
)

type CandidateEntity struct {
	Id             uint      `db:"id" gorm:"primaryKey" json:"id"`
	Name           string    `db:"name" json:"name"`
	ConstituencyId uint      `db:"constituency_id" gorm:"index" json:"constituency_id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
