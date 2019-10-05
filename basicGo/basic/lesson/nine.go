package lesson

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
遍历文件件
*/

func Nine(){
	// 写文件
	fileURL,_ := os.Getwd()
	path := strings.Replace(fileURL, "\\", "/", -1)	// 替换\\为/
	dirname := path + "/lesson/file"
	listFiles(dirname)
}

func listFiles(dirname string)  {
	fileInfos,err := ioutil.ReadDir(dirname)
	if err !=nil{
		log.Fatal(err)
	}

	for _,file := range fileInfos{
		fileName := dirname + "/" + file.Name()
		fmt.Println(fileName,"---")
	}
}