package internal

import (
	"leal-technical-test/config"

	"github.com/gin-gonic/gin"
)

// Server estructura que encapsula la lógica de inicialización del servidor
type Server struct {
	address   string
	ginMode   string
	ginServer *gin.Engine
	logger    config.ILogger
}

// NewServer es el constructor de la clase que inicializa el servidor
func NewServer() (*Server, error) {
	env := config.NewGetEnv()
	logger := config.NewLogger()
	ginServer := gin.New()
	gin.SetMode(env.GinMode)

	return &Server{
		address:   env.ServerPort,
		ginMode:   env.GinMode,
		ginServer: ginServer,
		logger:    logger,
	}, nil
}

// Run inicia el servidor
func (s *Server) Run() error {
	db := config.NewPostgresConnection()
	config.NewMigrator(db)
	defer db.Close()

	s.logger.Success("Starting server on =>", s.address)
	if err := s.ginServer.Run(s.address); err != nil {
		s.logger.Fatal("Failed to start server: ", err)
		return err
	}
	return nil
}
