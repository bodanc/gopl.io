package main

import (
	"fmt"
	"time"
)

const noDelay time.Duration = 0
const timeout = 5 * time.Minute

// when a sequence of constants is declared as a group, the right-hand side expression may be omitted
// for all but the first of the group, implying that the previous expression and its type should be used again;
const (
	a = 1
	b
	c = 2
	d
	e
	f
	g
)

type Weekday int

const (
	Sunday Weekday = 1 + iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)


type Flags uint8

const (
	FlagUp Flags = 1 << iota // flag is up
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

type powOfTwo int64

const (
	_ powOfTwo = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
)

func testConstants() {

	fmt.Printf("%T %[1]v\n", noDelay)     // time.Duration 0s
	fmt.Printf("%T %[1]v\n", timeout)     // time.Duration 5m0s
	fmt.Printf("%T %[1]v\n", time.Minute) // time.Duration 1m0s

	fmt.Println(a, b, c, d, e, f, g) // 1 1 2 2 2 2 2

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Println(January, February, March, April, May, June, July, August, September, October, November, December)

}

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}
