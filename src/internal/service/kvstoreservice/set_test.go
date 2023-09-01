package kvstoreservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/yigithankarabulut/kvstore/src/internal/service/kvstoreservice"
)

func TestSetWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := kvsStoreService.Set(ctx, nil); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestSetWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		setErr: errStorageSet,
	}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	setRequest := kvstoreservice.SetRequest{
		Key:   "vigo",
		Value: "vigo",
	}
	if _, err := kvsStoreService.Set(context.Background(), &setRequest); !errors.Is(
		err,
		errStorageSet,
	) {
		t.Error("error not occurred")
	}
}

func TestSet(t *testing.T) {
	mockStorage := &mockStorage{
		memoryDB: map[string]any{
			"key": "value",
		},
	}
	kvsStoreService := kvstoreservice.New(kvstoreservice.WithStorage(mockStorage))

	setRequest := kvstoreservice.SetRequest{
		Key:   "key",
		Value: "value",
	}
	if _, err := kvsStoreService.Set(context.Background(), &setRequest); err != nil {
		t.Error("error occurred")
	}
}
