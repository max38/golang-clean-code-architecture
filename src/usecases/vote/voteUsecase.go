package voteusecase

import (
	entityvote "github.com/max38/golang-clean-code-architecture/src/domain/entities/vote"
	repositories "github.com/max38/golang-clean-code-architecture/src/domain/repositories"
)

type IVoteUsecase interface {
	CountVoteByCandidateId(voteId uint) (*int64, error)
	CheckExistsVoteByUserId(userId uint) (bool, error)
	InsertVote(voteRegisterRequest *entityvote.VoteCreateWithUserRequest) (*entityvote.VoteEntity, error)
}

type voteUsecase struct {
	voteRepository repositories.IVoteRepository
}

func VoteUsecase(voteRepository repositories.IVoteRepository) IVoteUsecase {
	return &voteUsecase{
		voteRepository: voteRepository,
	}
}

func (u *voteUsecase) CountVoteByCandidateId(voteId uint) (*int64, error) {
	var count, err = u.voteRepository.CountVoteByCandidateId(voteId)
	if err != nil {
		return nil, err
	}

	return count, nil
}

func (u *voteUsecase) CheckExistsVoteByUserId(userId uint) (bool, error) {
	return u.voteRepository.CheckExistsVoteByUserId(userId)
}

func (u *voteUsecase) InsertVote(voteCreateRequest *entityvote.VoteCreateWithUserRequest) (*entityvote.VoteEntity, error) {

	var voteEntity, err = u.voteRepository.InsertVote(voteCreateRequest)
	if err != nil {
		return nil, err
	}

	return voteEntity, nil
}
