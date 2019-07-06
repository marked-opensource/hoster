package hoster_test

import (
	"hoster/hoster"
	"testing"
)

func TestActivationHandler_Handle(t *testing.T) {
	subject := hoster.NewActivationHandler()

	t.Run("Smoke test", func(t *testing.T) {
		subject.Handle(fakeCliContext(t, []string{"test"}))
	})
}
