package main

import (
	"github.com/mozyy/utils"
	"github.com/mozyy/website/go/datamanage"
	"github.com/mozyy/website/go/message"
)

type User struct {
	uid      int
	name     string
	phone    int
	password string
}

func (user *User) Regist(args *User, reply *message.Message) error {

	db := datamanage.GetDb()

	rows, err := db.Query("SELECT * FROM userinfo WHERE uid=?", user.uid)
	utils.PanicErr(err)

	if rows.Next() {
		return reply.Error("该用户已经注册")
	}

	stmt, err := db.Prepare("INSERT userinfo SET name=?,phone=?,password=?")
	utils.PanicErr(err)

	_, err = stmt.Exec(user.name, user.phone, user.password)
	utils.PanicErr(err)

	reply.Success("注册成功")

	return nil
}

func (user *User) Login(args *User, reply *message.Message) error {

	db := datamanage.GetDb()

	rows, err := db.Query("SELECT * FROM userinfo WHERE uid=?", user.uid)
	utils.PanicErr(err)

	if rows.Next() {
		return reply.Error("该用户已经注册")
	}

	stmt, err := db.Prepare("INSERT userinfo SET name=?,phone=?,password=?")
	utils.PanicErr(err)

	_, err = stmt.Exec(user.name, user.phone, user.password)
	utils.PanicErr(err)

	reply.Success("注册成功")

	return nil
}
