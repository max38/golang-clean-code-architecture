package entityconstituency

import "time"

type ConstituencyEntity struct {
	Id        uint      `db:"id" gorm:"primaryKey" json:"id"`
	Name      string    `db:"name" json:"name"`
	Province  string    `db:"province" json:"province"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
