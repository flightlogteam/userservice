package device

type Writer interface {
	// Create returns an string-based ID or an error
	Create(device *FlyingDevice) (string, error)
	// Update returns an error if transaction fails
	Update(device *FlyingDevice) error
}

type Reader interface {
	// GetDeviceByUser get all flying devices by userID
	GetDeviceByUser(userId string) ([]FlyingDevice, error)
	GetDeviceByID(deviceId string) (*FlyingDevice, error)
}

type FlyingDeviceRepository interface {
	Writer
	Reader
}

type FlyingDeviceService interface {
	Writer
	Reader
}
