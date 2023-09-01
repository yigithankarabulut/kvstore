package kvstorage

func (ms *memoryStorage) Update(key string, value any) (any, error) {
	if _, err := ms.Get(key); err != nil { // can not update! key doesn't exist
		return nil, err
	}

	res, _ := ms.Set(key, value)
	return res, nil
}
