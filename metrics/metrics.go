// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// PanicCounter measures the count of panics.
	PanicCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "tidb",
			Name:      "panic",
			Help:      "Counter of panic.",
		}, []string{LabelSession, LabelDomain, LabelDDL})
)

// metrics labels.
const (
	LabelSession  = "session"
	LabelDomain   = "domain"
	LabelDDL      = "ddl"
	LabelGCWorker = "gcworker"

	opSucc   = "op_succ"
	opFailed = "op_failed"
)

// RetLabel returns "op_succ" when err == nil and "op_failed" when err != nil.
// This could be useful when you need to observe the operation result.
func RetLabel(err error) string {
	if err == nil {
		return opSucc
	}
	return opFailed
}

func init() {
	prometheus.MustRegister(PanicCounter)
}
