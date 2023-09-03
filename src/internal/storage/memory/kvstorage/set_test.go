package kvstorage_test

import (
	"testing"

	"github.com/yigithankarabulut/kvstore/src/internal/storage/memory/kvstorage"
)

func TestSetEqualData(t *testing.T) {
	key := "key"
	memoryStorage := kvstorage.MemoryDB(map[string]any{})
	storage := kvstorage.New(kvstorage.WithMemoryDB(memoryStorage))
	_, _ = storage.Set(key, "value")

	if _, err := storage.Set(key, "value"); err == nil {
		t.Error("updated although key exists")
	}
}

func TestSet(t *testing.T) {
	key := "key"
	memoryStorage := kvstorage.MemoryDB(map[string]any{})
	storage := kvstorage.New(
		kvstorage.WithMemoryDB(memoryStorage),
	)

	_, _ = storage.Set(key, "value") //nolint:errcheck

	if _, ok := memoryStorage[key]; !ok {
		t.Error("value not equal")
	}
}
