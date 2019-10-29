package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type LogProcess struct {
	rc chan []byte
	wc chan string
	read Reader
	write Writer
}

type ReadFromFile struct {
	path string // 读取文件的路劲
}

type WriteToInfluxDB struct {
	influxDBDsn string // influx data source
}

func (r*ReadFromFile) Read(rc chan []byte) {
	// 读取模块
	f,err := os.Open(r.path)
	if err != nil{
		fmt.Println("open file err",err.Error())
	}

	// 从文件末尾开始逐行读取文件内容
	f.Seek(0,2)
	rd := bufio.NewReader(f)

	for  {
		line,err := rd.ReadBytes('\n')
		if err == io.EOF{
			time.Sleep(500*time.Millisecond)
		}else if err !=nil{
			fmt.Sprintf("ReadyBytes error",err.Error())
		}
		rc <- line[:len(line)-1]
	}

}

func (w *WriteToInfluxDB) Write(wc chan string){
	// 写入模块
	fmt.Println(<-wc)
}

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

func main()  {
	r := &ReadFromFile{
		path:        "./tmp/access.log",
	}

	w := &WriteToInfluxDB{
		influxDBDsn: "username&password..",
	}

	lp := &LogProcess{
		rc:make(chan []byte),
		wc:make(chan string),
		read:r,
		write:w,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(10*time.Second)
}

func (l *LogProcess) Process(){
	// 解析模块
	data := <- l.rc
	l.wc <- strings.ToUpper(string(data))
}
