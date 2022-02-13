package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataStore(t *testing.T) {
	myStore := NewDataStore()

	// add some
	myStore.Insert("FA84933", "BNI")
	myStore.Insert("FA84932", "PAYPAL")

	//peek
	assert.False(t, myStore.Exist("UI444"))
	assert.True(t, myStore.Exist("FA84932"))
}
