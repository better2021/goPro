package lesson

import "fmt"

/*
interface 接口
*/

type person01 struct {
	name string
}

// 定义接口
type Tranice interface {
	sing()
	dance()
	rap()
	playBasketball()
}

type Player interface {
	playBasketball()
}

type Flyer interface {
	fly()
}

// 方法
func show(t Tranice){
	t.sing()
	t.dance()
	t.rap()
	t.playBasketball()
}

func play(player Player){
	player.playBasketball()
}

func fly(flyer Flyer)  {
	flyer.fly()
}

func (p person01) sing(){
	fmt.Println(p.name,"sing")
}

func (p person01) dance(){
	fmt.Println(p.name,"dance")
}

func (p person01) rap(){
	fmt.Println(p.name,"rap")
}

func (p person01) playBasketball(){
	fmt.Println(p.name,"playBasketball")
}

func (p person01) play() {
	fmt.Println("player")
}

func (p person01) fly() {
	fmt.Println("fly")
}


func Face(){
	//joy := person01{name:"fangzi"}
	//show(joy)
	//play(joy)
	//fly(joy)

	play := eneny()
	play()
	play()
	play()

	adder(2,3)

	fmt.Printf("%s\n","----")

	sum := curryAdd(1)
	sum(2)
	sum(3)
}


/*
闭包
*/
func eneny() func() {
	key := 0
	return func() {
		key++
		fmt.Printf("%d\n",key)
	}
}

func adder(a,b int){
	fmt.Println(a+b)
}

func curryAdd(a int) func(int){
	return func(b int) {
		fmt.Printf("%v\n",a + b)
	}
}