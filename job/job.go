// Package job - Main process: pulling, processing, and saving the blockchain data
package job

import (
	"fmt"

	"github.com/akkien/ethereumetl/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//ExportAll : Export all data from block
func ExportAll() {
	connStr := "postgres://akkien:trungkien@127.0.0.1:5432/ropsten?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} else {
		if err = db.Ping(); err != nil {
			panic(err)
		}
		fmt.Println("Connected")

		rows, err := db.Queryx("SELECT * FROM blocks;")
		if err != nil {
			fmt.Println("Query Fail")
		} else {
			blocks := make([]model.Block, 0)
			for rows.Next() {
				block := model.Block{}
				err = rows.StructScan(&block)
				if err != nil {
					panic(err)
				}

				blocks = append(blocks, block)
			}

			fmt.Println("DB Response", blocks[0].Number)
		}
	}
}
