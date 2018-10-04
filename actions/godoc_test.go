package actions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_indexGodocs(t *testing.T) {
	r := require.New(t)
	r.NoError(indexGodocs())
	r.False(true)
}
