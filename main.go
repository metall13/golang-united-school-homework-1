package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	world := emoji.Sprint("Hello :world_map:!")
	fmt.Println(world)
}
