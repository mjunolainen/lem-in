package routes

import (
	"lem-in/structs"
)

var route []string
var routes = make(map[int][]string)
var counter int

func AllPossibleRoutes(currentRoom *structs.Room, previousRoom *structs.Room, endName string) map[int][]string {

	if currentRoom.Name == endName {
		counter++
		route = append(route, currentRoom.Name)
		//routes[counter] = route //sagar = same as below line
		routes[counter] = append(routes[counter], route...)
		route = route[:len(route)-1]
	} else {
		for i := range currentRoom.Connections {
			if currentRoom.Connections[i].Name != previousRoom.Name && currentRoom.Connections[i].AntCount == 0 {
				currentRoom.AntCount = 1
				route = append(route, currentRoom.Name)
				AllPossibleRoutes(currentRoom.Connections[i], currentRoom, endName)
				if len(route) != 0 {
					route = route[:len(route)-1]
				}
			}
		}
	}
	currentRoom.AntCount = 0

	return routes
}
