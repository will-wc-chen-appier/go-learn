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
