package main

import (
	"github.com/tealeg/xlsx"
	"go.yyue.dev/common/utils"
	"go.yyue.dev/zq/components"
)

// Server is server struct
type Server struct {
	Name string `json:"服务组启用服务器"`
	IP   string `json:"服务器IP地址"`
	Port string `json:"服务器对内端口"`
}

// Group is server group
type Group struct {
	Name    string `json:"虚拟服务器启用服务组"`
	Port    string `json:"服务组启用端口"`
	Servers []Server
}

// Result is result
type Result struct {
	IP     string `json:"虚拟服务器IP地址"`
	Groups []Group
}

func main() {
	ips := components.GetIps()
	groups := components.GetGroups()
	servers := components.GetServers()

	result := []Result{}
	for _, ip := range ips {
		item := Result{}
		item.IP = ip.IP
		for _, ipPorts := range ip.Ports {
			group := Group{}
			group.Name = ipPorts.Group
			group.Port = ipPorts.Port
			for _, groupItem := range groups {
				if groupItem.Group == ipPorts.Group {
					for _, groupMember := range groupItem.Members {
						server := Server{}
						server.Name = groupMember.Member
						server.Port = groupMember.Port
						for _, serverItem := range servers {
							if serverItem.Member == groupMember.Member {
								server.IP = serverItem.Server
								break
							}
						}
						group.Servers = append(group.Servers, server)
					}
				}
			}
			item.Groups = append(item.Groups, group)
		}
		result = append(result, item)
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	utils.PanicErr(err)
	// title
	row := sheet.AddRow()
	titles := []string{
		"虚拟服务器IP地址",
		"虚拟服务器启用服务组",
		"服务组启用端口",
		"服务组启用服务器",
		"服务器IP地址",
		"服务器对内端口",
	}
	row.WriteSlice(&titles, -1)
	for _, item := range result {
		c1Len := 0
		for i, group := range item.Groups {

			c1Len += len(group.Servers)
			for j, server := range group.Servers {
				contents := []string{
					item.IP,
					group.Name,
					group.Port,
					server.Name,
					server.IP,
					server.Port,
				}
				row := sheet.AddRow()
				row.WriteSlice(&contents, -1)
				if j == 0 {
					row.Cells[1].Merge(0, len(group.Servers)-1)
					row.Cells[2].Merge(0, len(group.Servers)-1)
				}
				if i == 0 && j == 0 {
					clen := 0
					for _, group := range item.Groups {
						clen += len(group.Servers)
						row.Cells[0].Merge(0, clen-1)
					}
				}
			}
		}

	}
	err = file.Save("test.xlsx")
	utils.PanicErr(err)

}
