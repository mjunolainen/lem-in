package prep

import (
	"lem-in/structs"
	"log"
	"os"
	"strings"
)

func ReadInput(file string) (*structs.Room, *structs.Room, string) {

	var antNumber int
	var startRoom string
	var endRoom string
	var startFlag bool
	var endFlag bool
	var rooms = make(map[string]*structs.Room)

	openedFile, err := os.ReadFile("testdata/" + file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(openedFile), string(rune(10)))
	antNumber = AtoiNoErrors(lines[0])
	lines = append(lines[1:])

	for _, line := range lines {
		if startFlag {
			roomSlice := strings.Split(line, " ")
			rooms[roomSlice[0]] = &structs.Room{
				Name:     roomSlice[0],
				X:        AtoiNoErrors(roomSlice[1]),
				Y:        AtoiNoErrors(roomSlice[2]),
				AntCount: antNumber,
			}
			startFlag = false
			startRoom = roomSlice[0]
			continue
		}
		if endFlag {
			roomSlice := strings.Split(line, " ")
			rooms[roomSlice[0]] = &structs.Room{
				Name:     roomSlice[0],
				X:        AtoiNoErrors(roomSlice[1]),
				Y:        AtoiNoErrors(roomSlice[2]),
				AntCount: 0,
			}
			endFlag = false
			endRoom = roomSlice[0]
			continue
		}

		switch {
		//startRoom is always the line after ##start
		case line == "##start":
			startFlag = true
		//endRoom is always the line after ##end
		case line == "##end":
			endFlag = true
		//room format is always [name] [X] [Y]
		case strings.Contains(line, " "):
			roomSlice := strings.Split(line, " ")
			rooms[roomSlice[0]] = &structs.Room{
				Name:     roomSlice[0],
				X:        AtoiNoErrors(roomSlice[1]),
				Y:        AtoiNoErrors(roomSlice[2]),
				AntCount: 0,
			}
		//connection format is always [fromRoom]-[toRoom]
		case strings.Contains(line, "-"):
			roomSlice := strings.Split(line, "-")

			fromRoom := rooms[roomSlice[0]]
			toRoom := rooms[roomSlice[1]]

			fromRoom.Connections = append(fromRoom.Connections, toRoom)
			toRoom.Connections = append(toRoom.Connections, fromRoom)

		default:
		}
	}
	return rooms[startRoom], rooms[endRoom], string(openedFile)
}
