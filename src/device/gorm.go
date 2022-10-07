package device

import (
	"time"

	"gorm.io/gorm"
)

type FlyingDeviceGormRepository struct {
	context *gorm.DB
}

// NewRepository instansiates a repository obect
func NewRepository(db *gorm.DB) *FlyingDeviceGormRepository {
	repo := FlyingDeviceGormRepository{context: db}
	return &repo
}

// Create returns an string-based ID or an error
func (r *FlyingDeviceGormRepository) Create(device *FlyingDevice) (string, error) {
	timestamp := time.Now()
	device.CreatedAt = &timestamp
	device.UpdatedAt = &timestamp
	result := r.context.Create(device)

	return device.ID, result.Error
}

// Update returns an error if transaction fails
func (r *FlyingDeviceGormRepository) Update(device *FlyingDevice) error {
	timestamp := time.Now()
	device.UpdatedAt = &timestamp
	return r.context.Save(device).Error
}

// GetDeviceByUser get all flying devices by userID
func (r *FlyingDeviceGormRepository) GetDeviceByUser(userId string) ([]FlyingDevice, error) {
	var devices []FlyingDevice

	result := r.context.Where("userId = ?", userId).Find(&devices)

	return devices, result.Error
}
