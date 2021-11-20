package structs

type Room struct {
	Name        string
	X, Y        int
	Connections []*Room
	AntCount    int
}
