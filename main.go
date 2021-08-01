package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isDigit(str string) bool {
	for _, x := range []rune(str) { //[]rune(str)转换成unicode
		if !unicode.IsDigit(x) {
			return false
		}

	}
	return true
}

func isUpper(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsUpper(x) {
			return false
		}
	}
	return true
}

func isLower(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsLower(x) {
			return false
		}
	}
	return true
}

// //func (str string) l_trans_u() {
// func l_trans_u(str string) {
// 	s := ""
// 	for _, x := range []rune(str) {
// 		if unicode.IsLower(x) {
// 			x = unicode.ToUpper(x)
// 		}
// 		s += string(x)
// 	}
// 	str = s
// 	return s
// }

func main() {
	str := strings.Split("acbdw,1269547,AASIDX,AIUydjs,12sjaA,3819247,ausSHSzio,IUFISsi", ",")
	fmt.Println(str)

	//for-range循环 同时返回下标和元素
	for _, s := range str {
		fmt.Println(s, isDigit(s))
		fmt.Println(s, isUpper(s))
		fmt.Println(s, isLower(s))
		//fmt.Println(l_trans_u(s))
	}

}
