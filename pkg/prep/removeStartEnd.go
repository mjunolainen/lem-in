package prep

import "lem-in/structs"

func RemoveStartEnd(a []structs.Routes) []structs.Routes {
	var output []structs.Routes
	for _, val := range a {
		val.Value = val.Value[1 : len(val.Value)-1]
		output = append(output, val)
	}
	return output
}
