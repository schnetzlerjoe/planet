package keeper

import (
	"github.com/schnetzlerjoe/planet/x/planet/types"
)

var _ types.QueryServer = Keeper{}
