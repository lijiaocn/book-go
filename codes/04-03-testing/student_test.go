//create: 2018/01/08 14:31:41 change: 2018/01/08 14:48:16 lijiaocn@foxmail.com
package main

import (
	"os"
	"testing"
)

var boy Student
var girl Student

func TestIsBoy(t *testing.T) {
	if !boy.IsBoy() {
		t.Errorf("i'm a boy ,but method IsBoy return fasle\n")
	}
	if girl.IsBoy() {
		t.Errorf("i'm a girl ,but method IsBoy return true\n")
	}
}

func TestMain(m *testing.M) {
	boy = Student{
		Gender: "Male",
	}
	girl = Student{
		Gender: "Female",
	}
	os.Exit(m.Run())
}

//If use vim-go, just run:
//GoTest
//GoTestFunc

//If in shell, just run:
//go test
