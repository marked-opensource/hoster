package hoster_test

import (
	"github.com/stretchr/testify/assert"
	"hoster/hoster"
	"testing"
)

func TestActivationHandler_Handle(t *testing.T) {
	subject := hoster.NewActivationHandler()

	t.Run("Smoke test", func(t *testing.T) {
		output := captureOutput(t, func() { subject.Handle(fakeCliContext(t, []string{"test"})) })
		assert.Equal(t, "TODO activate test\n", output)
	})
}
