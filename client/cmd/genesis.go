package cmd

import (
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/types"
	cmttime "github.com/cometbft/cometbft/types/time"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/piplabs/story/client/genutil"
	"github.com/piplabs/story/lib/netconf"
)

func MakeGenesis(network netconf.ID, valPubKeys ...crypto.PubKey) (*types.GenesisDoc, error) {
	power := cosmostypes.DefaultPowerReduction // Use any non-zero power for this single validator.

	cometVals := make([]types.GenesisValidator, 0, len(valPubKeys))
	for _, pubkey := range valPubKeys {
		cometVals = append(cometVals, types.GenesisValidator{
			Address: pubkey.Address(),
			PubKey:  pubkey,
			Power:   power.Int64(),
		})
	}

	return &types.GenesisDoc{
		ChainID:         network.Static().StoryConsensusChainIDStr(),
		GenesisTime:     cmttime.Now(),
		ConsensusParams: genutil.DefaultConsensusParams(),
		Validators:      cometVals,
		AppState:        []byte("{}"),
	}, nil
}
