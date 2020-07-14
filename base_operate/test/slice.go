package test

import (
	"fmt"
	"unicode/utf8"
)

func Test_slice() {
	fmt.Println()
	fmt.Println("-------------testing slice start---------------")

	var arr_1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr_1[0:4] //从索引0截取到索引4，共5个数
	fmt.Printf("s1: %v \n", s1)

	s2 := arr_1[2:] //从索引2截取到尾部
	fmt.Printf("s2: %v \n", s2)

	s3 := arr_1[:3] //从头截取3个元素
	fmt.Printf("s3: %v \n", s3)

	//通过make创建切片
	a := make([]int, 10)
	fmt.Printf("a: %v \n", a)
	fmt.Println("a的len=", len(a))
	fmt.Println("a的cap=", cap(a))

	// 创建一个数组长度是 len=10,cap=15
	b := make([]int, 10, 15)
	fmt.Printf("b: %v \n", b)
	fmt.Println("b的len=", len(b))
	fmt.Println("b的cap=", cap(b))

	for i := 0; i < 10; i++ {
		b = append(b, i+99) //必须写成 b=append(b,x) 否则报错  evaluated but not used
	}
	fmt.Printf("b追加后为: %v \n", b)

	/*************切片的初始化************/
	// 切片的初始化1 直接初始化
	slice_a := []int{1, 2, 3} //注意前面的中括号内没有指定长度和变长度的标识
	// 切片的初始化2 从数组中初始化一个切片 下面三种方法都是初始化一个切片
	slice_b := b[0:10]
	slice_c := b[:10]
	slice_d := b[10:]
	// 切片的初始化3 通过make函数初始化切片
	slice_e := make([]int, 10, 20) //make初始化一个切片 长度是10 cap是20
	fmt.Println(slice_a, slice_b, slice_c, slice_d, slice_e)
	/*************切片的初始化************/

	/*************切片和数组的区别************/
	//数组作为参数和赋值是都是值传递，相当于复制了一份数据赋值给新的变量
	//切片是引用类型

	// 数组是值传递测试
	arr_a := [3]string{"a", "b", "c"}
	arr_b := arr_a
	arr_a[1] = "e"
	fmt.Println("arr_a", arr_a)
	fmt.Println("arr_b", arr_b) //可以看出来修改arr_a的值 不影响数组 arr_b的值，说明是值传递

	// 下面是测试切片的case
	sli_a := []string{1: "b", 5: "e"}
	fmt.Println("sli_a", sli_a)
	sli_b := sli_a

	sli_b[5] = "f" //此处修改了sli_b的值，但是他会影响到sli_a的值，说明他们之间是引用。
	fmt.Println("sli_b", sli_b)
	fmt.Println("sli_a", sli_a) //此处sli_a的值已经被修改。

	/*************切片和数组的区别************/

	/*************切片的追加和copy************/
	sli_e := make([]int, 5)
	sli_e[2] = 2
	fmt.Println("sli_e:", sli_e)
	sli_e = append(sli_e, 6, 7, 8) //可以追加多个
	fmt.Println("sli_e:", sli_e)
	sli_f := make([]int, 3)
	sli_e = append(sli_e, sli_f...) //追加一个slice到另外一个slice上，注意格式需要加上...
	fmt.Println("sli_e:", sli_e)
	/*************切片的追加和copy************/

	//在unicode中，一个中文占两个字节，utf-8中一个中文占三个字节，golang默认的编码是utf-8编码，因此默认一个中文占三个字节，但是golang中的字符串底层实际上是一个byte数组
	str := "hello,世界！"
	slice_1 := []byte(str) //将字符串转化为字符切片
	slice_2 := []rune(str) //将字符串转化为rune切片
	fmt.Printf("slice_1: %v \n", slice_1)
	fmt.Printf("slice_2: %v \n", slice_2)

	fmt.Println("str的长度是:", len(str))
	fmt.Println("str的真实长度是:", utf8.RuneCountInString(str))

	// rune和byte的区别
	// 除开rune和byte底层的类型的区别，在使用上，rune能处理一切的字符，而byte仅仅局限在ascii
	for _, v := range str {
		fmt.Println(v)
	}
	for i := 0; i < len(slice_2); i++ {
		fmt.Println(slice_2[i])
		fmt.Println(string(slice_2[i]))
	}

	fmt.Println("-------------testing slice end---------------")
	fmt.Println()
}
