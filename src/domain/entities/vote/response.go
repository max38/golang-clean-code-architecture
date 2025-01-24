package entityvote

import (
	entitycandidate "github.com/max38/golang-clean-code-architecture/src/domain/entities/candidate"
)

type GetVotesByConstituencyIdResponse struct {
	Candidate entitycandidate.CandidateEntity `json:"candidate"`
	VoteCount int64                           `json:"vote_count"`
}

type VoteCreateResponse struct {
	Vote *VoteEntity `json:"vote"`
}
