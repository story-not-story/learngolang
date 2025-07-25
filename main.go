package main

import (
	"awesomeProject/imooc"
	"awesomeProject/imooc/moody"
	"cmp"
	"fmt"
	"sync"
	"time"
)

/*
	func devide(x int, y int) (int, error) {
		if y == 0 {
			return 0, errors.New("cannot devide 0")
		}
		return x / y, nil
	}
*/

func devide(x int, y int) int {
	if y == 0 {
		panic("cannot devide 0")
	}
	return x / y
}

var m = sync.Mutex{}

func add(p *int) {
	m.Lock()
	defer m.Unlock()
	*p++

}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
type Player interface {
	play()
}

type MyLesson struct {
	name string
	rate float64
}

func (l MyLesson) testtypeof() {
	fmt.Println("test")
}
func (l MyLesson) play() {
	l.rate = l.rate * 200.0
	fmt.Printf("Player %v is playing\n", l.rate)

}

type MyTextLession struct {
	MyLesson
	text string
}
type MyMovie struct {
	name   string
	actors string
}

func (y MyMovie) play() {
	fmt.Println(y.name)
}
func processString(c chan string) {
	/*for {
		s := <-c
		fmt.Printf("begin:%s\n", s)
		time.Sleep(5 * time.Second)
		fmt.Printf("end:%s\n", s)

	}*/
	for s := range c {
		fmt.Printf("begin:%s\n", s)
		time.Sleep(time.Minute)
		fmt.Printf("end:%s\n", s)
	}
}

func inputString1(c chan string) {
	time.Sleep(7 * time.Second)
	c <- "anwer1"
}

func inputString2(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "anwer2"
}
func inputString3(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "anwer3"
}

func devideZero(x int, y int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	i := devide(x, y)
	fmt.Println("devide called", i)
	return i
}

type MyNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func min1[T MyNumber](x, y T) T {
	if x < y {
		return x
	}
	return y
}
func min2[MyNumber2 ~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string](x, y MyNumber2) MyNumber2 {
	if x < y {
		return x
	}
	return y
}
func min3[MyNumber3 cmp.Ordered](x, y MyNumber3) MyNumber3 {
	if x < y {
		return x
	}
	return y
}
func main() {
	i := imooc.MyLesson{"test3", 1, "abc"}
	fmt.Println(i.Testvalue)
	mytest := imooc.Mytest{"aga"}
	fmt.Println(mytest.Name)
	lesson := moody.MyLesson{"test1", 1}
	myLesson := imooc.MyLesson{"test2", 2, "abc"}
	fmt.Println(lesson.LessonName())
	fmt.Println(myLesson.LessonName())
	/*err := os.Mkdir("log", 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
	err2 := os.MkdirAll("log/test1/test2", 0777)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	err3 := os.RemoveAll("log")
	if err3 != nil {
		fmt.Println(err3.Error())
	}

	file1, err10 := os.Open("test.txt")
	defer func(file1 *os.File) {
		err := file1.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(file1)
	if err10 != nil {
		fmt.Println(err10.Error())
	}
	readbytes := make([]byte, 1024)

	for {
		n, err11 := file1.Read(readbytes)
		if err11 != nil {
			fmt.Println(err11.Error())
		}
		if n == 0 {
			break
		}
	}
	_, err12 := os.Stdout.Write(readbytes)
	if err12 != nil {
		fmt.Println(err12.Error())
	}
	err13 := os.Remove("test.txt")
	if err13 != nil {
		fmt.Println(err13.Error())
	}*/

	/*
		file, err4 := os.Create("test.txt")
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println(err.Error())
			}
		}(file)
		if err4 != nil {
			fmt.Println(err4.Error())
		}

		nbytes, err5 := file.WriteString("hello world\r\n")
		if err5 != nil {
			fmt.Println(err5.Error())
		}
		_, err6 := file.WriteAt([]byte("hello world\r\n"), int64(nbytes))
		if err6 != nil {
			fmt.Println(err6.Error())
		}*/

	/*fmt.Println(min1("abc", "xyz"))
	fmt.Println(min2("abc", "xyz"))
	fmt.Println(min3("abc", "xyz"))*/
	//var myLesson = MyLesson{name: "test", rate: 2.0}
	//var myLesson2 = MyLesson{name: "test", rate: 2.0}

	/*Hypothetical illegal Go code:
	copy := &(MyLesson{name: "test", rate: 2.0}.name) // CANNOT take address of temporary copy!
	addr := &copy
	fmt.Println(addr)
	fmt.Println(&copy)*/
	/*valueOf := reflect.ValueOf(&lesson).Elem() // Get addressable value via pointer
	fmt.Println(valueOf)
	//fmt.Println(valueOf.Pointer())
	fmt.Println(valueOf.IsZero())
	valueOf.SetZero() // Now works
	fmt.Println(lesson.name)
	fmt.Println(lesson.rate)
	fmt.Println(valueOf.IsZero())*/
	/*lesson := MyLesson{name: "test", rate: 2.0}
	typeOf := reflect.TypeOf(lesson)
	structField, has := typeOf.FieldByName("name")
	method, exists := typeOf.MethodByName("testtypeof")
	if !exists {
		fmt.Println("method not found")
	} else {
		fmt.Println(method.Name)
		fmt.Println(method.Type)
	}
	if !has {
		fmt.Println("struct field not found")
	} else {
		fmt.Println(structField.Name)
		fmt.Println(structField.Type)
	}
	valueOf := reflect.ValueOf(lesson)
	valueOf.SetZero()
	fmt.Println(valueOf.IsZero())*/
	/*result := devideZero(10, 0)
	fmt.Println("devideZero called", result)*/
	/*
		m := MyLesson{"abc", 1.0}
		(&m).play()
		fmt.Printf("Player %v is playing\n", m.rate)*/

	/*var num int = 1
	for i := 0; i < 1000; i++ {
		go add(&num)
	}

	time.Sleep(10 * time.Second)
	fmt.Println(num)*/
	/*
		c1 := make(chan string)
		fmt.Println("main begin")
		go processString(c1)
		var temp string
		for {
			fmt.Scanln(&temp)
			c1 <- temp
		}*/
	/*
		c1 := make(chan string, 10)
		c2 := make(chan string, 10)
		c3 := make(chan string, 10)
		go inputString1(c1)
		go inputString2(c2)
		go inputString3(c3)
		select {
		case c1 <- "answerNo":
			fmt.Println("c1 winner")
		case <-c2:
			fmt.Println("c2 winner:")
		case s := <-c3:
			fmt.Println("c3 winner:", s)

		}*/
	/*
		var m = sync.Map{}
		m.Store(1, "hello")
		m.Store(2, "world")
		s, t := m.Load(3)
		fmt.Printf("%s,%v\n", s, t)*/
	/*
		var s1 Player = MyLesson{name: "abc1", rate: 2}
		s1.play()
		var s2 Player = MyMovie{
			name:   "abc2",
			actors: "panda",
		}
		s2.play()
		var s3 Player = MyTextLession{
			MyLesson: MyLesson{
				name: "abc3",
				rate: 2,
			},
			text: "fegew",
		}
		s3.play()
		var s4 MyTextLession = MyTextLession{
			MyLesson: MyLesson{
				name: "abc3",
				rate: 2,
			},
			text: "fegew",
		}
		fmt.Println(s4.MyLesson.name, s4.rate, s4.text)
	*/
	/*
		s1 := map[int]string{
			537: "hello",
			543: "world",
			234: "!!!",
		}
		s1[234] = "!"
		for k, v := range s1 {
			fmt.Printf("%d: %s\n", k, v)
		}

		s2 := make(map[int]string, 3)
		s2[537] = "hello"
		s2[543] = "world"
		s2[234] = "!!!"
		delete(s2, 543)
		m, n := s2[543]
		fmt.Println(m)
		fmt.Println(n)
		for i, j := range s2 {
			fmt.Printf("%d: %s\n", i, j)
		}
		fmt.Println(len(s2))*/

	/*	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
		// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
		var s1, s = "hello world", "test"

		var s2 = "abc"
		fmt.Println(s2)
		fmt.Printf("Hello and welcome, %s!\n", s)
		b := true
		c := 9
		fmt.Printf(
			"%s, %s\n",
			s1,
			s2,
		)
		fmt.Println(true, c)
		for i := 1; i <= 5; i++ {
			//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
			// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
			fmt.Println("i =", 100/i)
		}

		const mm, nn = true, -3
		const (
			tt = 5
			gg = 6
		)
		const (
			ll        = iota
			ii string = "test"
			jj int64  = 5
			ss        = iota
			pp        = iota
		)
		fmt.Println(ll)
		fmt.Println(ss)
		fmt.Println(pp)
		fmt.Println(a1)
		fmt.Println(a2)
		fmt.Println(a3)
		fmt.Println(t4)
		const (
			a1 = 1 << iota
			a2
			a3
			t4 = 7
			a4 = 1 << iota
			a5
			a6
		)
		fmt.Println(a4)
		fmt.Println(a5)
		fmt.Println(a6)
		var ff = math()
		fmt.Println(ff(4, 5, 6, 7, 8, 9))
	*/
	/*
		var myarremp = [...]int{11, 22, 33, 44, 55, 66}
		myarr := myarremp[:4]
		s5 := make([]int, len(myarr))
		for k, v := range s5 {
			fmt.Printf("%d:%d\n", k, v)

		}
		i := 0
		for {
			if i >= len(myarr) {
				break
			}
			fmt.Printf("%d\n", myarr[i])
			i++
		}

		i = 0
		for {
			if i >= len(myarremp) {
				break
			}
			fmt.Printf("%d\n", myarremp[i])
			i++
		}
	*/
	/*var g [5]int
	g[0] = 1
	g[1] = 2
	g[2] = 3
	g[3] = 4
	g[4] = 5
	fmt.Println(g)
	n := [4]int{1, 2, 3, 4}
	fmt.Println(n)
	m := [...]int{1, 2, 3}
	fmt.Println(m)
	t := [5]int{0: 1, 2: 3, 4: 5}
	fmt.Println(t)*/
	/*
		myarr2 := [3][4][5]int{
			{
				{1, 2, 3, 4, 5},
				{11, 22, 33, 44, 55},
				{111, 122, 133, 144, 155},
				{211, 222, 233, 244, 255},
			},
			{
				{1, 2, 3, 4, 5},
				{11, 22, 33, 44, 55},
				{111, 122, 133, 144, 155},
				{211, 222, 233, 244, 255},
			},
			{
				{1, 2, 3, 4, 5},
				{11, 22, 33, 44, 55},
				{111, 122, 133, 144, 155},
				{211, 222, 233, 244, 255},
			},
		}
		for i := 0; i < 3; i++ {
			//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
			// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>

			for j := 0; j < 4; j++ {
				//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
				// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>

				for k := 0; k < 5; k++ {
					//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
					// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
					fmt.Printf("%d,", myarr2[i][j][k])
				}
				fmt.Println()
			}
			fmt.Println()

		}*/

}

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func arr1(ii ...int) int {
	sumvalue := 0

	for i := 0; i < len(ii); i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		sumvalue = sumvalue + ii[i]
	}
	return sumvalue
}

func math() func(ii ...int) int {

	return func(ii ...int) int {
		sumvalue := 0

		for i := 0; i < len(ii); i++ {
			//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
			// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
			sumvalue = sumvalue + ii[i]
		}
		return sumvalue
	}

}
