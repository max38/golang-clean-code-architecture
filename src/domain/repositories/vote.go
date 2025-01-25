package repositories

import (
	entityvote "github.com/max38/golang-clean-code-architecture/src/domain/entities/vote"
)

type IVoteRepository interface {
	CountVoteByCandidateId(candidateId uint) (*int64, error)
	CheckExistsVoteByUserId(userId uint) (bool, error)
	InsertVote(voteRegisterRequest *entityvote.VoteCreateWithUserRequest) (*entityvote.VoteEntity, error)
}
