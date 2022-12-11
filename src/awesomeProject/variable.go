//package main
//
//var a = "G"
//
//func main() {
//	n()
//	m()
//	n()
//}
//
//func n() { print(a) }
//
//func m() {
//	a := "O"
//	print(a)
//}


//
//package main
// 局部变量
//var a = "G"

//func main() {
//	n()
//	m()
//	n()
//}
//
//func n() {
//	print(a)
//}
//
//func m() {
// 全局变量
//	a = "O"
//	print(a)
//}

//package main
//
//var a string
//
//func main() {
//	a = "G"
//	print(a)
//	f1()
//}
//
//func f1() {
//	a := "O"
//	print(a)
//	f2()
//}
//
//func f2() {
//	print(a)
//}
package main

var a int
func main()  {
	a = 0
	print(a)
	main1()
	main2()
}
func main1()  {
	a := 1
	print(a)
}
func main2()  {
	print(a)
}

// = 与 := 区别在于 = 是变量赋值, := 是定义一个变量并赋值, var的省略写法
// 一般情况 var 是用作于包级别的全局变量写法, := 用作与函数内的局部变量. 不过都可以这么写 只是大家的普遍情况.