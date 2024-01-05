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
import "github.com/gin-gonic/gin"


import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Configure the example database connection.
	openDB("mysql", func(db *sql.DB) {
		r := gin.Default()

		r.GET("/fetchAll", func(c *gin.Context) {
			allLogs, err := queryLogs(db, FetchAllLogs)
			if err != nil {
				panic(err)
			}
			for index, log := range allLogs {
				fmt.Printf("print %d log: %+v\n", index+1, log)
			}
			c.JSON(200, allLogs)
		})


		r.POST("/getOptions", func(c *gin.Context) {
			var featureMap map[string]string
			c.BindJSON(&featureMap)
			var feature string 
			feature = featureMap["feature"]
			options, err := getOptions(db, feature)
			if err != nil {
				panic(err)
			}
			c.JSON(200, options)
		})

		r.POST("/createLog", func(c *gin.Context) {
			var recievedlog Log
			c.BindJSON(&recievedlog)
		    createLog(db, recievedlog)
			fmt.Printf("Log created: %v\n", recievedlog)
		})


	
		r.POST("/queryLogs", func(c *gin.Context) {
			var features map[string]string
			c.BindJSON(&features)
			query := "SELECT * FROM logs WHERE"
			for key, value := range features {
				query += fmt.Sprintf(" %s = '%s' AND", key, value)
			}
			// Remove the last 'AND' from the query string
			query = query[:len(query)-4]
			fmt.Println("Generated query:", query)

			FetchedLogs, err := queryLogs(db, query)
			if err != nil {
				panic(err)
			}
			c.JSON(200, FetchedLogs)
		})
	
		r.Run(":3000")
	})
}

// func simpleExample(db *sql.DB) {
// 	// Create a player, who has a coin and a goods.
// 	err := createPlayer(db, Player{ID: "test", Coins: 1, Goods: 1})
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Get a player.
// 	testPlayer, err := getPlayer(db, "test")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("getPlayer: %+v\n", testPlayer)

// 	// Create players with bulk inserts. Insert 1919 players totally, with 114 players per batch.

// 	err = bulkInsertPlayers(db, randomPlayers(1919), 114)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Count players amount.
// 	playersCount, err := getCount(db)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("countPlayers: %d\n", playersCount)

// 	// Print 3 players.
// 	threePlayers, err := getPlayerByLimit(db, 3)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for index, player := range threePlayers {
// 		fmt.Printf("print %d player: %+v\n", index+1, player)
// 	}
// }

// func tradeExample(db *sql.DB) {
// 	// Player 1: id is "1", has only 100 coins.
// 	// Player 2: id is "2", has 114514 coins, and 20 goods.
// 	player1 := Player{ID: "1", Coins: 100}
// 	player2 := Player{ID: "2", Coins: 114514, Goods: 20}

// 	// Create two players "by hand", using the INSERT statement on the backend.
// 	if err := createPlayer(db, player1); err != nil {
// 		panic(err)
// 	}
// 	if err := createPlayer(db, player2); err != nil {
// 		panic(err)
// 	}

// 	// Player 1 wants to buy 10 goods from player 2.
// 	// It will cost 500 coins, but player 1 cannot afford it.
// 	fmt.Println("\nbuyGoods:\n    => this trade will fail")
// 	if err := buyGoods(db, player2.ID, player1.ID, 10, 500); err == nil {
// 		panic("there shouldn't be success")
// 	}

// 	// So player 1 has to reduce the incoming quantity to two.
// 	fmt.Println("\nbuyGoods:\n    => this trade will success")
// 	if err := buyGoods(db, player2.ID, player1.ID, 2, 100); err != nil {
// 		panic(err)
// 	}
// }

func openDB(driverName string, runnable func(db *sql.DB)) {
	db, err := sql.Open(driverName, getDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	runnable(db)
}

func getDSN() string {
	tidbHost := getEnvWithDefault("TIDB_HOST", "127.0.0.1")
	tidbPort := getEnvWithDefault("TIDB_PORT", "4000")
	tidbUser := getEnvWithDefault("TIDB_USER", "root")
	tidbPassword := getEnvWithDefault("TIDB_PASSWORD", "")
	tidbDBName := getEnvWithDefault("TIDB_DB_NAME", "test")
	useSSL := getEnvWithDefault("USE_SSL", "false")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s",
		tidbUser, tidbPassword, tidbHost, tidbPort, tidbDBName, useSSL)
}

func getEnvWithDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
