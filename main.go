package main

import "fmt"

type testStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (t *testStruct) Shoot() bool {
	if t.On == false {
		fmt.Print("no ")
		return false
	}
	if t.Ammo > 0 {
		t.Ammo -= 1
		fmt.Print("бах! ")
		return true
	} else {
		fmt.Print("осечка ")
		return false
	}
}

func (t *testStruct) RideBike() bool {
	if t.On == false {
		return false
	}
	if t.Power > 0 {
		t.Power -= 1
		return true
	} else {
		return false
	}
}

var st1 testStruct
var st2 = testStruct{true, 5, 3}
var st3 = testStruct{
	On:    true,
	Ammo:  5,
	Power: 1,
}

func main() {

	fmt.Println(" \n go \n")

	fmt.Println(st1)
	fmt.Println(st2)
	fmt.Println(st3)

	fmt.Println("поедем")

	fmt.Println(st1.RideBike())
	fmt.Println(st2.RideBike())
	fmt.Println(st3.RideBike())

	fmt.Println("проехали - посмотрим что получилось")

	fmt.Println(st1)
	fmt.Println(st2)
	fmt.Println(st3)

	fmt.Println("поедем 2")

	fmt.Println(st1.RideBike())
	fmt.Println(st2.RideBike())
	fmt.Println(st3.RideBike())

	fmt.Println("проехали 2 - посмотрим что получилось")

	fmt.Println(st1)
	fmt.Println(st2)
	fmt.Println(st3)

	fmt.Println("шмаляем")

	fmt.Println(st1.Shoot())
	fmt.Println(st2.Shoot())
	fmt.Println(st3.Shoot())

	// ts := new(testStruct)
	// ts.Ammo = 10

	// fmt.Println("------------", ts.Ammo)
}
