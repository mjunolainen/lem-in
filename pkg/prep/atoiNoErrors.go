package prep

import (
	"log"
	"strconv"
)

func AtoiNoErrors(s string) int {
	output, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return output
}
