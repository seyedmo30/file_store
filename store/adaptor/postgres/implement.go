package postgres

import (
	"context"
	"store/pkg/logs"
)

func CreateMetadata(ctx context.Context) error {
	db := GetDB()

	// Example Query 1: Select all rows from a table
	rows1, err := db.Query("SELECT name , type FROM store_information")
	if err != nil {
		logs.Connect().Fatal(err.Error())
	}
	defer rows1.Close()

	// Process the query results
	for rows1.Next() {
		var column1, column2 string
		if err := rows1.Scan(&column1, &column2); err != nil {
			logs.Connect().Fatal(err.Error())
		}
		logs.Connect().Debug(column1 + column2)
	}

	if err := rows1.Err(); err != nil {
		logs.Connect().Fatal(err.Error())
	}

	// // Example Query 2: Select rows with a condition
	// rows2, err := db.Query("SELECT * FROM your_table WHERE some_column = $1", "some_value")
	// if err != nil {
	// 	logs.Connect().Fatal(err.Error())
	// }
	// defer rows2.Close()

	// // Process the query results
	// for rows2.Next() {
	// 	var column1, column2 string
	// 	if err := rows2.Scan(&column1, &column2); err != nil {
	// 		logs.Connect().Fatal(err.Error())
	// 	}
	// 	logs.Connect().Debug(column1 + column2)
	// }

	// if err := rows2.Err(); err != nil {
	// 	logs.Connect().Fatal(err.Error())
	// }

	// // Example Query 3: Insert data into a table
	// _, err = db.Exec("INSERT INTO your_table (column1, column2) VALUES ($1, $2)", "value1", "value2")
	// if err != nil {
	// 	logs.Connect().Fatal(err.Error())
	// }

	// // Example Query 4: Update data in a table
	// _, err = db.Exec("UPDATE your_table SET column1 = $1 WHERE column2 = $2", "new_value", "old_value")
	// if err != nil {
	// 	logs.Connect().Fatal(err.Error())
	// }

	// // Example Query 5: Delete data from a table
	// _, err = db.Exec("DELETE FROM your_table WHERE column1 = $1", "value_to_delete")
	// if err != nil {
	// 	logs.Connect().Fatal(err.Error())
	// }
	return nil
}


func RetrieveStore(ctx context.Context , )