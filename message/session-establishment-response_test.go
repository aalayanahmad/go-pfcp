// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"net"
	"testing"
	"time"

	"github.com/aalayanahmad/go-pfcp/ie"
	"github.com/aalayanahmad/go-pfcp/message"

	"github.com/aalayanahmad/go-pfcp/internal/testutil"
)

func TestSessionEstablishmentResponse(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Single IE",
			Structured: message.NewSessionEstablishmentResponse(
				mp, fo, seid, seq, pri,
				ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
				ie.NewCause(ie.CauseRequestAccepted),
				ie.NewOffendingIE(ie.Cause),
				ie.NewFSEID(0x1111111122222222, net.ParseIP("127.0.0.1"), nil),
				ie.NewCreatedPDR(
					ie.NewPDRID(0xffff),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewUEIPAddress(0x02, "127.0.0.1", "", 0, 0),
				),
				ie.NewLoadControlInformation(ie.NewSequenceNumber(0xffffffff), ie.NewMetric(0x01)),
				ie.NewOverloadControlInformation(
					ie.NewSequenceNumber(0xffffffff),
					ie.NewMetric(0x01),
					ie.NewTimer(20*time.Hour),
					ie.NewOCIFlags(0x01),
				),
				ie.NewFQCSID("127.0.0.1", 1),
				ie.NewFailedRuleID(ie.RuleIDTypePDR, 0xffff),
				ie.NewCreatedTrafficEndpoint(
					ie.NewTrafficEndpointID(0x01),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewUEIPAddress(0x02, "127.0.0.1", "", 0, 0),
				),
				ie.NewCreatedBridgeInfoForTSC(
					ie.NewDSTTPortNumber(0xffffffff),
					ie.NewNWTTPortNumber(0xffffffff),
					ie.NewTSNBridgeID(mac1),
				),
				ie.NewProvideATSSSControlInformation(
					ie.NewMPTCPControlInformation(1),
					ie.NewATSSSLLControlInformation(1),
					ie.NewPMFControlInformation(1),
				),
			),
			Serialized: []byte{
				0x21, 0x33, 0x01, 0x12, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x3c, 0x00, 0x1d, 0x02, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x03, 0x65, 0x70, 0x63, 0x0b, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x03, 0x6f, 0x72, 0x67,
				0x00, 0x13, 0x00, 0x01, 0x01,
				0x00, 0x28, 0x00, 0x02, 0x00, 0x13,
				0x00, 0x39, 0x00, 0x0d, 0x02, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x08, 0x00, 0x29,
				0x00, 0x38, 0x00, 0x02, 0xff, 0xff,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x5d, 0x00, 0x05, 0x02, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x33, 0x00, 0x0d, 0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff, 0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x36, 0x00, 0x17,
				0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x37, 0x00, 0x01, 0x82,
				0x00, 0x6e, 0x00, 0x01, 0x01,
				0x00, 0x41, 0x00, 0x07, 0x01, 0x7f, 0x00, 0x00, 0x01, 0x00, 0x01,
				0x00, 0x72, 0x00, 0x03, 0x00, 0xff, 0xff,
				0x00, 0x80, 0x00, 0x28,
				0x00, 0x83, 0x00, 0x01, 0x01,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x5d, 0x00, 0x05, 0x02, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0xc3, 0x00, 0x1b,
				0x00, 0xc4, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xc5, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xc6, 0x00, 0x07, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01,
				0x00, 0xdc, 0x00, 0x0f,
				0x00, 0xde, 0x00, 0x01, 0x01,
				0x00, 0xdf, 0x00, 0x01, 0x01,
				0x00, 0xe0, 0x00, 0x01, 0x01,
			},
		}, {
			Description: "Multiple IEs",
			Structured: message.NewSessionEstablishmentResponse(
				mp, fo, seid, seq, pri,
				ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
				ie.NewCause(ie.CauseRequestAccepted),
				ie.NewOffendingIE(ie.Cause),
				ie.NewFSEID(0x1111111122222222, net.ParseIP("127.0.0.1"), nil),
				ie.NewCreatedPDR(
					ie.NewPDRID(0xffff),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewUEIPAddress(0x02, "127.0.0.1", "", 0, 0),
				),
				ie.NewCreatedPDR(
					ie.NewPDRID(0xeeee),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewUEIPAddress(0x02, "127.0.0.1", "", 0, 0),
				),
				ie.NewLoadControlInformation(ie.NewSequenceNumber(0xffffffff), ie.NewMetric(0x01)),
				ie.NewOverloadControlInformation(
					ie.NewSequenceNumber(0xffffffff),
					ie.NewMetric(0x01),
					ie.NewTimer(20*time.Hour),
					ie.NewOCIFlags(0x01),
				),
				ie.NewFQCSID("127.0.0.1", 1),
				ie.NewFailedRuleID(ie.RuleIDTypePDR, 0xffff),
				ie.NewCreatedTrafficEndpoint(
					ie.NewTrafficEndpointID(0x01),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewUEIPAddress(0x02, "127.0.0.1", "", 0, 0),
				),
				ie.NewCreatedBridgeInfoForTSC(
					ie.NewDSTTPortNumber(0xffffffff),
					ie.NewNWTTPortNumber(0xffffffff),
					ie.NewTSNBridgeID(tsnBridgeMac),
				),
				ie.NewProvideATSSSControlInformation(
					ie.NewMPTCPControlInformation(1),
					ie.NewATSSSLLControlInformation(1),
					ie.NewPMFControlInformation(1),
				),
			),
			Serialized: []byte{
				0x21, 0x33, 0x01, 0x41, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x3c, 0x00, 0x1d, 0x02, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x03, 0x65, 0x70, 0x63, 0x0b, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x03, 0x6f, 0x72, 0x67,
				0x00, 0x13, 0x00, 0x01, 0x01,
				0x00, 0x28, 0x00, 0x02, 0x00, 0x13,
				0x00, 0x39, 0x00, 0x0d, 0x02, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x08, 0x00, 0x29,
				0x00, 0x38, 0x00, 0x02, 0xff, 0xff,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x5d, 0x00, 0x05, 0x02, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x08, 0x00, 0x29,
				0x00, 0x38, 0x00, 0x02, 0xee, 0xee,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x5d, 0x00, 0x05, 0x02, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x33, 0x00, 0x0d, 0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff, 0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x36, 0x00, 0x17,
				0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x37, 0x00, 0x01, 0x82,
				0x00, 0x6e, 0x00, 0x01, 0x01,
				0x00, 0x41, 0x00, 0x07, 0x01, 0x7f, 0x00, 0x00, 0x01, 0x00, 0x01,
				0x00, 0x72, 0x00, 0x03, 0x00, 0xff, 0xff,
				0x00, 0x80, 0x00, 0x28,
				0x00, 0x83, 0x00, 0x01, 0x01,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x15, 0x00, 0x09, 0x01, 0x11, 0x11, 0x11, 0x11, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0x5d, 0x00, 0x05, 0x02, 0x7f, 0x00, 0x00, 0x01,
				0x00, 0xc3, 0x00, 0x1d,
				0x00, 0xc4, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xc5, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xc6, 0x00, 0x09, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x05, 0x67, 0x89,
				0x00, 0xdc, 0x00, 0x0f,
				0x00, 0xde, 0x00, 0x01, 0x01,
				0x00, 0xdf, 0x00, 0x01, 0x01,
				0x00, 0xe0, 0x00, 0x01, 0x01,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseSessionEstablishmentResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
