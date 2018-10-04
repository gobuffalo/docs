package actions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLatestVersion(t *testing.T) {
	r := require.New(t)

	r.NotZero(buffaloVersion)
	r.NotEqual(buffaloVersion, "unknown")
}
