// Copyright 2022 PingCAP, Inc.
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

package tp

import "github.com/pingcap/tiflow/cdc/model"

var _ schedule = &balancer{}

type balancer struct{}

func newBalancer() *balancer {
	return nil
}

func (b *balancer) Name() string {
	return "balancer"
}

func (b *balancer) Schedule(
	currentTables []model.TableID,
	captures map[model.CaptureID]*model.CaptureInfo,
	captureTables map[model.CaptureID]captureStatus,
) []*scheduleTask {
	return nil
}