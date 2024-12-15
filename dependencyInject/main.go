package main

import "fmt"

// rock climbing game
type RockClimber struct {
	rocksClimbed int
	sp           SafetyPlacer
}

type SafetyPlacer struct {
	kind int
}

func (c *RockClimber) climbRock() {
	c.rocksClimbed++
	if c.rocksClimbed == 10 {
		c.sp.placeSafeties()
	}
}

// heavy task has to check all cases
func (c *SafetyPlacer) placeSafeties() {
	fmt.Println("Placing safeties")
}

func main() {
	climber := RockClimber{}
	for i := 0; i <= 9; i++ {
		climber.climbRock()
		fmt.Println(i)
	}
}
