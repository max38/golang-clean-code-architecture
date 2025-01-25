package main

import (
	config "github.com/max38/golang-clean-code-architecture/src/config"
	entities "github.com/max38/golang-clean-code-architecture/src/domain/entities"
	gofiberserver "github.com/max38/golang-clean-code-architecture/src/infrastructure/gofiber/server"
	postgrescandidaterepository "github.com/max38/golang-clean-code-architecture/src/interface/repositories/postgres/candidate"
	postgresconstituencyrepository "github.com/max38/golang-clean-code-architecture/src/interface/repositories/postgres/constituency"
	postgresuserrepository "github.com/max38/golang-clean-code-architecture/src/interface/repositories/postgres/user"
	postgresvoterepository "github.com/max38/golang-clean-code-architecture/src/interface/repositories/postgres/vote"
	candidateusecase "github.com/max38/golang-clean-code-architecture/src/usecases/candidate"
	constituencyusecase "github.com/max38/golang-clean-code-architecture/src/usecases/constituency"
	userusecase "github.com/max38/golang-clean-code-architecture/src/usecases/user"
	voteusecase "github.com/max38/golang-clean-code-architecture/src/usecases/vote"
)

// @title Training Kawaii Shop API
// @version 1.0
// @description This is a sample swagger for Training Kawaii Shop
// @termsOfService http://swagger.io/terms/
// @contact.name Sukhum Butrkam
// @contact.email sukhum_butrkam@hotmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	config.Load(".env.dev")

	// User
	var userRepository = postgresuserrepository.UserRepository()
	var userUsecase = userusecase.UserUsecase(userRepository)

	// Constituency
	var constituencyRepository = postgresconstituencyrepository.ConstituencyRepository()
	var constituencyUsecase = constituencyusecase.ConstituencyUsecase(constituencyRepository)

	// Candidate
	var candidateRepository = postgrescandidaterepository.CandidateRepository()
	var candidateUsecase = candidateusecase.CandidateUsecase(candidateRepository)

	// Vote
	var voteRepository = postgresvoterepository.VoteRepository()
	var voteUsecase = voteusecase.VoteUsecase(voteRepository)

	var applicationEntity = entities.ApplicationEntity{
		UserUsecase:         userUsecase,
		ConstituencyUsecase: constituencyUsecase,
		CandidateUsecase:    candidateUsecase,
		VoteUsecase:         voteUsecase,
	}

	var server gofiberserver.IServer = gofiberserver.NewServer(applicationEntity)
	server.Start()
}
