import re
import json
filename='./sources/贝森虚拟服务器IP地址.log'

ips = []  # 所有结果写入的数组
file = open(filename, encoding='utf-8')
constent = file.read()
file.close() # 文件用完, 立即关闭, 防止内存泄漏

# 正则表达式, 提前编译, 提升效率
pattern = re.compile('^slb virtual-server \S+ (\d+\.\d+\.\d+\.\d+).*?((?=port)[^!]*)!$', re.M | re.S)
subPattern = re.compile('port (\d+).*?service-group (\S+)', re.M | re.S)

matchs = pattern.findall(constent) # 查找所有的 slb 到 ! 的内容

# print(matchs)

for match in matchs: # 遍历所有匹配结果
	ip = {'ip': match[0], 'service-groups':[]} # 生成一个结果 字典, service-groups 数组用于存放服务组信息
	subMatchs = subPattern.findall(match[1]) # 匹配 port 和 service-group
	for subMatch in subMatchs:
		ip['service-groups'].append({  # 在结果字典的service-groups 里插入port 和service-group
			'port': subMatch[0],
			'service-group': subMatch[1], 
		})
	ips.append(ip) # 把ip 插入ips

resultStr = json.dumps(ips, sort_keys=True, indent=2) # 把结果转为json字符串, 便于查看

resultFile = open('./result.txt','w')
resultFile.write(resultStr)