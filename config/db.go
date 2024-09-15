package config

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type IDatabaseConnection interface {
	GetDB() *gorm.DB
	Connect() error
	Close() error
	Ping() error
}

type postgresConnection struct {
	host       string
	port       string
	user       string
	password   string
	name       string
	sslMode    string
	gormMode   string
	connection *gorm.DB
	logger     ILogger
}

var (
	dbConnectionInstance *postgresConnection
	databaseOnce         sync.Once
)

func NewPostgresConnection() IDatabaseConnection {
	databaseOnce.Do(func() {
		host := NewGetEnv().PostgresDBHost
		port := NewGetEnv().PostgresDBPort
		user := NewGetEnv().PostgresDBUser
		password := NewGetEnv().PostgresDBPassword
		name := NewGetEnv().PostgresDBName
		sslMode := NewGetEnv().PostgresDBSSLMode
		gormMode := NewGetEnv().GormMode

		// Initialize database instance
		dbConnectionInstance = &postgresConnection{
			host:     host,
			port:     port,
			user:     user,
			password: password,
			name:     name,
			sslMode:  sslMode,
			gormMode: gormMode,
		}

		// Initialize logger
		dbConnectionInstance.logger = NewLogger()

		// Read environment variables
		err := dbConnectionInstance.readEnvironmentVariables()
		if err != nil {
			dbConnectionInstance.logger.Fatal("failed to read environment variables: %v", err)
		}

		// Connect to database
		err = dbConnectionInstance.Connect()
		if err != nil {
			dbConnectionInstance.logger.Fatal("failed to connect to database: %v", err)
		}

		// Ping database
		err = dbConnectionInstance.Ping()
		if err != nil {
			dbConnectionInstance.logger.Fatal("failed to ping database: %v", err)
		}

		dbConnectionInstance.logger.Success("connected to database")
	})

	// Return database instance
	return dbConnectionInstance
}

func (p *postgresConnection) readEnvironmentVariables() error {

	if p.host == "" {
		return fmt.Errorf("missing required environment variable: POSTGRES_DB_HOST")
	}

	if p.port == "" {
		return fmt.Errorf("missing required environment variable: POSTGRES_DB_PORT")
	}

	if p.user == "" {
		return fmt.Errorf("missing required environment variable: POSTGRES_DB_USER")
	}

	if p.password == "" {
		return fmt.Errorf("missing required environment variable: POSTGRES_DB_PASSWORD")
	}

	if p.name == "" {
		return fmt.Errorf("missing required environment variable: POSTGRES_DB_NAME")
	}

	if p.sslMode == "" {
		return fmt.Errorf("missing required environment variable: POSTGRES_DB_SSLMODE")
	}

	return nil
}

func (p *postgresConnection) Connect() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.host, p.port, p.user, p.password, p.name, p.sslMode)

	gorMode := p.gormMode

	logMode := gormLogger.Silent
	if gorMode == "on" {
		logMode = gormLogger.Info
	}

	var err error
	p.connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: false, Logger: gormLogger.Default.LogMode(logMode)})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	return nil
}

func (p *postgresConnection) Close() error {
	sqlDB, err := p.connection.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}

	return nil
}

func (p *postgresConnection) Ping() error {
	if p.connection == nil {
		return fmt.Errorf("connection is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sqlDB, err := p.connection.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	err = sqlDB.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

func (p *postgresConnection) GetDB() *gorm.DB {
	return p.connection
}
