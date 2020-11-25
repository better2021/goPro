package study

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Sync(){
	// testSunc()
	// objSync()
	// wait()
	result()
}

func testSunc()  {
	var mutex sync.Mutex
	num := 0

	//  开启1000个协程，每个协程都让共享数据
	for i:=0;i<1000;i++{
		go func() {
			mutex.Lock()	// 加错，阻塞其他协程获取锁
			num +=1
			mutex.Unlock()	// 解锁
		}()
	}

	// 大致模拟协程结束 等待3秒
	time.Sleep(time.Second*3)

	// 输出1000，如果没有加锁，则输出的数据很大可能不是1000
	fmt.Println("num=",num)
}

type Account struct {
	money int
	lock *sync.Mutex
}

// 对象加锁
func objSync(){
	a := &Account{
		money: 0,
		lock:  &sync.Mutex{},
	}

	for i:=0;i<10;i++{
		go func(num int) {
			a.Add(num)
		}(10)
	}

	time.Sleep(time.Second*2)
	a.Query()
}

func (a * Account)Query(){
	fmt.Println("当前金额为：",a.money)
}

func (a *Account) Add(num int){
	a.lock.Lock()
	a.money += num
	a.lock.Unlock()
}


/*
等待组 sync.WaitGroup
*/
func wait(){
	var wg sync.WaitGroup 	// 声明一个等待组

	var urls = []string{
		"https://www.baidu.com",
		"http://www.163.com",
		"https://www.weibo.com",
	}

	for _,url := range urls{
		wg.Add(1)		// 每个任务开始，等待组+1
		go func(url string) {
			defer wg.Done()	// (wg *WaitGroup) Done() 等待组计数器-1，等同于Add传入负值
			res,err := http.Get(url)
			fmt.Println(url,err,res.Header)
		}(url)
	}

	wg.Wait()
	fmt.Println("over")
}

func result(){
	var mt sync.Mutex
	var wg sync.WaitGroup
	var money = 1000

	// 开启10个协程，每个协程内部，循环1000次，每次循环+10
	for i :=0 ;i<10;i++{
		wg.Add(1)
		go func(index int) {
			mt.Lock()
			fmt.Printf("协程%d抢到锁\n",index)
			for j:=0;j<100;j++{
				money += 10 //  多个协程对 money产生了竞争
			}
			fmt.Printf("协程%d准备解锁\n",index)
			mt.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("最终的money为：",money)
}