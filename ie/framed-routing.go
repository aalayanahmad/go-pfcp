// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

// Framed-Routing definitions.
//
// Ref: https://tools.ietf.org/html/rfc2865#section-5.10
const (
	FramedRoutingNone                    uint32 = 0
	FramedRoutingSendRoutingPackets      uint32 = 1
	FramedRoutingListenForRoutingPackets uint32 = 2
	FramedRoutingSendAndListen           uint32 = 3
)

// NewFramedRouting creates a new FramedRouting IE.
func NewFramedRouting(routing uint32) *IE {
	return newUint32ValIE(FramedRouting, routing)
}

// FramedRouting returns FramedRouting in uint32 if the type of IE matches.
func (i *IE) FramedRouting() (uint32, error) {
	switch i.Type {
	case FramedRouting:
		return i.ValueAsUint32()
	case CreateTrafficEndpoint:
		ies, err := i.CreateTrafficEndpoint()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == FramedRouting {
				return x.FramedRouting()
			}
		}
		return 0, ErrIENotFound
	case UpdateTrafficEndpoint:
		ies, err := i.UpdateTrafficEndpoint()
		if err != nil {
			return 0, err
		}
		for _, x := range ies {
			if x.Type == FramedRouting {
				return x.FramedRouting()
			}
		}
		return 0, ErrIENotFound
	default:
		return 0, &InvalidTypeError{Type: i.Type}
	}
}
