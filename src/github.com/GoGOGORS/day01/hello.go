package main

import "fmt"

//批量声明时，如果某一行声明之后没有赋值则默认同上
const (
	n1 = 100
	n2
	n3
)

//iota go语言的常量计数器，在const关键字出现的时候将被置为0，
//之后const中每新增一行常量声明iota将+1
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

//iota
const (
	b1 = iota //0
	b2 = iota //1
	_  = iota
	b3 = iota //3
)

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

const (
	d1, d2 = iota + 1, iota + 2 //1, 2
	d3, d4 = iota + 1, iota + 2 //2, 4
)

// 定义数量级
const (
	_         = iota
	KB        = 1 << (10 * iota) //1024
	MB        = 1 << (10 * iota) //1048576
	GB        = 1 << (10 * iota) //1073741824
	TB uint64 = 1 << (10 * iota) // 1099511627776
	PB uint64 = 1 << (10 * iota) //1125899906842624
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)

	fmt.Println("-----------------------")

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)

	fmt.Println("-----------------------")

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

	fmt.Println("-----------------------")

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

	fmt.Println("-----------------------")

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println("-----------------------")

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

}
