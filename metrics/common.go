/*
 * Zenoss CONFIDENTIAL
 * __________________
 *
 *  This software Copyright (c) Zenoss, Inc. 2019
 *  All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of Zenoss Incorporated
 * and its suppliers, if any.  The intellectual and technical concepts contained herein are owned
 * and proprietary to Zenoss Incorporated and its suppliers and may be covered by U.S. and Foreign
 * Patents, patents in process, and are protected by U.S. and foreign trade secret or copyright law.
 * Dissemination of this information or reproduction of any this material herein is strictly forbidden
 * unless prior written permission by an authorized officer is obtained from Zenoss Incorporated.
 */
package metrics

import (
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	KeyWorker, _ = tag.NewKey("worker")

	DefaultBatchSizeDistribution    = view.Distribution(1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536)
	DefaultMinutesDistribution      = view.Distribution(1, 5, 10, 15, 20, 30, 45, 60, 120, 180, 240, 360, 480, 720, 1440, 2880, 4320, 10080, 20160, 43200)
	DefaultHoursDistribution        = view.Distribution(1, 2, 3, 6, 9, 12, 15, 18, 24, 36, 48, 60, 72, 96, 120, 168, 360, 720, 1440, 2160, 4320, 8760, 17520, 35040)
	Log10Distribution               = view.Distribution(1, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9)
	DefaultMillisecondsDistribution = view.Distribution(1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000)
)
