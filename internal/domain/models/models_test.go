package models

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Configura una base de datos en memoria para realizar pruebas
func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrar los modelos
	err = db.AutoMigrate(&Store{}, &Branch{}, &Campaign{}, &AccumulatedReward{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Prueba para crear una tienda y asociar una sucursal
func TestCreateBranchForStore(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}

	store := Store{Name: "Test Store", ConversionFactor: 1.0}
	db.Create(&store)

	branch := Branch{
		Name:    "Test Branch",
		StoreID: store.ID,
		Address: "123 Main St",
	}
	db.Create(&branch)

	// Verificar si la sucursal está asociada a la tienda correctamente
	if branch.StoreID != store.ID {
		t.Errorf("Expected StoreID to be %d, got %d", store.ID, branch.StoreID)
	}
}

// Prueba para crear una campaña asociada a una sucursal
func TestCreateCampaignForBranch(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}

	branch := Branch{Name: "Test Branch", Address: "123 Main St"}
	db.Create(&branch)

	campaign := Campaign{
		Name:       "Double Points",
		BranchID:   branch.ID,
		Type:       "double",
		Percentage: 20.0,
		StartDate:  time.Now(),
		EndDate:    time.Now().AddDate(0, 0, 7),
	}
	db.Create(&campaign)

	// Verificar si la campaña está asociada a la sucursal
	if campaign.BranchID != branch.ID {
		t.Errorf("Expected BranchID to be %d, got %d", branch.ID, campaign.BranchID)
	}
}

// Prueba para agregar recompensas acumuladas para un usuario y una tienda
func TestAddAccumulatedReward(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}

	store := Store{Name: "Test Store", ConversionFactor: 1.5}
	db.Create(&store)

	user := User{Name: "John Doe"} // Asume que tienes una estructura de usuario
	db.Create(&user)

	reward := AccumulatedReward{
		UserID:              user.ID,
		StoreID:             store.ID,
		PointsAccumulated:   100.0,
		CashbackAccumulated: 10.0,
	}
	db.Create(&reward)

	// Verificar si las recompensas acumuladas se crearon correctamente
	if reward.UserID != user.ID {
		t.Errorf("Expected UserID to be %d, got %d", user.ID, reward.UserID)
	}
	if reward.StoreID != store.ID {
		t.Errorf("Expected StoreID to be %d, got %d", store.ID, reward.StoreID)
	}
	if reward.PointsAccumulated != 100.0 {
		t.Errorf("Expected PointsAccumulated to be 100.0, got %f", reward.PointsAccumulated)
	}
}
