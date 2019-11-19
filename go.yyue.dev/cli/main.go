package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/c-bata/go-prompt"
	"github.com/urfave/cli/v2"
	util "go.yyue.dev/cli/utils"
	"go.yyue.dev/common/utils"
)

// Env is the environment
type Env struct {
	Name        string
	Description string
}

// Object is the object package
type Object struct {
	Name         string
	Description  string
	RelativePath string
	CmdTemp      string
}

// ObjectConfig is the object and envs
type ObjectConfig struct {
	Object
	Envs []Env
}

var (
	// WorkPath is work space path
	WorkPath = filepath.Join("d:/", "work")
	dev      = Env{Name: "dev", Description: "开发环境"}
	test     = Env{Name: "test", Description: "测试环境"}
	open     = Env{Name: "open", Description: "开放平台"}
	prod     = Env{Name: "prod", Description: "正式环境"}
)

var objects = []ObjectConfig{
	ObjectConfig{
		Object: Object{
			Name:         "video-session",
			Description:  "国君视频会议",
			CmdTemp:      "npm run build",
			RelativePath: "./video-session",
		},
		Envs: []Env{
			dev,
			test,
			open,
			prod,
		},
	},
	ObjectConfig{
		Object: Object{
			Name:        "video-session2",
			Description: "国君视频会议",
			CmdTemp:     "npm run",
		},
		Envs: []Env{
			dev,
			test,
			prod,
		},
	},
	ObjectConfig{
		Object: Object{
			Name:        "video-session3",
			Description: "国君视频会议",
			CmdTemp:     "npm run",
		},
		Envs: []Env{
			dev,
			prod,
		},
	},
}

func main2() {

	var object, env string
	commands := []*cli.Command{
		&cli.Command{
			Name:    "deploy",
			Aliases: []string{"d"},
			Usage:   "发布项目",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "object",
					Aliases:     []string{"o"},
					Usage:       "select object `NAME` to deploy",
					Destination: &object,
				},
				&cli.StringFlag{
					Name:        "env",
					Aliases:     []string{"e"},
					Usage:       "select env `ENV` to deploy",
					Destination: &env,
				},
			},
			Action: func(c *cli.Context) error {

				d := &Deploy{
					ObjectName: object,
					EnvName:    env,
				}
				prompt1 := prompt.New(d.Executor, d.Complete, prompt.OptionPrefix("> "), prompt.OptionTitle("deploy:"))
				prompt1.Run()
				return nil
			},
		},
	}
	app := &cli.App{
		Name:  "ttd",
		Usage: "ttd command line interface",
		Authors: []*cli.Author{
			&cli.Author{
				Name: "yang yue",
			},
		},
		Commands: commands,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := checkObject(objects)
	utils.PanicErr(err)

	err = app.Run(os.Args)
	utils.PanicErr(err)

}

func main() {
	err := util.UploadDeploy("d:/work/evidence-security/dist/evidence-security-one/evidence-security-one.zip", "http://192.168.88.122:9990/upfile.php?pk_id=202")
	if err != nil {
		fmt.Println(err)
	}
}

// Deploy is deploy struct
type Deploy struct {
	ObjectName string
	EnvName    string
	Object
	Env
}

// Complete is for go-prompt Complete function
func (deploy *Deploy) Complete(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	args := strings.Split(d.TextBeforeCursor(), " ")
	switch len(args) {
	case 1:
		s = getObjectSuggest(objects)
		s = append(s, prompt.Suggest{Text: "exit", Description: "退出"})
	case 2:
		object, err := getObjectConfig(args[0], objects)
		if err != nil {
			return []prompt.Suggest{}
		}
		s = getEnvSuggest(object.Envs)
	}
	return prompt.FilterContains(s, d.GetWordBeforeCursor(), true)
}

// Executor is deploy executor impl go-prompt Executor
func (deploy *Deploy) Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}
	args := strings.Split(s, " ")
	if len(args) == 2 {
		object, err := getObjectConfig(args[0], objects)
		if err != nil {
			fmt.Println("Object error")
			return
		}
		deploy.Object = object.Object
		env, err := getEnv(args[1], object.Envs)
		if err != nil {
			fmt.Println("Env error")
			return
		}
		deploy.Env = env
	}
	if deploy.Object.Name == "" || deploy.Env.Name == "" {
		fmt.Println("need object and env")
		return
	}
	t := template.Must(template.New("exec").Parse(deploy.Object.CmdTemp))
	buffer := bytes.NewBufferString("")
	err := t.Execute(buffer, deploy.Env)
	if err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
		return
	}
	cmd := exec.Command("bash", "-c", buffer.String())
	dir := filepath.Join(WorkPath, deploy.Object.RelativePath)
	fmt.Printf("Got error: %s\n", dir)
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}

func getObjectSuggest(objects []ObjectConfig) []prompt.Suggest {
	s := []prompt.Suggest{}
	for _, obj := range objects {
		item := prompt.Suggest{Text: obj.Name, Description: obj.Description}
		s = append(s, item)
	}
	return s
}

func getEnvSuggest(envs []Env) []prompt.Suggest {
	s := []prompt.Suggest{}
	for _, env := range envs {
		item := prompt.Suggest{Text: env.Name, Description: env.Description}
		s = append(s, item)
	}
	return s
}

func getObjectConfig(name string, objects []ObjectConfig) (ObjectConfig, error) {
	if name == "" {
		return ObjectConfig{}, fmt.Errorf("object name cannot be empty")
	}
	for _, obj := range objects {
		if obj.Name == name {
			return obj, nil
		}
	}
	return ObjectConfig{}, fmt.Errorf("invalid object: %s", name)
}

func getEnv(name string, envs []Env) (Env, error) {
	if name == "" {
		return Env{}, fmt.Errorf("object name cannot be empty")
	}
	for _, env := range envs {
		if env.Name == name {
			return env, nil
		}
	}
	return Env{}, fmt.Errorf("invalid object: %s", name)
}

func checkObject(objs []ObjectConfig) error {
	objMap := map[string]bool{}
	for _, obj := range objects {
		objName := obj.Name
		if objMap[objName] {
			return fmt.Errorf("multiple object config: %s", objName)
		}
		objMap[objName] = true
		envMap := map[string]bool{}
		for _, env := range obj.Envs {
			envName := env.Name
			if envMap[envName] {
				return fmt.Errorf("multiple env config: %s:%s", objName, envName)
			}
			envMap[envName] = true
		}
	}
	return nil
}
