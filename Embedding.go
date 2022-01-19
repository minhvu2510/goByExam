package main

import "fmt"

type base struct {
	num int
}

func (bas *base) descibe() string {
	return fmt.Sprintf("base with num=%v", bas.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	type desciption interface {
		descibe() string // base đã implement cái này -> co = dc interface
	}
	var dm desciption = &co
	fmt.Println(dm.descibe())
}
