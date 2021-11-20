package move

import "fmt"

func SendAntsSpecial(solution []string, numberOfAnts int) {
	for i := 0; i < numberOfAnts; i++ {
		fmt.Printf("L%d-%s ", i+1, solution[1])
	}
	fmt.Println()
}
