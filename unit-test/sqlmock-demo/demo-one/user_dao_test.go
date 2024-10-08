package main

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	// 创建数据库连接模拟对象
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// 预期的查询行为和结果
	expectedId := 1
	expectedName := "John Doe"
	rows := sqlmock.NewRows([]string{"name"}).AddRow(expectedName)
	mock.ExpectQuery("^SELECT name FROM users WHERE id = \\$1").WithArgs(expectedId).WillReturnRows(rows)

	// 调用被测试的函数
	name, err := getUserInfo(db, expectedId)
	if err != nil {
		t.Errorf("error occurred: %v", err)
	}

	// 验证结果
	if name != expectedName {
		t.Errorf("expected name %s, but got %s", expectedName, name)
	}
	// 确保所有预期的操作都已完成
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestInsertUserInfo(t *testing.T) {
	// 创建模拟数据库连接
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// 预期的插入操作
	expectedName := "Alice"
	expectedAge := 25
	mock.ExpectExec("INSERT INTO users").
		WithArgs(expectedName, expectedAge).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// 调用被测试的插入函数
	err = insertUserInfo(db, expectedName, expectedAge)
	if err != nil {
		t.Errorf("error occurred during insertion: %v", err)
	}

	// 检查预期是否满足
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// a successful case
func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if err = recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// a failing test case
func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	// now we execute our method
	if err = recordStats(db, 2, 3); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
