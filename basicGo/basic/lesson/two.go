package lesson

import "fmt"

/**
 go语言的指针（值传递）
当一个指针被定义后没有分配到任何变量时，它的值为nil
&（变量）取地址，*（变量）取值
nil来指针也称为空指针
nil表示零值或者空值
 */
func Two()  {
	var a int = 20 // 声明实际变量
	var ip *int  // 声明指针变量
	ip = &a // 指针变量存储地址
	fmt.Printf("a变量:%v\n",a)
	fmt.Printf("a变量的地址是：%x\n",&a)
	fmt.Printf("ip变量存储的指针地址:%x\n",ip)
	/*使用指针访问*/
	fmt.Printf("*ip变量的值 :%d\n",*ip)

	p := new(int)
	*p = 666
	fmt.Println(*p,&p,p)

	x,y:=10,20
	// 通过函数交换x和y
	swap(&x,&y)
	fmt.Println(x,y,"--")

	var arr []int
	for i:=0;i<25;i++{
		fmt.Println(i)
		arr = append(arr,i)
	}
	fmt.Println(arr)
	newArr := arr[5:10]
	fmt.Println(newArr,len(newArr))
	fmt.Println(len(arr))
	fmt.Println(arr[len(arr)-1])

	// 使用空接口类型定义任意类型的数组
	//var anyArr []interface{}
	anyArr := make([]interface{},3)
	anyArr[0] = 15
	anyArr[1] = "20"
	anyArr[2] = 23.5
	anyArr = append(anyArr,"qq")
	fmt.Println(anyArr,len(anyArr))
}

func swap(a,b *int){
	*a,*b=*b,*a
	fmt.Println(*a,*b,"--++")
}