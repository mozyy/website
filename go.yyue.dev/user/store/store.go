package store

import (
	"bytes"
	"database/sql"
	"fmt"
	"text/template"

	"go.yyue.dev/database"
	"go.yyue.dev/user/proto"
)

const (
	databaseName = "yyue_user"
	table        = "userinfo"
)

var (
	DB           *sql.DB
	editTemplate = template.Must(template.New("edit").Parse(fmt.Sprintf(`
	UPDATE %s
	SET {{if .Name}}name="{{.Name}}", {{end}}{{if .Mobile}}mobile={{.Mobile}}, {{end}}
	{{if .Password}}password="{{.Password}}", {{end}}{{if .Uuid}}uuid="{{.Uuid}}", {{end}}
	{{if .State}}state={{.State}}, {{end}}{{if .UpdateTime}}update_time={{.UpdateTime}}, {{end}}
	{{if .LoginTime}}login_time={{.LoginTime}}, {{end}} WHERE id={{.Id}}
	`, table)))
)

func ConnectDB() {
	var err error
	DB, err = database.GetDB(databaseName)
	if err != nil {
		panic("get user database error: " + err.Error())
	}
}

func Add(u proto.UserInfo) error {
	_, err := DB.Exec(fmt.Sprintf(`INSERT INTO %s 
	(id, name, mobile, password, uuid, state, create_time, update_time, login_time) 
	VALUES (%d, %q, %q, %q, %q, %d, %q, %q, %q);`, table, u.GetId(), u.GetName(),
		u.GetMobile(), u.GetPassword(), u.GetUuid(), u.GetState(), u.GetCreateTime(),
		u.GetUpdateTime(), u.GetLoginTime()))
	return err
}

func Edit(u proto.UserInfo) error {
	buf := &bytes.Buffer{}
	err := editTemplate.Execute(buf, u)
	if err != nil {
		return err
	}
	_, err = DB.Exec(string(buf.Bytes()))
	return err
}

func Del(ID uint32) error {
	_, err := DB.Exec(fmt.Sprintf(`UPDATE %s SET state=-1`, table))
	return err
}

func Get(ID uint32) (u proto.UserInfo, err error) {
	row := &sql.Rows{}
	row, err = DB.Query(fmt.Sprintf(`SELECT * FROM %s WHERE id = %d;`, table, ID))
	if err != nil {
		return
	}
	err = scan(row, &u)
	return
}
func GetByMobile(mobile string) (u proto.UserInfo, err error) {
	row := &sql.Rows{}
	row, err = DB.Query(fmt.Sprintf(`SELECT * FROM %s WHERE mobile = %s;`, table, mobile))
	if err != nil {
		return
	}
	if row.Next() {
		err = scan(row, &u)
	}
	return
}

func GetAll() ([]proto.UserInfo, error) {
	users := []proto.UserInfo{}
	rows, err := DB.Query(`SELECT * FROM userinfo WHERE state != -1`)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		u := proto.UserInfo{}
		err := scan(rows, &u)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, nil
}

func scan(rows *sql.Rows, u *proto.UserInfo) error {
	return rows.Scan(&u.Id, &u.Name, &u.Mobile, &u.Password, &u.Uuid, &u.State, &u.CreateTime, &u.UpdateTime, &u.LoginTime)
}
