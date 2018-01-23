package main

import "fmt"
import evalf "oekaki/evalfunc"

func main() {
	var cycle1 evalf.Cycle
	var cycle2 evalf.Cycle
	cycle1.NewCycle(2.0,3.0,0.0)
	cycle2.NewCycle(1.0,1.0,1.0)
	var reducible evalf.Reducible
	reducible.NewReducible(&cycle1, &cycle2)
	fmt.Println("計算結果：", reducible.Eval(2.0, 3.0))
}
