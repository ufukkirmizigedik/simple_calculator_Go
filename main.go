package main

import "fmt"

type kadin string
type erkek string

func (e erkek) yuru() string {
	return "ben erkegim ve yuruyorum"
}

func (k kadin) yuru() string {
	return "ben bir kadinim ve yuruyorum"
}

type Ufuk interface {
	yuru()
}

func (u Ufuk) mars() string {
	u.yuru()
}

func main() {
	fmt.Println(mars)
}
