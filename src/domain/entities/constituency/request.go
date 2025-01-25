package entityconstituency

type ConstituencyCreateRequest struct {
	Name     string `json:"name"`
	Province string `json:"province"`
}

type ConstituencyUpdateRequest struct {
	Id uint `json:"id"`
	ConstituencyCreateRequest
}
