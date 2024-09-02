package main

import (
	"crypto/sha1"
)

var DigVecUpper = []byte{
	'0', '1', '2', '3',
	'4', '5', '6', '7',
	'8', '9', 'A', 'B',
	'C', 'D', 'E', 'F',
	'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R',
	'S', 'T', 'U', 'V',
	'W', 'X', 'Y', 'Z',
}

func Octet2Hex(src []byte) string {
	var result string
	for _, s := range src {
		result += string(DigVecUpper[s>>4])
		result += string(DigVecUpper[s&0x0F])
	}
	return result
}

func main() {
	password := "123456"
	result := sha1.Sum([]byte(password))
	result = sha1.Sum(result[:])
	scrambled := "*" + Octet2Hex(result[:])
	println(scrambled)
}
