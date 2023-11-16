package mocks_test

import (
	"mocks/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapters(t *testing.T) {
	mockExt := mocks.NewExternalAdapter(t)
	mockInt := mocks.NewInternalAdapter(t)

	mockExt.On("Do", "ping").Return("pong", nil)
	mockInt.On("Do").Return(true, nil)

	assert.True(t, mockExt.AssertNotCalled(t, "Do", "ping"))
}
