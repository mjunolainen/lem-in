package routes

import (
	"lem-in/structs"
	"strconv"
)

func BestRoutes(a [][]string, antCount2 int, goodRoutes []structs.Routes) int {
	var b []int
	var c []int
	antCount3 := antCount2
	for _, val := range a {
		for i := range val {
			for _, goodRoute := range goodRoutes {
				p, _ := strconv.Atoi(val[i])
				if goodRoute.Key == p {
					b = append(b, len(goodRoute.Value))
					break
				}
			}
		}
		for antCount3 > 0 {
			index := ShortestRoute(b)
			b[index]++
			antCount3--
		}
		antCount3 = antCount2
		w := LongestRoute(b)
		c = append(c, w)
		b = nil

	}
	return ShortestRoute(c)
}
