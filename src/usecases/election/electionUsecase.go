package election

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	entityelection "github.com/max38/golang-clean-code-architecture/src/domain/entities/election"
)

type IElectionUsecase interface {
	AddResult(result entityelection.ElectionNewresultRequest) error
	GetResult(id int) (entityelection.ElectionNewresultRequest, error)
	GetAll() ([]entityelection.ElectionNewresultRequest, error)
	GetScoreboard() (entityelection.Scoreboard, error)
	Reset() error
}

type electionUsecase struct {
	data map[int]entityelection.ElectionNewresultRequest
}

func ElectionUsecase(data map[int]entityelection.ElectionNewresultRequest) *electionUsecase {
	return &electionUsecase{
		data: data,
	}
}

type NotFound string

func (e NotFound) Error() string {
	return string(e)
}

func (e electionUsecase) AddResult(result entityelection.ElectionNewresultRequest) error {
	e.data[result.Id] = result
	return nil
}

func (e electionUsecase) GetResult(id int) (entityelection.ElectionNewresultRequest, error) {
	res, ok := e.data[id]
	if !ok {
		return entityelection.ElectionNewresultRequest{}, NotFound(fmt.Sprintf("%d", id))
	}
	return res, nil
}

func (e electionUsecase) GetAll() ([]entityelection.ElectionNewresultRequest, error) {
	var results []entityelection.ElectionNewresultRequest
	for _, v := range e.data {
		results = append(results, v)
	}
	return results, nil
}

func (e electionUsecase) GetScoreboard() (*entityelection.Scoreboard, error) {
	err := e.Reset()
	if err != nil {
		return nil, err
	}

	err = e.autoAddResult(650)
	if err != nil {
		return nil, err
	}

	// Retrieve all election results from the result service
	results, err := e.GetAll()
	if err != nil {
		return nil, err
	}

	// Calculate seats and total votes for each party
	seats, totalVotes, voteShare, totalAllVotes := e.calseatsAndTotalVote(results)

	// Determine the winning party if it has reached the required seat count
	var winner string
	for party, seatCount := range seats {
		if seatCount >= 325 {
			winner = party
			break
		}
	}

	// Calculate the percentage share of votes for each party
	for party, votes := range totalVotes {
		voteShare[party] = (float64(votes) / float64(totalAllVotes)) * 100
	}

	// Construct the scoreboard response with calculated data
	result := &entityelection.Scoreboard{
		Seats:      seats,
		Winner:     winner,
		TotalVotes: totalVotes,
		VoteShare:  voteShare,
	}

	return result, nil
}

func (e electionUsecase) Reset() error {
	e.data = make(map[int]entityelection.ElectionNewresultRequest, 0)
	return nil
}

func (e electionUsecase) calseatsAndTotalVote(results []entityelection.ElectionNewresultRequest) (seats map[string]int, totalVotes map[string]uint, voteShare map[string]float64, totalAllVotes uint) {
	seats = make(map[string]int)
	totalVotes = make(map[string]uint)
	voteShare = make(map[string]float64)

	for _, result := range results {
		var maxVotes uint
		var winnerParty string
		for _, partyResult := range result.PartyResults {
			totalVotes[partyResult.Party] += partyResult.Votes
			totalAllVotes += partyResult.Votes
			if partyResult.Votes > maxVotes {
				maxVotes = partyResult.Votes
				winnerParty = partyResult.Party
			}
		}
		// Increment seat count for the party with the most votes in the constituency
		if winnerParty != "" {
			seats[winnerParty]++
		}
	}

	return seats, totalVotes, voteShare, totalAllVotes
}

func (e electionUsecase) autoAddResult(results int) error {
	for i := 1; i <= results; i++ {
		filename := fmt.Sprintf("./api/sample-election-results/result%03d.json", i)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		var cr entityelection.ElectionNewresultRequest
		err = json.Unmarshal(data, &cr)
		if err != nil {
			return err
		}

		err = e.AddResult(cr)
		if err != nil {
			return err
		}
	}
	return nil
}
