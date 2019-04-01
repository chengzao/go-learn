// 条件判断
package main

import "fmt"

func main() {
	gotoF()
}

func gotoF() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签  标签名是大小写敏感的
	println(i)
	if i < 50 {
		i++
		goto Here //跳转到Here去
	}
}

/*
for expression1; expression2; expression3 {
	//...
}
*/

func totalSumF() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	println(sum)
}

func breakF() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			break //← 终止这个循环，只打印 0 到 5
		}
		println(i)
	}

}

func breakF2() {
J:
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break J //终止 j 循环
			}
			println(i)
		}
	}
}

// continue忽略当前循环体内的剩下代码，继续下一轮迭代
func continueF() {
	for i := 0; i < 10; i++ {
		if i < 5 {
			// println(i)  //1234
			continue //忽略本次循环内容,继续下一次循环
		}
		println(i) //56789
	}
}

func foRangeF() {
	var map1 = map[string]int{"key1": 100, "key2": 200}
	for k, v := range map1 {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}
}

//使用_来丢弃不需要的返回值
func foRangeF2() {
	var map1 = map[string]int{"key1": 100, "key2": 200}
	for _, v := range map1 {
		fmt.Println("map's val:", v)
	}
}
