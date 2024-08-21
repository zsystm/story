package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/storyprotocol/iliad/e2e/types"
)

func TestNextRPCAddress(t *testing.T) {
	t.Parallel()
	c := types.NewPublicChain(types.EVMChain{}, []string{"1 ", " 2", "3"})

	require.Equal(t, "1", c.NextRPCAddress())
	require.Equal(t, "2", c.NextRPCAddress())
	require.Equal(t, "3", c.NextRPCAddress())
	require.Equal(t, "1", c.NextRPCAddress())
}
