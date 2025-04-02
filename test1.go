package main

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"rsc.io/quote"
)

func add(x, y int) int {
	return x + y
}

func main() {
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(arr[2])
	fmt.Println(slice[3])

	go func() {
		fmt.Println("go routine")
	}()

	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("오늘")
	case today + 1:
		fmt.Println("내일")
	case today + 2:
		fmt.Println("이틀 후")
	default:
		fmt.Println("너무 나중이네")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MAC OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s \n", os)
	}

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	var p *int

	i := 100
	p = &i

	fmt.Println(*p)
	*p = 101
	fmt.Println(*p)

	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		//p  = &Vertex{3, 4}
	)

	fmt.Println(v1.X)
	fmt.Println(v1.Y)
	fmt.Println("-----")
	fmt.Println(v2.X)
	fmt.Println(v2.Y)
	fmt.Println("-----")
	fmt.Println(v3.X)
	fmt.Println(v3.Y)
	fmt.Println("-----")
	//fmt.Println(*p.X)

	names := [4]string{
		"john",
		"paul",
		"george",
		"singo",
	}

	a := names[0:2]
	b := names[1:3]
	fmt.Println(names)
	fmt.Println(a)
	fmt.Println(b)

	b[0] = "jonnung"
	fmt.Println(names)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(quote.Go())
}
