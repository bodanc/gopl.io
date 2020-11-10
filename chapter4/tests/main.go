package main

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {

	arrays()
	slices()
	maps()

}
