package main

import (
	"fmt"
	"lem-in/pkg/move"
	"lem-in/pkg/prep"
	"lem-in/pkg/routes"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments")
		return
	}
	startRoom, endRoom, fileInput := prep.ReadInput(os.Args[1])
	antCount := startRoom.AntCount

	if antCount < 1 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		return
	}

	//find all possible routes
	allRoutes := routes.AllPossibleRoutes(startRoom, startRoom, endRoom.Name)

	//sort all possible routes by length
	sortedRoutes := routes.SortRoutes(allRoutes)

	//remove start and end rooms to compare routes
	cleanedRoutes := prep.RemoveStartEnd(sortedRoutes)

	//find all non-overlapping route combos
	allRouteCombos := routes.AllRouteCombos(cleanedRoutes)

	if len(allRouteCombos) < 1 {
		fmt.Println("ERROR: invalid data format, no path from ##start to ##end")
		return
	}
	fmt.Println(fileInput)

	//find the best routes for the given number of ants
	shortestRouteIndex := routes.BestRoutes(allRouteCombos, antCount, cleanedRoutes)

	//send ants from start to end
	move.SendAnts(allRouteCombos[shortestRouteIndex], antCount, sortedRoutes, *endRoom)

	//TODO celebrate the finished project :)
}
