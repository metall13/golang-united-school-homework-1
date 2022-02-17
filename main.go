package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	world := emoji.Sprint(":world_map:!")
	fmt.Println("Hello" + world)
}
