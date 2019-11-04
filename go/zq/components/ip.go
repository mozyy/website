package components

import (
	"bufio"
	"log"
	"os"
	"strings"

	"yyue.dev/common/utils"
)

const res = "./resources/"

type Port struct {
	Port  string
	Group string
}

type Ip struct {
	Ip    string
	Ports []Port
}

func GetIps() []Ip {
	path := res + "贝森虚拟服务器IP地址.log"
	f, err := os.Open(path)
	utils.PanicErr(err)
	defer f.Close()
	scan := bufio.NewScanner(f)
	ips := []Ip{}
	for scan.Scan() {
		text := scan.Text()
		// fmt.Println(text)
		if strings.HasPrefix(text, "slb") {
			ip := strings.Split(text, " ")[3]
			item := Ip{
				Ip: ip,
			}
			for scan.Scan() {
				content := strings.TrimSpace(scan.Text())
				if content == "!" {
					break
				}
				port := Port{}
				if strings.HasPrefix(content, "port") {
					port.Port = strings.Split(content, " ")[1]
					for scan.Scan() {
						content := strings.TrimSpace(scan.Text())
						if strings.HasPrefix(content, "service-group") {
							port.Group = strings.Split(content, " ")[1]
							item.Ports = append(item.Ports, port)
							break
						}
					}
				}
			}
			ips = append(ips, item)
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	return ips
}
