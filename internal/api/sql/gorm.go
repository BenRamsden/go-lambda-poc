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
	*gorm.DB
}

func NewSQLRepository(db *gorm.DB) Repository {
	return &repository{db}
}

// Migrate implements Repository.
func (r *repository) Migrate() error {
	return r.AutoMigrate(&Asset{})
}

// Base UUID model
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

// Pagination

func P(pagination model.Pagination) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		var page, count int = 0, 10
		if pagination.Page > 0 {
			page = pagination.Page
		}

		if pagination.Count > 0 && pagination.Count < 100 {
			count = pagination.Count
		}

		return tx.Limit(count).Offset(page * count)
	}

}
