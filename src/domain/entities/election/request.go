package election

import "encoding/json"

type ElectionNewresultRequest struct {
	Id           int           `json:"id"`
	Name         string        `json:"name"`
	SeqNo        int           `json:"seqNo"`
	PartyResults []PartyResult `json:"partyResults"`
}

type PartyResult struct {
	Party string      `json:"party"`
	Votes uint        `json:"votes"`
	Share json.Number `json:"share"`
}

type Scoreboard struct {
	Seats      map[string]int     `json:"seats"`
	Winner     string             `json:"winner"`
	TotalVotes map[string]uint    `json:"total_votes"`
	VoteShare  map[string]float64 `json:"vote_share"`
}
