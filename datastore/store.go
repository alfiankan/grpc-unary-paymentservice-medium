package datastore

import "sync"

type DataStore struct {
	sync.RWMutex
	PaymentOrder map[string]string
}

func NewDataStore() DataStore {
	return DataStore{
		PaymentOrder: make(map[string]string),
	}
}

func (store *DataStore) Insert(orderID string, paymentMethod string) {
	store.PaymentOrder[orderID] = paymentMethod
}

func (store *DataStore) Exist(orderID string) bool {
	if store.PaymentOrder[orderID] == "" {
		return false
	}
	return true
}
