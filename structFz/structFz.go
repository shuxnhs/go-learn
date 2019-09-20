package main

import "fmt"

// 类的简单封装
type People struct {
	name string
	age int
	hobby []string
}

func (p *People)GetName()  {
	fmt.Println("名字是", p.name, "开始吃饭")
}

func (p *People)GetAge()  {
	fmt.Println("今年", p.age, "岁")
}

func (p *People)GetHobby()  {
	for k, v := range p.hobby {
		fmt.Println("key =", k, "value = ", v)
	}
}

// 类的继承
type Men struct {
	// 继承类people
	People
	hair bool
}

func (man *Men)GetHair()  {
	if man.hair != true{
		fmt.Println("男人的头发是短的")
	}else {
		fmt.Println("男人的头发是长的")
	}
}

// 重写父类的方法
func (man *Men)GetAge()  {
	fmt.Println("男人的年龄是：", man.People.age)
}

// 多态，接口
type Player interface {
	Input(music string)(happy int)
	Output()
}

// mp3Player来继承抽象接口
type Mp3 struct {
	name string
}

func (mp3 *Mp3)Input(music string)(happy int)  {
	happy = 100
	fmt.Println(mp3.name, "播放器🎵，开始播放", music, "开心值为", happy)
	return happy
}

func (mp3 *Mp3)Output()  {
	fmt.Println("mp3输出")
}

// mp4Player来继承抽象接口
type Mp4 struct {
	name string
}

func (mp4 *Mp4)Input(music string)(happy int)  {
	happy = 1000
	fmt.Println(mp4.name, "播放器🎵，开始播放", music, "开心值为", happy)
	return happy
}

func (mp4 *Mp4)Output()  {
	fmt.Println("mp4输出")
}


func main()  {
	people1 := People{name:"hxh", age:11}
	people1.GetAge()

	people2 := People{}
	people2.name = "wxz"
	people2.GetName()
	// 默认为0
	people2.GetAge()

	people3 := People{"hxh", 11, []string{"吃","喝", "拉", "撒"}}
	people3.GetHobby()

	// 新的创建方法
	man1 := new(Men)
	man1.People = People{name:"wxz", age: 11}
	man1.hair = true
	man1.GetHair()
	man1.GetAge()

	// 多态
	player := make([]Player, 0)
	player = append(player, &Mp3{"MP3"})
	player = append(player, &Mp4{"mp4"})
	for _, v := range player{
		// 类型断言
		switch v.(type) {
		// 开始判断类型
		case *Mp4:
			fmt.Println("我是mp4")
			v.Input("鸡你太美")
			v.Output()
		default:
			fmt.Println("我是mp3")
			v.Input("鸡你太美")
			v.Output()
		}

	}

}
