package config

import (
	"fmt"
	"leal-technical-test/internal/domain/models"

	"gorm.io/gorm"
)

// Migrator es la clase encargada de las migraciones de la base de datos
type Migrator struct {
	db     *gorm.DB
	logger ILogger
}

// NewMigrator es el constructor de la clase que inicializa la conexión
func NewMigrator(connection IDatabaseConnection) (*Migrator, error) {
	// Obtener la instancia de la conexión a la base de datos
	db := connection.GetDB()
	logger := NewLogger()
	if db == nil {
		return nil, fmt.Errorf("failed to get database connection")
	}

	return &Migrator{
		db:     db,
		logger: logger,
	}, nil
}

// Migrate realiza las migraciones de las entidades de la base de datos
func (m *Migrator) Migrate() error {
	// Realiza las migraciones de las entidades de la base de datos
	err := m.db.AutoMigrate(
		models.AccumulatedReward{},
		models.Branch{},
		models.Campaign{},
		models.Reward{},
		models.Transaction{},
		models.User{},
		models.Store{},
	)

	if err != nil {
		m.logger.Error(fmt.Sprintf("Error al migrar la base de datos: %v", err))
		return err
	}

	m.logger.Success("Migraciones completadas exitosamente")
	return nil
}
