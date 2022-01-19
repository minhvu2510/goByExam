package main

import "fmt"

type SalaryCaculator interface {
	caculate() int
	// devs() int
}

type Permanent struct {
	name     string
	basicpay int
	pf       int
}
type Contract struct {
	name     string
	basicpay int
}

func (per *Permanent) caculate() int {
	return per.basicpay + per.pf
}
func (Con *Contract) caculate() int {
	return Con.basicpay
}

// tính tổng phải trả củac công ty

func totalExpense(p []SalaryCaculator) {
	total := 0
	for _, cal := range p {
		total += cal.caculate()
	}
	fmt.Println("===>", total)
}
func main() {
	pemp1 := Permanent{"minhvu", 34, 20}
	pemp2 := Permanent{"vudeptrai", 6000, 30}
	cemp1 := Contract{"dm", 3000}
	employees := []SalaryCaculator{&pemp1, &pemp2, &cemp1} // khai bao dc vì cac doi tuong da có đầy đủ method trong interface
	totalExpense(employees)

}
