package conversion

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SetupBech32Prefix() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("p2p", "p2ppub")
	config.SetBech32PrefixForValidator("p2pv", "p2pvpub")
	config.SetBech32PrefixForConsensusNode("p2pc", "p2pcpub")
}
