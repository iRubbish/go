//package main
//
//func main() {
//	//var a int
//	var b int32
//	var c  int32
//	//a = 15
//	c = 12
//	b = b + c	 // 编译错误
//	c = b + 5    // 因为 5 是常量，所以可以通过编译
//	print(b)
//	print(c)
//}
package main
import "fmt"
func main() {
	var x uint8 = 15
	var y uint8 = 4
	fmt.Println("%b",x,y)
	fmt.Printf("%08b\n", x &^ y)  // 00001011
}
