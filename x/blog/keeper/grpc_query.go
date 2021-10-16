package keeper

import (
	"github.com/schnetzlerjoe/planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
