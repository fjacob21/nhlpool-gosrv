package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSqliteStore(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
}
