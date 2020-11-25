package study

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Six(){
	//getInfo()
	//getPath()
	//createFile()
	//deleteFile()

	//readFile()
	writeText()
	ioRead()
}

// 获取文件描述信息 os.Stat()
func getInfo(){
	fileInfo,err := os.Stat("./study/abc.txt")
	if err !=nil{
		fmt.Println("stat err:",err)
		return
	}
	fmt.Println("%v\n",fileInfo.Name(),fileInfo.Size(),fileInfo.IsDir(),fileInfo.Mode(),fileInfo.Sys())
}

// 路径、目录操作
func getPath(){
	// 路劲操作
	filename := "./test.txt"
	fmt.Println(filepath.IsAbs(filename)) // 判断是否是绝对路径
	fmt.Println(filepath.Abs(filename)) // 转换为绝对路劲

	//  创建目录
	err := os.Mkdir("./test",os.ModePerm)
	if err !=nil{
		fmt.Println("mkdir err:",err)
		return
	}

	// 创建多级目录
	errAll := os.MkdirAll("./aa/bb",os.ModePerm)
	if errAll !=nil{
		fmt.Println("mkdirall fail",errAll)
		return
	}
}

// 创建文件
func createFile(){
	// Create方式在创建文件时，如果文件不存在则创建，如果文件存在，则会将现有文件修改为空文件
	pathName := "./study/test.txt"
	f,err := os.Create(pathName)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(f)

	file,err := os.OpenFile(pathName,os.O_APPEND|os.O_RDWR,os.ModeAppend)
	if err !=nil{
		fmt.Println("open file err: ", err)
		return
	}
	// 关闭文件
	defer file.Close()

	// 写入内容
	str := "hello golang"
	n,err := file.Write([]byte(str))
	if err !=nil{
		fmt.Println("write fail",err)
		return
	}
	fmt.Println("write number =",n)
}

// 删除文件
func deleteFile(){
	pathName := "./study/test.txt"
	isExist:=Exist(pathName)
	if isExist{
		err := os.Remove(pathName)
		if err !=nil{
			fmt.Println("remove err:",err)
			return
		}
	}else {
		fmt.Println("文件不存在")
	}
}

// 判断文件是否存在
func Exist(filename string) bool{
	_,err:=os.Stat(filename)
	return err == nil || os.IsExist(err)
}


/*---------------------------------------*/

// 文件读取
func readFile(){
	// 打开文件，获取文件指针
	pathName := "./study/abc.txt"
	f,err := os.Open(pathName)
	if err !=nil{
		fmt.Println("open file err:",err)
		return
	}

	// 关闭文件
	defer f.Close()

	readByte := make([]byte, 128)			// 指定要读取的长度
	for {

		fmt.Println("f222==", f)
		n, err := f.Read(readByte)			// 将数据读取如切片，返回值 n 是实际读取到的字节数

		if err != nil && err != io.EOF{		// 如果读到了文件末尾：EOF 即 end of file
			fmt.Println("read file err: ", err)
			break
		}

		fmt.Println("read: ", string(readByte[:n]))
		if n < 128 {
			fmt.Print("read end")
			break
		}

	}

}

// io/ioutil 包文件读取
func ioRead(){
	pathName := "./study/abc.txt"
	bytes,err := ioutil.ReadFile(pathName)
	if err !=nil{
		fmt.Println("read err",err)
		return
	}
	fmt.Println(string(bytes))
}

// 文件写入
func writeText(){
	pathName := "./study/abc.txt"
	// 直接写入
	/*
	f,err := os.OpenFile(pathName,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err !=nil{
		fmt.Println("openfile err",err)
		return
	}
	defer f.Close()

	n,err := f.Write([]byte("hello golang , come on !"))
	if err != nil{
		fmt.Println("write err",err)
		return
	}
	fmt.Println("写入的字节长度",n)
	 */

	// io/ioutil 包文件写入
	str := "加油 golang 好好学习 day day up !"
	err := ioutil.WriteFile(pathName,[]byte(str),os.ModePerm)
	if err != nil{
		fmt.Println("write err:",err)
	}
}