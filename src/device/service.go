package device

import "github.com/pkg/errors"

type DeviceService struct {
	deviceRepo FlyingDeviceRepository
}

func NewDeviceService(deviceRepo FlyingDeviceRepository) FlyingDeviceService {
	return &DeviceService{
		deviceRepo: deviceRepo,
	}
}

// Create returns an string-based ID or an error
func (s *DeviceService) Create(device *FlyingDevice) (string, error) {

	if !device.IsValid() {
		return "", errors.New("Device missing mandatory fields")
	}

	id, err := s.deviceRepo.Create(device)

	if err != nil {
		return "", errors.Wrap(err, "Unable to create flying-device")
	}

	return id, nil
}

// Update returns an error if transaction fails
func (s *DeviceService) Update(device *FlyingDevice) error {
	if !device.IsValid() {
		return errors.New("Device missing mandatory fields")
	}

	err := s.deviceRepo.Update(device)

	if err != nil {
		return errors.Wrap(err, "Unable to update flying-device")
	}

	return nil
}

// GetDeviceByUser get all flying devices by userID
func (s *DeviceService) GetDeviceByUser(userId string) ([]FlyingDevice, error) {
	devices, err := s.deviceRepo.GetDeviceByUser(userId)

	if err != nil {
		return nil, errors.Wrap(err, "Unable to fetch users")
	}

	return devices, nil
}

func (s *DeviceService) GetDeviceByID(deviceId string) (*FlyingDevice, error) {
	return s.deviceRepo.GetDeviceByID(deviceId)
}
