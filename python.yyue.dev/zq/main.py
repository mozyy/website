import re
import json
filename='./sources/贝森虚拟服务器IP地址.log'

ips = []  # 所有结果写入的数组

with open(filename, encoding='utf-8') as file:
	while True:
		line = file.readline()
		if line:
			if line.find('slb')>-1:
				ip = {'ip':line.split(' ')[3],'service-groups':[]}
				while True:
					line = file.readline().strip()
					group = {}
					if line != '!':
						if line.find('port') > -1:
							group['port'] = line.split(' ')[1]
							while True:
								line = file.readline()
								line = line.strip()
								if line.find('service-group') > -1:
									# print(line)
									group['service-group'] = line.split(' ')[1]
									ip['service-groups'].append(group)
									break
					else:
						break
				ips.append(ip)
		else:
			break

resultStr = json.dumps(ips, sort_keys=True, indent=2) # 把结果转为json字符串, 便于查看

resultFile = open('./result2.txt','w')
resultFile.write(resultStr)