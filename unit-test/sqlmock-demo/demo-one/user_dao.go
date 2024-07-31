package main

import (
	"database/sql"
)

// 被测试的函数，从数据库中获取用户信息
func getUserInfo(db *sql.DB, userId int) (string, error) {
	var name string
	err := db.QueryRow("SELECT name FROM users WHERE id = $1", userId).Scan(&name)
	return name, err
}

// 被测试的插入数据函数
func insertUserInfo(db *sql.DB, name string, age int) error {
	_, err := db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", name, age)
	return err
}

// 官方示例
func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		return
	}
	return
}
