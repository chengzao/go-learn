package main

import "fmt"

func main() {
	// mapFindF()
	// mapDeleteF()
	varMap()
}

func varMap() {
	//声明1
	var numbers = map[string]int{}
	numbers["ta"] = 11
	fmt.Println(numbers)
	//声明2
	numbers2 := make(map[string]int)
	numbers2["one"] = 1
	numbers2["two"] = 2
	numbers2["three"] = 3
	fmt.Println(numbers2, "==> ", numbers2["three"])
}

func mapFindF() {
	var map1 = map[string]int{"key1": 100, "key2": 200}
	//如果Key存在，将Key对应的Value赋值给v，OK== true. 否则 v 是0，OK==false.
	v, OK := map1["key1"]
	if OK {
		fmt.Println(v, OK)
	} else {
		fmt.Println(v)
	}
	// 这里 不是 :=，是 = ，因为这些变量已经定义过了。
	v, OK = map1["key3"]
	if OK {
		fmt.Println(v, OK)
	} else {
		fmt.Println(v)
	}
}

func mapDeleteF() {
	var map1 = map[string]int{"key1": 100, "key2": 200, "key3": 300}
	for k, v := range map1 {
		fmt.Println(k, v)
		if k == "key2" {
			delete(map1, k)
		}
		if k == "key3" {
			map1["key4"] = 400
		}
	}

	fmt.Println(map1)
}
