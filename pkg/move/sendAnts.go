package move

import (
	"fmt"
	"lem-in/pkg/routes"
	"lem-in/structs"
	"math"
	"strconv"
)

func SendAnts(goodRoutes []string, numberOfAnts int, ourRoutes []structs.Routes, endRoom structs.Room) {

	var solution [][]string
	var paths []int
	for _, val := range goodRoutes {
		for i := range ourRoutes {
			if strconv.Itoa((ourRoutes[i].Key)) == val {
				solution = append(solution, ourRoutes[i].Value)
			}
		}
	}

	for i := range solution {
		if len(solution[i]) <= 2 {
			SendAntsSpecial(solution[i], numberOfAnts)
			return
		}
	}

	var finished, step, currentPath int
	a := make([]structs.Ant, numberOfAnts)

	for i := range solution {
		paths = append(paths, len(solution[i]))
	}

	for i := 0; i < numberOfAnts; i++ {
		id := routes.ShortestRoute(paths)
		paths[id]++
	}
	for i := range paths {
		paths[i] -= len(solution[i])
	}
	var usedRooms []string
	for finished < numberOfAnts {
		usedRooms = nil
		step++
		currentPath = 0
		for i := 0; i < int(math.Min(float64(numberOfAnts), float64((step)*len(solution)))); i++ {
			if !a[i].OnWay {
			kitty:
				if paths[currentPath%len(solution)] != 0 {
					a[i].Path = currentPath % len(solution)
					paths[currentPath%len(solution)]--
				} else {
					currentPath++
					goto kitty
				}
				a[i].OnWay = true
				currentPath++
			}
			a[i].Room++
			if !a[i].Finished {

				if !CheckIfRoomUsedOnSameStep(usedRooms, solution[a[i].Path][a[i].Room], endRoom.Name) {
					fmt.Printf("L%d-%s ", i+1, solution[a[i].Path][a[i].Room])
					usedRooms = append(usedRooms, solution[a[i].Path][a[i].Room])
				} else {
					a[i].Room--
				}
			}
			if a[i].Room == len(solution[a[i].Path])-1 {
				a[i].Finished = true
				finished++
			}
		}
		fmt.Println()
	}
}
