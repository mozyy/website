package model

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"go.yyue.dev/common/message"
	"go.yyue.dev/common/utils"
	"go.yyue.dev/user/proto"
	"go.yyue.dev/user/store"
)

type errLog struct {
	msg string
	err error
}

func (e errLog) Error() string {
	info := fmt.Sprintf("[[error]] msg: %s, err: %v\n", e.msg, e.err)
	log.Printf(info)
	return info
}

type User struct {
}

func (user *User) Register(ctx context.Context, req *proto.UserInfo, msg *message.Message) error {
	if req.GetMobile() == "" || req.GetPassword() == "" {
		return errLog{"参数不合法", nil}
	}
	if err := store.DB.Ping(); err != nil {
		return errLog{"数据库连接失败", err}
	}
	if u, err := store.GetByMobile(req.GetMobile()); err != nil || u.GetMobile() == req.GetMobile() {
		return errLog{"手机号己注册", err}
	}
	uid := uuid.New()
	// TODO: 可能会死循环
	// // 防止生成的 id 重复
	// for {
	// 	if _, err := store.Get(uid.ID()); err != nil {
	// 		break
	// 	}
	// 	uid = uuid.New()
	// }
	req.Uuid = uid.String()
	req.Id = uid.ID()
	h := md5.Sum([]byte(req.Uuid + req.GetPassword()))
	req.Password = fmt.Sprintf("%x", h)
	req.State = 1
	now := time.Now().Format(time.RFC3339)
	req.CreateTime = now
	req.UpdateTime = now
	err := store.Add(*req)
	if err == nil {
		log.Printf("register sucess: %s\n", req.GetMobile())
	}
	return err
}

func (user *User) Login(context.Context, *proto.LoginRequest, *proto.UserInfo) error {
	return nil
}
func (user *User) GetInfo(context.Context, *proto.GetInfoRequest, *proto.UserInfo) error {
	return nil
}

func RegisterService() {

	srv := micro.NewService(
		micro.Name("user"),
		micro.Flags(cli.StringFlag{
			Name:  "wait",
			Usage: "wait for some url",
		}),
	)
	srv.Init(
		micro.Action(func(c *cli.Context) {
			if wait := c.String("wait"); wait != "" {
				utils.WaitForIt(wait)
			}
			store.ConnectDB()
		}),
	)
	proto.RegisterUserServiceHandler(srv.Server(), &User{})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
