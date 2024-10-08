// Code generated by "callbackgen -type BboMonitor"; DO NOT EDIT.

package twap

import (
	"github.com/c9s/bbgo/pkg/types"
)

func (m *BboMonitor) OnUpdate(cb func(bid types.PriceVolume, ask types.PriceVolume)) {
	m.updateCallbacks = append(m.updateCallbacks, cb)
}

func (m *BboMonitor) EmitUpdate(bid types.PriceVolume, ask types.PriceVolume) {
	for _, cb := range m.updateCallbacks {
		cb(bid, ask)
	}
}
