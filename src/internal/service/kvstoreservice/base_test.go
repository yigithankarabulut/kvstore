package kvstoreservice_test //nolint:typecheck

import (
	"errors"

	"github.com/yigithankarabulut/kvstore/src/internal/storage/memory/kvstorage"
)

var (
	errStorageDelete = errors.New("storage delete error")
	errStorageGet    = errors.New("storage get error")
	errStorageSet    = errors.New("storage set error")
	errStorageUpdate = errors.New("storage update error")
)

type mockStorage struct {
	deleteErr error
	getErr    error
	updateErr error
	setErr    error
	memoryDB  kvstorage.MemoryDB
}

func (m *mockStorage) Delete(_ string) error {
	return m.deleteErr
}

func (m *mockStorage) Get(_ string) (any, error) {
	return nil, m.getErr
}

func (m *mockStorage) List() kvstorage.MemoryDB {
	return m.memoryDB
}

func (m *mockStorage) Set(_ string, _ any) (any, error) {
	return nil, m.setErr
}

func (m *mockStorage) Update(_ string, _ any) (any, error) {
	return nil, m.updateErr
}
