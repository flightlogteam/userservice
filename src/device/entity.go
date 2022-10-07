package device

import "time"

type FlyingDevice struct {
	ID         string `gorm:"primaryKey;column:id;size:64"`
	Model      string
	Make       string
	Size       string
	DeviceType FlyingDeviceType
	Nickname   string
	UpdatedAt  *time.Time `gorm:"column:updatedat;type:time"`
	CreatedAt  *time.Time `gorm:"column:createdat;type:time"`
	DeletedAt  *time.Time `gorm:"column:deletedat;type:time"`
	UserID     string     `gorm:"size:64;column:userId"`
}

func (FlyingDevice) TableName() string {
	return "FlyingDevice"
}

type FlyingDeviceType int

const (
	PARAGLIDER FlyingDeviceType = 0
	SPEEDRIDER FlyingDeviceType = 1
	HANGGLIDER FlyingDeviceType = 2
	AIRBALLOON FlyingDeviceType = 3
)

func (f *FlyingDevice) IsValid() bool {
	return (len(f.Model) > 0 && len(f.Make) > 0 && f.DeviceType < 4 && f.DeviceType > 0)
}
