package components

import (
	"bufio"
	"log"
	"os"
	"strings"

	"go.yyue.dev/common/utils"
)

type Member struct {
	Member string
	Port   string
}
type Group struct {
	Group   string
	Members []Member
}

func GetGroups() []Group {
	path := res + "贝森虚拟服务组.txt"
	f, err := os.Open(path)
	utils.PanicErr(err)
	defer f.Close()
	scan := bufio.NewScanner(f)
	groups := []Group{}
	for scan.Scan() {
		text := scan.Text()
		// fmt.Println(text)
		if strings.HasPrefix(text, "slb") {
			title := strings.Split(text, " ")
			item := Group{
				Group: title[2],
			}
			for scan.Scan() {
				content := strings.TrimSpace(scan.Text())
				if content == "!" {
					break
				}
				if strings.HasPrefix(content, "member") {
					info := strings.Split(strings.Split(content, " ")[1], ":")
					member := Member{
						Member: info[0],
						Port:   info[1],
					}
					item.Members = append(item.Members, member)
				}
			}
			groups = append(groups, item)
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	return groups
}
