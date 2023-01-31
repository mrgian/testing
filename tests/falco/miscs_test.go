package testfalco

import (
	"testing"

	"github.com/jasondellaluce/falco-testing/pkg/falco"
	"github.com/jasondellaluce/falco-testing/tests"
	"github.com/jasondellaluce/falco-testing/tests/data/configs"

	"github.com/stretchr/testify/assert"
)

func TestMiscs_StartupFail(t *testing.T) {
	runner := tests.NewFalcoExecutableRunner(t)
	t.Run("empty-config", func(t *testing.T) {
		res := falco.Test(runner, falco.WithConfig(configs.EmptyConfig))
		assert.NotNil(t, res.Err(), "%s", res.Stderr())
		assert.Equal(t, res.ExitCode(), 1)
		assert.Contains(t, res.Stderr(), "You must specify at least one rules file")
	})
}
