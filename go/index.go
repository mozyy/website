package main

import (
	"database/sql"
	"fmt"

	"github.com/mozyy/website/go/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.99.100:3306)/user?charset=utf8")
	utils.CheckErr(err)

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	utils.CheckErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	utils.CheckErr(err)

	id, err := res.LastInsertId()
	utils.CheckErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	utils.CheckErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	utils.CheckErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		utils.CheckErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	utils.CheckErr(err)

	res, err = stmt.Exec(id)
	utils.CheckErr(err)

	affect, err = res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect)

	db.Close()
}
