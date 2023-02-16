// Copyright 2019-2023 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message

import (
	"github.com/wmnsk/go-pfcp/ie"
)

// SessionSetDeletionRequest is a SessionSetDeletionRequest formed PFCP Header and its IEs above.
type SessionSetDeletionRequest struct {
	*Header
	NodeID *ie.IE
	FQCSID *ie.IE
	IEs    []*ie.IE
}

// NewSessionSetDeletionRequest creates a new SessionSetDeletionRequest.
func NewSessionSetDeletionRequest(seq uint32, id, csid *ie.IE, ies ...*ie.IE) *SessionSetDeletionRequest {
	m := &SessionSetDeletionRequest{
		Header: NewHeader(
			1, 0, 0, 0,
			MsgTypeSessionSetDeletionRequest, 0, seq, 0,
			nil,
		),
		NodeID: id,
		FQCSID: csid,
		IEs:    ies,
	}
	m.SetLength()

	return m
}

// Marshal returns the byte sequence generated from a SessionSetDeletionRequest.
func (m *SessionSetDeletionRequest) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}

	return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (m *SessionSetDeletionRequest) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if i := m.NodeID; i != nil {
		if err := i.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += i.MarshalLen()
	}
	if i := m.FQCSID; i != nil {
		if err := i.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += i.MarshalLen()
	}

	for _, ie := range m.IEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	m.Header.SetLength()
	return m.Header.MarshalTo(b)
}

// ParseSessionSetDeletionRequest decodes a given byte sequence as a SessionSetDeletionRequest.
func ParseSessionSetDeletionRequest(b []byte) (*SessionSetDeletionRequest, error) {
	m := &SessionSetDeletionRequest{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes a given byte sequence as a SessionSetDeletionRequest.
func (m *SessionSetDeletionRequest) UnmarshalBinary(b []byte) error {
	var err error
	m.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(m.Header.Payload) < 2 {
		return nil
	}

	ies, err := ie.ParseMultiIEs(m.Header.Payload)
	if err != nil {
		return err
	}

	for _, i := range ies {
		switch i.Type {
		case ie.NodeID:
			m.NodeID = i
		case ie.FQCSID:
			m.FQCSID = i
		default:
			m.IEs = append(m.IEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length of Data.
func (m *SessionSetDeletionRequest) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)

	if i := m.NodeID; i != nil {
		l += i.MarshalLen()
	}
	if i := m.FQCSID; i != nil {
		l += i.MarshalLen()
	}

	for _, ie := range m.IEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}

	return l
}

// SetLength sets the length in Length field.
func (m *SessionSetDeletionRequest) SetLength() {
	m.Header.Length = uint16(m.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *SessionSetDeletionRequest) MessageTypeName() string {
	return "Session Set Deletion Request"
}

// SEID returns the SEID in uint64.
func (m *SessionSetDeletionRequest) SEID() uint64 {
	return m.Header.seid()
}
