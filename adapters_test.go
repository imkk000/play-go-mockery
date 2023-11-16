package mocks_test

import (
	"mocks/mocks"
	"testing"
)

func TestAdapters(t *testing.T) {
	t.Cleanup(func() {
		t.Error("clean up stuff")
	})

	mockExt := mocks.NewExternalAdapter(t)
	mockInt := mocks.NewInternalAdapter(t)

	mockExt.On("Do", "ping").Return("pong", nil)
	mockInt.On("Do").Return(true, nil)
}
