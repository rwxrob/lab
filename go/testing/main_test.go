package main_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var _ = MyOwnTest()

func compile() {
	exec.Command("go", "build").Output()
}

func MyOwnTest() error {
	compile()
	fmt.Println(os.Getenv("PWD"))
	out, err := exec.Command("./some").Output()
	log.Print(string(out), err)
	os.Remove("./some")
	os.Exit(0)
	return nil
}

/*
func TestFoo(t *testing.T) {
	t.Fatal()
}
*/
