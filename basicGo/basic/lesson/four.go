package lesson

import (
	"fmt"
	"io"
	"os"
)


// 拷贝文件
func Four()  {
	fmt.Print(os.Getwd())
	srcFile := "E:/study/goStudy/basicGo/basic/lesson/file/aa.txt"
	descFile:= "E:/study/goStudy/basicGo/basic/lesson/file/bb.txt"
	total,err:= CopyFile(srcFile,descFile)
	fmt.Print(total,err)

	user:="user"
	fmt.Print("/home/" + user)
}

func CopyFile(srcFile,descFile string)(int,error){
	file1,err:=os.Open(srcFile)
	if err!=nil{
		return 0,err
	}
	file2,err:=os.OpenFile(descFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err!=nil{
		return 0,err
	}
	defer file1.Close()
	defer file2.Close()

	// 读写
	bs:=make([]byte,1024,1024)
	n:=-1 // 读取的数据量
	total:= 0
	for{
		n,err= file1.Read(bs)
		if err==io.EOF||n==0{
			fmt.Println("拷贝完毕")
			break
		}else if err!=nil{
			fmt.Println(err)
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}
