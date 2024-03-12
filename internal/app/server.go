package app

import (
	"Library_WebAPI/internal/handler"
	"Library_WebAPI/pkg/config"
	"Library_WebAPI/pkg/database"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hashicorp/go-hclog"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
	"syscall"
)

type Handler interface {
	HealthCheck(c *fiber.Ctx) error

	CreateAuthor(c *fiber.Ctx) error
	ReadAuthors(c *fiber.Ctx) error
	UpdateAuthor(c *fiber.Ctx) error
	DeleteAuthor(c *fiber.Ctx) error

	CreateBook(c *fiber.Ctx) error
	ReadBooks(c *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error

	CreateReader(c *fiber.Ctx) error
	ReadReaders(c *fiber.Ctx) error
	UpdateReader(c *fiber.Ctx) error
	DeleteReader(c *fiber.Ctx) error

	GetAuthorBooks(c *fiber.Ctx) error
	GetReaderBooks(c *fiber.Ctx) error
}

type Server struct {
	App     *fiber.App
	Logger  hclog.Logger
	Handler Handler
}

func selectiveLogging(c *fiber.Ctx) error {
	log := logger.New()

	if c.Path() == "/healthz" || c.Path() == "/metrics" {
		return c.Next()
	}

	return log(c)
}

func Start(conf *config.Config) error {
	s := new(Server)

	s.App = fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	s.App.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, DELETE, PATCH",
		AllowHeaders:     "Origin,X-Requested-With, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           120,
	}))

	logger := hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
		Level:      hclog.Debug,
	})

	s.Logger = logger.Named("Server")

	var err error

	db, err := database.ConnectDB(conf.Db, s.Logger.Named("DB Connection"))
	if err != nil {
		return err
	}
	defer db.Close()

	s.Handler = handler.NewHandler(db, logger)

	prometheus := fiberprometheus.New("library-web-api")
	prometheus.RegisterAt(s.App, "/metrics")
	s.App.Use(prometheus.Middleware)

	s.NewRouter()

	s.App.Use(selectiveLogging)

	go func() {
		if err = s.App.Listen(":" + conf.Port); err != nil {
			s.Logger.Error("Failed listen", "error", err)
			return
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c

	s.Logger.Info("Gracefully shutting down...")
	_ = s.App.Shutdown()

	s.Logger.Info("Running cleanup tasks...")
	return err
}
