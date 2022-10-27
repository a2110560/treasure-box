package main

import (
	"fmt"
	"math"
	"strings"
)

var Chars string = "mOK2GHtZF4T5P6dnXarflzjkVSy9LMCvIqwDeWghsN78BUExAuY3JpQRboic01"

func main() {
	//rand.Seed(time.Now().UnixNano())
	//s := rand.Intn(15) + 5

	str := encode(1230)
	fmt.Println(str)
	num := decode(str)
	fmt.Println(num)
}
func encode(num uint64) string {
	var bytes []byte
	for num > 0 {
		bytes = append(bytes, Chars[num%62])
		num = num / 62
	}
	for len(bytes) < 6 {
		if len(bytes) == 6 {
			break
		}
		bytes = append(bytes, Chars[0])
	}
	reverse(bytes)
	//var bytes2 []byte
	//for companycode > 0 {
	//	bytes2 = append(bytes2, c[companycode%62])
	//	companycode = companycode / 62
	//}
	//for len(bytes2) < checkcode {
	//	if len(bytes2) == checkcode {
	//		break
	//	}
	//	bytes2 = append(bytes2, c[0])
	//}
	//reverse(bytes2)
	return string(bytes)
}

func decode(str string) int64 {
	var num int64
	//var nums int64
	n := len(str) //9
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(Chars, str[i])
		num += int64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	////只取後3
	//for i := checkcode; i > 0; i-- {
	//	pos := strings.IndexByte(c, str[len(str)-i])
	//	nums += int64(math.Pow(62, float64(i-1)) * float64(pos))
	//}
	return num
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
