package kvstorage_test

import (
	"reflect"
	"testing"

	"github.com/yigithankarabulut/kvstore/src/internal/storage/memory/kvstorage"
)

func TestListEmpty(t *testing.T) {
	memoryStorage := kvstorage.MemoryDB(map[string]any{})

	storage := kvstorage.New(kvstorage.WithMemoryDB(memoryStorage))

	value := storage.List()

	if !reflect.DeepEqual(value, storage.List()) {
		t.Error("value not equal")
	}
}

func TestList(t *testing.T) {
	key := "key"
	memoryStorage := kvstorage.MemoryDB(map[string]any{
		key: "value",
	})
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	value := storage.List()

	if !reflect.DeepEqual(value, memoryStorage) {
		t.Error("value not equal")
	}
}
