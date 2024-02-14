package sql

import (
	"time"

	"github.com/google/uuid"
	"github.com/jugo-io/go-poc/internal/api/model"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate() error
	model.AssetRepository
}

type repository struct {
	DB *gorm.DB
}

func NewSQLRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

// Migrate implements Repository.
func (*repository) Migrate() error {
	panic("unimplemented")
}

type UUIDModel struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *UUIDModel) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now
	return nil
}

func (m *UUIDModel) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
