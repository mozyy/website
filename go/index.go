package main

import (
	"database/sql"
	"fmt"

	"go/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.99.100:3306)/user?charset=utf8")
	utils.PanicErr(err)

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	utils.PanicErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	utils.PanicErr(err)

	id, err := res.LastInsertId()
	utils.PanicErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	utils.PanicErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	utils.PanicErr(err)

	affect, err := res.RowsAffected()
	utils.PanicErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	utils.PanicErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		utils.PanicErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	utils.PanicErr(err)

	res, err = stmt.Exec(id)
	utils.PanicErr(err)

	affect, err = res.RowsAffected()
	utils.PanicErr(err)

	fmt.Println(affect)

	db.Close()
}
