<<<<<<< Updated upstream
package gofiber
=======
// filepath: /Users/noppachai/Documents/สมัครงาน/Testสัมงาน/golang-clean-code-architecture/src/infrastructure/gofiber/main.go
package main
>>>>>>> Stashed changes

import (
	config "github.com/max38/golang-clean-code-architecture/src/config"
	entities "github.com/max38/golang-clean-code-architecture/src/domain/entities"
	_ "github.com/max38/golang-clean-code-architecture/src/infrastructure/gofiber/docs" // Import generated docs
	gofiberserver "github.com/max38/golang-clean-code-architecture/src/infrastructure/gofiber/server"
	"github.com/max38/golang-clean-code-architecture/src/infrastructure/gofiber/swagger" // Import the swagger setup
	postgresuserrepository "github.com/max38/golang-clean-code-architecture/src/interface/repositories/postgres/user"
	userusecase "github.com/max38/golang-clean-code-architecture/src/usecases/user"
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
	// Load environment variables from a configuration file
	config.Load(".env.dev")

	// Initialize repositories and use cases
	var userRepository = postgresuserrepository.UserRepository()
	var userUsecase = userusecase.UserUsecase(userRepository)

	// Create an application entity with the user usecase
	var applicationEntity = entities.ApplicationEntity{
		UserUsecase: userUsecase,
	}

	// Initialize the server and setup Swagger
	var server gofiberserver.IServer = gofiberserver.NewServer(applicationEntity)

<<<<<<< Updated upstream
	// Start the server
=======
	swagger.SetupSwagger(server.App())

>>>>>>> Stashed changes
	server.Start()
}
