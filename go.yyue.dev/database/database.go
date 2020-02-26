package database

import (
	"database/sql"
	"fmt"
	"time"

	// sql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbs = make(map[string]*sql.DB)
)

// connect will connect database
func connect(database string) (*sql.DB, error) {
	// dbc := utils.GetConfig().Database
	user, password, domain, port := "root", "123456", "162.219.127.105", "3306"
	// dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?%s`, dbc.User, dbc.Password,
	// 	dbc.Domain, dbc.Port, database, "charset=utf8&parseTime=true")
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?%s`, user, password,
		domain, port, database, "charset=utf8&parseTime=true")
	fmt.Println(dsn)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//设置连接池
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(10 * time.Minute)
	return conn, conn.Ping()
}

func GetDB(database string) (*sql.DB, error) {
	if db, ok := dbs[database]; ok && db.Ping() == nil {
		return db, nil
	}

	db, err := connect(database)
	if err != nil {
		return nil, err
	}
	dbs[database] = db
	return db, nil
}
