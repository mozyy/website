package model

import (
	"fmt"

	"github.com/mozyy/website/go/datamanage"
	"github.com/mozyy/website/go/message"
	"github.com/mozyy/website/go/utils"
)

// CREATE TABLE `userinfo` (
// 	`uid` INT(10) NOT NULL AUTO_INCREMENT,
// 	`name` VARCHAR(64) NULL DEFAULT NULL,
// 	`phone` VARCHAR(64) NULL DEFAULT NULL,
// 	`password` VARCHAR(64) NULL DEFAULT NULL,
// 	`created` TIMESTAMP CURRENT_TIMESTAMP,
// 	`updated` TIMESTAMP NO UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`uid`)
// );

type User struct {
	UID      int
	Name     string
	Phone    int
	Password string
}

var db = datamanage.GetDb("user")

func (u *User) Regist(user User, reply *message.Message) error {

	fmt.Println(user)

	rows, err := db.Query("SELECT * FROM userinfo WHERE uid=?", user.UID)
	utils.PanicErr(err)

	if rows.Next() {
		reply.Error("该用户已经注册")
		return nil
	}

	stmt, err := db.Prepare("INSERT userinfo SET name=?,phone=?,password=?")
	utils.PanicErr(err)

	_, err = stmt.Exec(user.Name, user.Phone, user.Password)
	utils.PanicErr(err)

	reply.Success("注册成功")

	return nil
}

func (user *User) Login(args User, reply *message.Message) error {

	fmt.Println(args)

	rows, err := db.Query("SELECT * FROM userinfo WHERE uid=?", args.UID)
	utils.PanicErr(err)

	if rows.Next() {
		u := User{}
		// rows.Scan(u.UID, u.Name, u.Phone, u.Password)
		reply.Success("登录成功")
		reply.Data = u
		return nil
	}

	reply.Error("登录失败")
	return nil
}
