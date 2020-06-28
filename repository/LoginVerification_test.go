package repository_test

import (
	"testing"

	"github.com/ariefsam/go-chat/configuration"
	"github.com/jinzhu/copier"

	"github.com/ariefsam/go-chat/entity"
	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/repository"
)

func TestLoginVerificationRepository(t *testing.T) {
	var repo repository.LoginVerification
	copier.Copy(&repo, &configuration.Config.MySQL)
	repo.AutoMigrate()
	repo.Flush()

	lv := entity.LoginVerification{
		ID:               "id001",
		DeviceID:         "dev001",
		PhoneNumber:      "625445",
		ExpiredTimestamp: 156600,
		UserID:           "id0001",
		VerificationCode: "332211",
	}
	t.Run("Normal save data", func(t *testing.T) {
		err := repo.Save(lv)
		assert.NoError(t, err)
	})

	t.Run("Get data", func(t *testing.T) {
		getLv := repo.Get(lv.PhoneNumber, lv.DeviceID, lv.ExpiredTimestamp-100, &lv.VerificationCode)
		assert.Equal(t, 1, len(getLv))
		if len(getLv) == 1 {
			assert.Equal(t, lv, getLv[0])
		}
	})

}
