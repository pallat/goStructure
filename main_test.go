package main

import (
  "testing"
)

func testShow(t *testing.T) {
        f := fuse{&logger{msg:"test"},&second{}}
        a:=f.show("pallat") 

	if a != "pallat." {
		t.Errorf("%s",a)
	}
}
