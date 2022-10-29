package stella

type Language struct {
	Id            int
	Name          string
	Version       string
	Image         string // Docker Image
	Cmd           string // Command to run the language's program for submission
	EntryFileName string
}

var Langauges = []Language{
	{
		Id:            1,
		Name:          "Go",
		Version:       "latest",
		Image:         "golang",
		Cmd:           "go run /main.go",
		EntryFileName: "main.go",
	},
	{
		Id:            2,
		Name:          "NodeJS",
		Version:       "latest",
		Image:         "node",
		Cmd:           "node main.js",
		EntryFileName: "main.js",
	},
}
