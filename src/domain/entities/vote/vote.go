package entityvote

import "time"

type VoteEntity struct {
	Id          uint      `db:"id" gorm:"primaryKey" json:"id"`
	CandidateId uint      `db:"candidate_id" gorm:"index;not null" json:"candidate_id"`
	UserId      uint      `db:"user_id" gorm:"index;not null" json:"user_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
