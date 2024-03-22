package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // 导入SQLite3驱动
)

func main() {
	// 打开数据库连接
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	createTableSQL := `CREATE TABLE IF NOT EXISTS projects (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,   
        "name" TEXT,
        "description" TEXT
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO projects(name, description) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("My Project", "This is a sample project")
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	rows, err := db.Query("SELECT id, name, description FROM projects")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, description string
		err = rows.Scan(&id, &name, &description)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, description)
	}

	// 检查查询过程中可能出现的错误
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
