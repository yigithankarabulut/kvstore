package kvstorage_test

import (
	"testing"

	"github.com/yigithankarabulut/kvstore/src/internal/storage/memory/kvstorage"
)

// db de bulunan keyin aynı adıyla yeni post edilince set fonksiyonundaki ilk getten error dönücek testi yaz TODO: testi yaz

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
