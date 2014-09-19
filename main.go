package main

import "fmt"

type logger struct {
  msg string
}

type fuse struct {
  *logger
  *second
}

type second struct {}

func (s *second) show(m string) string {
  return m+ "."
}

func main() {
	f := fuse{&logger{msg:"test"},&second{}}
	a:=f.show("pallat") + f.msg
	fmt.Println(a)
}
