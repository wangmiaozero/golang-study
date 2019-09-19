/*以下是无效的标识符：

1ab（以数字开头）
case（Go 语言的关键字）
a+b（运算符是不允许的）*/
/*关键字:
break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var*/
/*36个预定义标识符
append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr*/
package main
import "fmt"
func main() {
	fmt.Println("google"+"Runoob")
	fmt.Println("a"+"b")
}
// 声明变量: Go 语言中变量的声明必须使用空格隔开，如：
var age int ;