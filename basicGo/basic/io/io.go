package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件操作
func main()  {
	// readFile()
	// readByIoutil()
	// write()
	// writeByte()
	//writeByiutil()

	trun()
}

// 读文件,只能读取很短的文件
func readFile(){
	fileObj,err:=os.Open("./test.txt")		// 相对路劲
	if err!=nil{
		fmt.Println("upen file failed",err)
	}

	defer fileObj.Close() // 关闭文件

	// 读取文件的内容
	var tmp = make([]byte,128)
	n,err:=fileObj.Read(tmp)
	if err != nil{
		fmt.Println("read from file failed",err)
	}
	fmt.Printf("read %d bytes from file.\n",n)
	fmt.Println(string(tmp))
}

// ioutil 可以一次性读取文件，但文件不能太大
func readByIoutil(){
	content,err:=ioutil.ReadFile("./test.txt")
	if err != nil{
		fmt.Println("read file by iouttil failed,err:%v\n",err)
		return
	}
	fmt.Println(string(content))
}

/*
写入文件
os.O_APPEND 会追加内容到文件后面
os.O_TRUNC 会覆盖前面的内容
 */
func write(){
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err!=nil{
		fmt.Println("open file failed",err)
		return
	}
	defer file.Close()
	str := "敖德萨多阿萨德haha"
	file.Write([]byte(str))
	file.WriteString("hello 呀呀呀123456")
}

// 文件会先输入缓冲区，可以减少与磁盘的交互
func writeByte(){
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err!=nil{
		fmt.Println("open file failed",err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("小发明家123")
	write.Flush() // 将缓冲区的内容写入磁盘
}

/*
 将字节类型直接写入
 不用打开和关闭文件
 */

func writeByiutil(){
	str:= "我命由我不由天"
	err:=ioutil.WriteFile("./test.txt",[]byte(str),0644)
	if err!=nil{
		fmt.Printf("write file failed%v\n",err)
		return
	}
}

/*
	裁剪一个文件到100个字节。
 	如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
	如果文件本来超过100个字节，则超过的字节会被抛弃。
	这样我们总是得到精确的100个字节的文件。
	传入0则会清空文件。
*/
func trun(){
	//fmt.Println(123)
	//err := os.Truncate("./test.txt",10)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(err)

	fileInfo,err := os.Stat("test.txt")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(fileInfo)
	// 文件名称
	fmt.Println("file name:",fileInfo.Name())
	fmt.Println("size in bytes:",fileInfo.Size())
	fmt.Println("permissions:",fileInfo.Mode())
	fmt.Println("Last modified:",fileInfo.ModTime())
	fmt.Println("is directory:",fileInfo.IsDir())
	fmt.Println(fileInfo.Sys())

}