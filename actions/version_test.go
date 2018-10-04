package actions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLatestVersion(t *testing.T) {
	r := require.New(t)

	v := findLatestVersion()
	r.NotZero(v)
	r.NotEqual(v, "unknown")
}
