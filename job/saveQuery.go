package job

func QueryBlocks() {
	// rows, err := db.Queryx("SELECT * FROM blocks;")
	// if err != nil {
	// 	fmt.Println("Query Fail")
	// } else {
	// 	blocks := make([]model.Block, 0)
	// 	for rows.Next() {
	// 		block := model.Block{}
	// 		err = rows.StructScan(&block)
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		blocks = append(blocks, block)
	// 	}

	// 	fmt.Println("DB Response", blocks[0].Number)
	// }
}
