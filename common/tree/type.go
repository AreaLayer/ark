package tree

import (
	"github.com/vulpemventures/go-elements/psetv2"
)

type TreeFactory func(outpoint psetv2.InputArgs) (CongestionTree, error)

type Receiver struct {
	Script VtxoScript
	Amount uint64
}
