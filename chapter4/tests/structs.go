package main

import (
	"fmt"
	"time"
)

type emptyStruct struct {
}

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int // no managers please; mentors instead!!!
}

type point struct {
	x, y, z int
}

type address struct {
	hostname string
	port     int
}

func structs() {

	var bogdan Employee

	// because bogdan is a variable, its fields are too, so we can assign to a field
	bogdan.Position = "surfer"
	bogdan.Salary -= 5000 // demoted, for writing too few lines of code :)

	// we can take the memory address of a field and access it through a pointer
	position := &bogdan.Position
	fmt.Println(position)                     // 0xc0000441c0
	*position = *position + " extraordinaire" // promoted, for doing battle with the Kraken
	fmt.Println(*position)                    // surfer extraordinaire

	id := &bogdan.ID
	fmt.Println(id)  // 0xc000044180
	fmt.Println(*id) // 0
	*id = 1000
	fmt.Println(*id) // 1000

	// dot notation for accessing struct fields also works with a pointer to a struct
	var employeeOfTheMonth = &bogdan
	fmt.Println(employeeOfTheMonth)    // &{struct}
	fmt.Println(employeeOfTheMonth.ID) // 1000
	employeeOfTheMonth.Position += " // dreams of Portugal"
	(*employeeOfTheMonth).Position += " or Sri Lanka"

	employeeOfTheMonth.Salary = 2500
	(*employeeOfTheMonth).Salary = 3000
	fmt.Println(employeeOfTheMonth.Salary)
	fmt.Println((*employeeOfTheMonth).Salary)

	p := point{1, 2, 3}
	fmt.Println(p)
	p = point{x: 64}
	fmt.Println(p)

	fmt.Println("Scale(factor: 10)", Scale(point{10, 20, 30}, 10))

	fmt.Println("Scale(factor: 64)", Scale(point{1, 2, 3}, 64))

	e := Employee{Salary: 5000}
	bonus := Bonus(&e, 400)
	fmt.Println(bonus)

	AwardAnnualRaise(&e)
	fmt.Println(e.Salary)

	// point pointer!
	pp := &point{56, 44, 82}
	ScaleByReference(pp, 2)
	fmt.Println(*pp)

	hitCount := make(map[address]int)
	hitCount[address{"www.bluenote.com", 80}]++
	hitCount[address{"www.bluenote.com", 80}]++
	fmt.Printf("%v\n", hitCount[address{"www.bluenote.com", 80}])

	addrPtr := &address{"foo.local", 443}
	fmt.Printf("%s %d\n", (*addrPtr).hostname, addrPtr.port)

}

// struct values can be passed as arguments to functions and returned from them
func Scale(p point, factor int) point {
	return point{p.x * factor, p.y * factor, p.z * factor}
}

func ScaleByReference(p *point, factor int) {
	p.x = p.x * factor
	p.y = p.y * factor
	p.z = p.z * factor
}

// for efficiency, larger struct types are usually passed to or returned from functions indirectly, using a pointer
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 110 / 100
}
