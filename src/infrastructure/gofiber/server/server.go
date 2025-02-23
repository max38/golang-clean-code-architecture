package gofiberserver

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "os/signal"

    "github.com/gofiber/fiber/v2"
    "github.com/max38/golang-clean-code-architecture/src/config"
    "github.com/max38/golang-clean-code-architecture/src/domain/entities"
)

type IServer interface {
<<<<<<< Updated upstream
	Start()
	App() *fiber.App
=======
    Start()
    App() *fiber.App
>>>>>>> Stashed changes
}

type gofiberServer struct {
    app         *fiber.App
    application entities.ApplicationEntity
}

func NewServer(application entities.ApplicationEntity) IServer {
    return &gofiberServer{
        app: fiber.New(
            fiber.Config{
                AppName:     config.Config.String("GOAPP_NAME"),
                JSONEncoder: json.Marshal,
                JSONDecoder: json.Unmarshal,
            },
        ),
        application: application,
    }
}

func (s *gofiberServer) Start() {
    // Init Middleware
    var middlewares IMiddlewares = InitMiddlewares()
    s.app.Use(middlewares.Cors())
    s.app.Use(middlewares.Logger())

    // Init Module
    var routerGroupApiV1 = s.app.Group("/api/v1")
    var modules = InitModule(routerGroupApiV1, s, middlewares)
    modules.MonitorModule()
    modules.UsersModule()
    modules.SwaggerModule()
    modules.CRUDModule()

    // Router Check Middleware
    s.app.Use(middlewares.RouterCheck())

    // Graceful Shutdown
    var channelSignalInterrupt = make(chan os.Signal, 1)
    signal.Notify(channelSignalInterrupt, os.Interrupt)
    go func() {
        sig := <-channelSignalInterrupt
        log.Printf("Received signal %v, shutting down...", sig)
        if err := s.app.Shutdown(); err != nil {
            log.Fatalf("Error shutting down server: %v", err)
        }
        log.Println("Server has shut down gracefully.")
    }()

    // Listen to host:port
    var serverListen string = fmt.Sprintf("%s:%d", config.Config.String("GOAPP_HOST"), config.Config.Int("GOAPP_PORT"))
    log.Printf("Server is starting on %v", serverListen)

    // Error handling on Listen
    if err := s.app.Listen(serverListen); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
<<<<<<< Updated upstream
func (s *gofiberServer) App() *fiber.App {
	return s.app
}
=======

func (s *gofiberServer) App() *fiber.App {
    return s.app
}
>>>>>>> Stashed changes
