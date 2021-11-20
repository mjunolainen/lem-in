package routes

import (
	"lem-in/structs"
	"strconv"
)

func AllRouteCombos(cleanedRoutes []structs.Routes) [][]string {

	var routes []string
	var routeComboKeys [][]string
	var keys []string

	for _, val := range cleanedRoutes {
		routes = nil
		keys = nil
		routes = append(routes, val.Value...)
		keys = append(keys, strconv.Itoa(val.Key))
		for _, val2 := range cleanedRoutes {
			if val.Key == val2.Key {
				continue
			}
			if CompareTwoRoutes(routes, val2.Value) {
				routes = append(routes, val2.Value...)
				keys = append(keys, strconv.Itoa(val2.Key))
			}
		}
		routeComboKeys = append(routeComboKeys, keys)
	}
	return routeComboKeys
}
