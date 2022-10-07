package device

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//var db *gorm.DB
//var err error
var dbName string = "test-device.db"

func createSqliteDatabase() *FlyingDeviceGormRepository {
	db, _ := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	db.AutoMigrate(FlyingDevice{})
	return NewRepository(db)
}

func removeDatabase() {
	os.Remove(dbName)
}

func TestFlyingDeviceGormRepository_Create(t *testing.T) {
	repo := createSqliteDatabase()
	defer removeDatabase()

	device := &FlyingDevice{
		ID:         "ID",
		Model:      "Alpha 6",
		Make:       "Advance",
		Size:       "26",
		DeviceType: PARAGLIDER,
		Nickname:   "MyFirstTime",
		UserID:     "BIGBANG",
	}

	id, err := repo.Create(device)

	assert.NoError(t, err, "Threw an unexpected error")
	assert.NotEmpty(t, id, "ID shoul be set")

	var resultEntity FlyingDevice
	repo.context.First(&resultEntity)

	assert.Equal(t, device.Make, resultEntity.Make)
	assert.Equal(t, device.Model, resultEntity.Model)
	assert.Equal(t, device.Size, resultEntity.Size)
	assert.Equal(t, device.DeviceType, resultEntity.DeviceType)
	assert.Equal(t, device.Nickname, resultEntity.Nickname)
	assert.Equal(t, device.ID, resultEntity.ID)

}

func TestFlyingDeviceGormRepository_Update(t *testing.T) {
	repo := createSqliteDatabase()
	defer removeDatabase()

	device := &FlyingDevice{
		ID:         "ID",
		Model:      "Alpha 6",
		Make:       "Advance",
		Size:       "26",
		DeviceType: PARAGLIDER,
		Nickname:   "MyFirstTime",
		UserID:     "BIGBANG",
	}

	repo.Create(device)

	var resultBeforeUpdate FlyingDevice
	repo.context.First(&resultBeforeUpdate)

	device.Nickname = "Edit Nickname"
	device.Size = "27"
	device.DeviceType = HANGGLIDER
	device.Model = "New model"
	device.Make = "New make"

	repo.Update(device)

	var resultAfterUpdate FlyingDevice
	repo.context.First(&resultAfterUpdate)

	assert.True(t, resultBeforeUpdate.UpdatedAt.Before(*resultAfterUpdate.UpdatedAt), "Timestamp should be updated")

	assert.NotEqual(t, resultBeforeUpdate.Nickname, resultAfterUpdate.Nickname)
	assert.NotEqual(t, resultBeforeUpdate.Size, resultAfterUpdate.Size)
	assert.Equal(t, resultBeforeUpdate.UserID, resultAfterUpdate.UserID)
	assert.NotEqual(t, resultBeforeUpdate.Model, resultAfterUpdate.Model)
	assert.NotEqual(t, resultBeforeUpdate.Make, resultAfterUpdate.Make)
}

func TestFlyingDeviceGormRepository_GetDeviceByUser(t *testing.T) {
	repo := createSqliteDatabase()
	defer removeDatabase()

	var device FlyingDevice

	for i := 0; i < 10; i++ {
		device = FlyingDevice{
			ID:         fmt.Sprintf("ID-%v", i),
			Model:      "Alpha 6",
			Make:       "Advance",
			Size:       "26",
			DeviceType: PARAGLIDER,
			Nickname:   "MyFirstTime",
			UserID:     "UserId01",
		}

		repo.Create(&device)
	}

	devices, err := repo.GetDeviceByUser("UserId01")
	assert.NoError(t, err)
	assert.Len(t, devices, 10)
}

func TestFlyingDeviceGormRepository_GetDeviceByUser_When_No_Devices(t *testing.T) {
	repo := createSqliteDatabase()
	defer removeDatabase()

	devices, err := repo.GetDeviceByUser("Some-ID-without-any-devices")

	assert.NoError(t, err, "There should be no error when the user has no flying devices")

	assert.Len(t, devices, 0, "Empty array should be empty")
}
