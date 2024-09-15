package repository

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Interfaz que representa la conexión a la base de datos
type IDatabaseConnection interface {
	GetDB() *gorm.DB
	Connect() error
	Close() error
}

// MockDBConnection simula la conexión a la base de datos
type MockDBConnection struct {
	DB *gorm.DB
}

// Simula el método GetDB
func (m *MockDBConnection) GetDB() *gorm.DB {
	return m.DB
}

// Simula el método Close
func (m *MockDBConnection) Close() error {
	return nil // No hace nada en este mock
}

// Simula el método Connect
func (m *MockDBConnection) Connect() error {
	return nil // Simulación simple que no realiza ninguna acción
}

// Simula el método Ping
func (m *MockDBConnection) Ping() error {
	return nil // Simulación simple que no realiza ninguna acción
}

// setupTestDB configura una base de datos SQLite en memoria para las pruebas
func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrar los modelos necesarios
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Prueba para el método Create de UserRepository
func TestCreateUser(t *testing.T) {
	// Configura la base de datos en memoria
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}

	// Crear el mock de la conexión a la base de datos
	mockDB := &MockDBConnection{DB: db}

	// Inicializar el repositorio con el mock
	userRepo := repository.NewUserRepository(mockDB)

	// Crear un nuevo usuario
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	err = userRepo.Create(&user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Verificar que el ID del usuario se haya asignado correctamente
	if user.ID == 0 {
		t.Errorf("Expected user ID to be set after creation")
	}
}
