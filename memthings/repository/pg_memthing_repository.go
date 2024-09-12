package repository

import (
	"github.com/JOn2Thou/memrizr/memthings/model"
	"gorm.io/gorm"
)

type MemthingRepository interface {
	CreateMemthing(memthing *model.Memthing) error
	GetMemthingByID(id uint) (*model.Memthing, error)
}

type PGMemthingRepository struct {
	db *gorm.DB
}

// NewPGMemthingRepository creates a new instance of PGMemthingRepository
func NewPGMemthingRepository(db *gorm.DB) *PGMemthingRepository {
	return &PGMemthingRepository{db: db}
}

// CreateMemthing inserts a new memthing record into the database
func (r *PGMemthingRepository) CreateMemthing(memthing *model.Memthing) error {
	return r.db.Create(memthing).Error
}

// GetMemthingByID fetches a memthing record by its ID
func (r *PGMemthingRepository) GetMemthingByID(id uint) (*model.Memthing, error) {
	var memthing model.Memthing
	if err := r.db.First(&memthing, id).Error; err != nil {
		return nil, err
	}
	return &memthing, nil
}
