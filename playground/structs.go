package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPersonPtr(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func newPersonCopy(name string) person {
	p := person{name: name}
	p.age = 30
	return p
}

func structTest() {
	fmt.Println(newPersonPtr("will"))
	fmt.Println(newPersonCopy("kang"))

	// fmt.Println(&person{name: "Ann", age: 40})
	// fmt.Println(person{name: "Fred"})
	// fmt.Println(person{age: 100})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sptr := &s
	fmt.Println(sptr, sptr.name)

	mumu := struct {
		name   string
		isGood bool
	}{
		name:   "mumu",
		isGood: true,
	}
	fmt.Println(mumu)
}

type rect struct {
	width, height, _area, _perim int
}

func (r *rect) area() int {
	r._area = r.width * r.height
	return r._area
}

func returnArea(r *rect) int {
	return r.area()
}

func (r rect) perim() int {
	r._perim = 2*r.width + 2*r.height
	return r._perim
}

func methodTest() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	fmt.Println(r)
	// test := func(r *rect) int {
	// 	return r.area()
	// }
	// fmt.Println("area: ", test(&r))
	// fmt.Println("area: ", returnArea(&r))

	// rp := &r
	// fmt.Println("area: ", rp.area())
	// fmt.Println("perim:", rp.perim())
}

type base struct {
	num int
}

func (b *base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base //embedded
	str  string
}

type container2 struct {
	base base // not embedded
	str  string
}

type base2 struct {
	num int
}

func (b base2) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container3 struct {
	base2
	str string
}

func embedTest() {
	co := container{
		base: base{num: 42},
		str:  "hello",
	}
	fmt.Println(co.describe())
	fmt.Println("also num:", co.base.num)
	fmt.Println(co.base.describe())
	fmt.Println((&co).describe())
	fmt.Println((&co.base).describe())
	fmt.Println((&(co.base)).num)

	co3 := container3{
		base2: base2{num: 100},
		str:   "hello2",
	}
	fmt.Println(co3.describe())
	fmt.Println("also num:", co3.base2.num)
	fmt.Println(co3.base2.describe())

	type describer interface {
		describe() string
	}

	// var d describer = co //can't compile
	// fmt.Println(d.describe())
	var d describer = &co
	fmt.Println(d.describe())

	var d3 describer = co3
	fmt.Println(d3.describe())
}
