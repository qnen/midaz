package command

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"

	"github.com/LerianStudio/midaz/components/ledger/internal/adapters/mongodb"
	"github.com/LerianStudio/midaz/pkg/mmodel"

	"github.com/stretchr/testify/assert"
)

// TestMetadataCreateSuccess is responsible to test MetadataCreate with success
func TestMetadataCreateSuccess(t *testing.T) {
	metadata := mongodb.Metadata{ID: primitive.NewObjectID()}
	collection := reflect.TypeOf(mmodel.Organization{}).Name()
	uc := UseCase{
		MetadataRepo: mongodb.NewMockRepository(gomock.NewController(t)),
	}

	uc.MetadataRepo.(*mongodb.MockRepository).
		EXPECT().
		Create(gomock.Any(), collection, &metadata).
		Return(nil).
		Times(1)

	err := uc.MetadataRepo.Create(context.TODO(), collection, &metadata)
	assert.Nil(t, err)
}

// TestMetadataCreateError is responsible to test MetadataCreate with error
func TestMetadataCreateError(t *testing.T) {
	errMSG := "err to create metadata on mongodb"
	metadata := mongodb.Metadata{ID: primitive.NewObjectID()}
	collection := reflect.TypeOf(mmodel.Organization{}).Name()
	uc := UseCase{
		MetadataRepo: mongodb.NewMockRepository(gomock.NewController(t)),
	}

	uc.MetadataRepo.(*mongodb.MockRepository).
		EXPECT().
		Create(gomock.Any(), collection, &metadata).
		Return(errors.New(errMSG)).
		Times(1)

	err := uc.MetadataRepo.Create(context.TODO(), collection, &metadata)
	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), errMSG)
}
