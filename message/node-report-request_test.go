// Copyright 2019-2024 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"
	"time"

	"github.com/aalayanahmad/go-pfcp/ie"
	"github.com/aalayanahmad/go-pfcp/message"

	"github.com/aalayanahmad/go-pfcp/internal/testutil"
)

func TestNodeReportRequest(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Single IE",
			Structured: message.NewNodeReportRequest(seq,
				ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
				ie.NewNodeReportType(0x01),
				ie.NewUserPlanePathFailureReport(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				),
				ie.NewUserPlanePathRecoveryReport(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				),
				ie.NewClockDriftReport(
					ie.NewTSNTimeDomainNumber(255),
					ie.NewTimeOffsetThreshold(10*time.Second),
					ie.NewCumulativeRateRatioThreshold(0xffffffff),
					ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				),
				ie.NewGTPUPathQoSReport(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
					ie.NewGTPUPathInterfaceType(1, 1),
					ie.NewQoSReportTrigger(1, 1, 1),
					ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewQoSInformationInGTPUPathQoSReport(
						ie.NewAveragePacketDelay(10*time.Second),
						ie.NewMinimumPacketDelay(10*time.Second),
						ie.NewMaximumPacketDelay(10*time.Second),
						ie.NewTransportLevelMarking(0x1111),
					),
				),
			),
			Serialized: []byte{
				0x20, 0x0c, 0x01, 0x00, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x3c, 0x00, 0x1d, 0x02, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x03, 0x65, 0x70, 0x63, 0x0b, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x03, 0x6f, 0x72, 0x67,
				0x00, 0x65, 0x00, 0x01, 0x01,
				0x00, 0x66, 0x00, 0x23,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xbb, 0x00, 0x23,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xcd, 0x00, 0x21,
				0x00, 0xce, 0x00, 0x01, 0xff,
				0x00, 0xcf, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02, 0x54, 0x0b, 0xe4, 0x00,
				0x00, 0xd0, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xef, 0x00, 0x5f,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xf1, 0x00, 0x01, 0x03,
				0x00, 0xed, 0x00, 0x01, 0x07,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xf0, 0x00, 0x1e,
				0x00, 0xea, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xeb, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xec, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0x1e, 0x00, 0x02, 0x11, 0x11,
			},
		}, {
			Description: "Multiple IEs",
			Structured: message.NewNodeReportRequest(seq,
				ie.NewNodeID("", "", "go-pfcp.epc.3gppnetwork.org"),
				ie.NewNodeReportType(0x01),
				ie.NewUserPlanePathFailureReport(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				),
				ie.NewUserPlanePathRecoveryReport(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				),
				ie.NewClockDriftReport(
					ie.NewTSNTimeDomainNumber(255),
					ie.NewTimeOffsetThreshold(10*time.Second),
					ie.NewCumulativeRateRatioThreshold(0xffffffff),
					ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				),
				ie.NewClockDriftReport(
					ie.NewTSNTimeDomainNumber(1),
					ie.NewTimeOffsetThreshold(10*time.Second),
					ie.NewCumulativeRateRatioThreshold(0xffffffff),
					ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				),
				ie.NewGTPUPathQoSReport(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
					ie.NewGTPUPathInterfaceType(1, 1),
					ie.NewQoSReportTrigger(1, 1, 1),
					ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewQoSInformationInGTPUPathQoSReport(
						ie.NewAveragePacketDelay(10*time.Second),
						ie.NewMinimumPacketDelay(10*time.Second),
						ie.NewMaximumPacketDelay(10*time.Second),
						ie.NewTransportLevelMarking(0x1111),
					),
				),
				ie.NewGTPUPathQoSReport(
					ie.NewRemoteGTPUPeer(0x0e, "127.0.0.2", "", ie.DstInterfaceAccess, "some.instance.example"),
					ie.NewGTPUPathInterfaceType(1, 1),
					ie.NewQoSReportTrigger(1, 1, 1),
					ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewQoSInformationInGTPUPathQoSReport(
						ie.NewAveragePacketDelay(10*time.Second),
						ie.NewMinimumPacketDelay(10*time.Second),
						ie.NewMaximumPacketDelay(10*time.Second),
						ie.NewTransportLevelMarking(0x1111),
					),
				),
			),
			Serialized: []byte{
				0x20, 0x0c, 0x01, 0x88, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x3c, 0x00, 0x1d, 0x02, 0x07, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x03, 0x65, 0x70, 0x63, 0x0b, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x03, 0x6f, 0x72, 0x67,
				0x00, 0x65, 0x00, 0x01, 0x01,
				0x00, 0x66, 0x00, 0x23,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xbb, 0x00, 0x23,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xcd, 0x00, 0x21,
				0x00, 0xce, 0x00, 0x01, 0xff,
				0x00, 0xcf, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02, 0x54, 0x0b, 0xe4, 0x00,
				0x00, 0xd0, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xcd, 0x00, 0x21,
				0x00, 0xce, 0x00, 0x01, 0x01,
				0x00, 0xcf, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02, 0x54, 0x0b, 0xe4, 0x00,
				0x00, 0xd0, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xef, 0x00, 0x5f,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x01,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xf1, 0x00, 0x01, 0x03,
				0x00, 0xed, 0x00, 0x01, 0x07,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xf0, 0x00, 0x1e,
				0x00, 0xea, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xeb, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xec, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0x1e, 0x00, 0x02, 0x11, 0x11,
				0x00, 0xef, 0x00, 0x5f,
				0x00, 0x67, 0x00, 0x1f,
				0x0e,
				0x7f, 0x00, 0x00, 0x02,
				0x00, 0x01, 0x00,
				0x00, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				0x00, 0xf1, 0x00, 0x01, 0x03,
				0x00, 0xed, 0x00, 0x01, 0x07,
				0x00, 0x9c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0xf0, 0x00, 0x1e,
				0x00, 0xea, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xeb, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0xec, 0x00, 0x04, 0x00, 0x00, 0x27, 0x10,
				0x00, 0x1e, 0x00, 0x02, 0x11, 0x11,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseNodeReportRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
