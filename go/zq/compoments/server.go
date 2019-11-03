package compoments

import (
	"bufio"
	"log"
	"os"
	"strings"

	"yyue.dev/common/utils"
)

type Server struct {
	Server string
	Member string
}

func GetServers() []Server {
	path := res + "贝森虚拟服务器.txt"
	f, err := os.Open(path)
	utils.PanicErr(err)
	defer f.Close()
	scan := bufio.NewScanner(f)
	servers := []Server{}
	for scan.Scan() {
		text := scan.Text()
		// fmt.Println(text)
		if strings.HasPrefix(text, "slb") {
			info := strings.Split(text, " ")
			item := Server{
				Member: info[2],
				Server: info[3],
			}
			servers = append(servers, item)
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	return servers
}
