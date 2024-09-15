package config

import (
	"errors"
	"fmt"
	"leal-technical-test/internal/domain/models"

	"golang.org/x/crypto/bcrypt"
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
		models.User{}, // Realiza la migración de User
		models.Store{},
	)
	if err != nil {
		m.logger.Error(fmt.Sprintf("Error al migrar la base de datos: %v", err))
		return err
	}

	// Crear un usuario por defecto después de migrar la tabla User
	defaultUser := models.User{
		Name:  "Admin",
		Email: "admin@example.com",
	}

	// Hashear la contraseña antes de guardarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		m.logger.Error(fmt.Sprintf("Error al hashear la contraseña: %v", err))
		return err
	}
	defaultUser.Password = string(hashedPassword)

	// Verificar si el usuario ya existe, si no, crearlo
	var user models.User
	result := m.db.Where("email = ?", defaultUser.Email).First(&user)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Crear el usuario por defecto si no existe
		if err := m.db.Create(&defaultUser).Error; err != nil {
			m.logger.Error(fmt.Sprintf("Error al crear el usuario por defecto: %v", err))
			return err
		}
		m.logger.Success("Usuario por defecto creado exitosamente")
	} else {
		m.logger.Info("El usuario por defecto ya existe, no se creó uno nuevo")
	}

	m.logger.Success("Migraciones completadas exitosamente")
	return nil
}
