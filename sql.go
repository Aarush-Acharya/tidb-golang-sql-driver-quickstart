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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

const (
	CreateLogSQL                 = "INSERT INTO logs (level, message, resourceId, timestamp, traceId, spanId, commit, parentResourceId) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	CreateLogsTableSQL           = "CREATE TABLE logs (`level` VARCHAR(100), `message` VARCHAR(100), `resourceId` VARCHAR(100), `timestamp` VARCHAR(100), `traceId` VARCHAR(100), `spanId` VARCHAR(100), `commit` VARCHAR(100), `parentResourceId` VARCHAR(100));"
	FetchFeatureDistinctValues   = "SELECT DISTINCT ? FROM logs"
	FetchAllLogs                 = "SELECT * FROM logs"
	DropLogsTableSQL             = "DROP TABLE IF EXISTS logs"
	// CreatePlayerSQL        = "INSERT INTO player (id, coins, goods) VALUES (?, ?, ?)"
	// GetPlayerSQL           = "SELECT id, coins, goods FROM player WHERE id = ?"
	// GetCountSQL            = "SELECT count(*) FROM player"
	// GetPlayerWithLockSQL   = GetPlayerSQL + " FOR UPDATE"
	// UpdatePlayerSQL        = "UPDATE player set goods = goods + ?, coins = coins + ? WHERE id = ?"
	// GetPlayerByLimitSQL    = "SELECT id, coins, goods FROM player LIMIT ?"
	// DropTableSQL           = "DROP TABLE IF EXISTS player"
	// CreateTableSQL         = "CREATE TABLE player ( `id` VARCHAR(36), `coins` INTEGER, `goods` INTEGER, PRIMARY KEY (`id`) );"
)
