package hoster_test

import (
	"github.com/stretchr/testify/assert"
	"hoster/hoster"
	"testing"
)

func TestDeactivationHandler_Handle(t *testing.T) {
	subject := hoster.NewDeactivationHandler()

	t.Run("Smoke test", func(t *testing.T) {
		output := captureOutput(t, func() { subject.Handle(fakeCliContext(t, []string{"test"})) })
		assert.Equal(t, "TODO deactivate test\n", output)
	})
}
