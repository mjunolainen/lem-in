package routes

import (
	"lem-in/structs"
	"sort"
)

//SortRoutes arranges the routes based on the length or how many steps an ant has to make
//to get from the start to the end
func SortRoutes(allRoutes map[int][]string) []structs.Routes {

	var sortedRoutes []structs.Routes
	for k, v := range allRoutes {
		sortedRoutes = append(sortedRoutes, structs.Routes{Key: k, Value: v})
	}

	sort.Slice(sortedRoutes, func(i, j int) bool {
		return len(sortedRoutes[i].Value) < len(sortedRoutes[j].Value)
	})
	return sortedRoutes
}
