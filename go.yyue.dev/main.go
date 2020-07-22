package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
			panic(e)
	}
}
var reg = regexp.MustCompile(`\d+$`)

func main() {
	file,err:=os.Open("proto/beike/beike.proto")
	check(err)
	f, err := os.Create("proto/beike/beike2.proto")
	check(err)
	defer f.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		b := scan.Bytes()
		r := reg.ReplaceAllFunc(b, func(d []byte) []byte {
			fmt.Printf("2text: %s, match: %s, add: %s\n", b, d, []byte{d[0]+1})
			i, err := strconv.Atoi(string(d))
			check(err)
			return []byte(strconv.Itoa(i+1)+";")
		})
		_, err := f.Write(append(r, '\n'))
		check(err)

		// if len(matchs) == 0 {
		// 	// fmt.Printf("2text: %s, %s", b, matchs)
		// 	_, err := f.Write(append(b, '\n'))
		// 	check(err)
		// } else {
		// 	fmt.Printf("text: %s, %s\n", b, matchs)
		// 	_, err := f.Write()
		// 	check(err)
		// }
	}
	err = f.Sync()
	check(err)
}
