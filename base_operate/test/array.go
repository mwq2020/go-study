package test

import "fmt"

func Test_array() {
	// 数组类型的基本定义
	fmt.Println()
	fmt.Println("-------------testing array start---------------")
	var arr_1 [10]int              //定义一个类型是int的数组，它有10个元素，每个元素的默认值是0
	var arr_2 [5]string            //定义一个类型是字符串类型的数组，长度是5，每个元素的默认值都是空
	var arr_3 = [3]int{11, 13, 15} //定义一个类型是int的数组，长度是3，并初始化数组的值为11，13，15
	arr_4 := [3]int{17, 19, 21}    //定义一个类型是int的数组，长度是3，并初始化数组的值为11，13，15

	fmt.Printf("arr_1:  %v \n", arr_1)
	fmt.Printf("arr_2:  %v \n", arr_2)
	fmt.Printf("arr_3:  %v \n", arr_3)
	fmt.Printf("arr_4:  %v \n", arr_4)

	arr_5 := [...]int{1, 2, 3, 4} //数组的长度由后面的初始化长度来定义
	fmt.Printf("arr_5:  %v \n", arr_5)

	arr_6 := [3]int{1: 1, 2: 3} //指定总长度，并通过索引值初始化，没初始化话的元素使用默认类型
	fmt.Printf("arr_6:  %v \n", arr_6)

	arr_7 := [...]int{1: 1, 2: 3, 7: 9} //不指定总长度，通过索引值进行初始化，数组的长度由最后一个索引值确定，没有初始化的元素使用默认类型
	fmt.Printf("arr_7:  %v \n", arr_7)

	// 数组的特点
	// 数组创建完长度就固定了，不可以在追加元素
	// 数组是值类型的，数组赋值或者作为函数参数都是值拷贝
	// 数组的长度是数组的组成部分，[10]int和[20]int表示不同的类型
	// 可以根据数组创建切片

	// 数组的遍历 range遍历
	for k, v := range arr_7 {
		fmt.Println("arr_7:[k,v]", k, v)
	}

	//通过索引遍历数组及访问数组的值
	for i := 0; i < len(arr_6); i++ {
		fmt.Println("arr_6[i]", i, arr_6[i])
	}

	fmt.Println("-------------testing array end---------------")
	fmt.Println()
}
