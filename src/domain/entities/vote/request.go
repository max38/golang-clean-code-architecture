package entityvote

type VoteCreateRequest struct {
	CandidateId uint `json:"candidate_id"`
}

type VoteCreateWithUserRequest struct {
	VoteCreateRequest
	UserId uint `json:"user_id"`
}
