package hoster_test

import (
	"hoster/hoster"
	"testing"
)

func TestDeactivationHandler_Handle(t *testing.T) {
	subject := hoster.NewDeactivationHandler()

	t.Run("Smoke test", func(t *testing.T) {
		subject.Handle(fakeCliContext(t, []string{"test"}))
	})
}
