package main

import (
	"fmt"
	"time"

	"github.com/bloodydevil/string"
)

func main() {
	fmt.Println(string.Reverse("This should print Hello world backwards!"))

	fmt.Println("The time is", time.Now())
}
