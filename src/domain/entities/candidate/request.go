package entitycandidate

type CandidateCreateRequest struct {
	Name           string `json:"name"`
	ConstituencyId uint   `json:"constituency_id"`
}

type CandidateUpdateRequest struct {
	Id uint `json:"id"`
	CandidateCreateRequest
}
