package stella

type Language struct {
	Id            int
	Name          string
	Version       string
	Image         string   // Docker Image
	Cmd           []string // Command to run the language's program for submission
	EntryFileName string
}

var Langauges = []Language{
	{
		Id:            1,
		Name:          "Go",
		Version:       "latest",
		Image:         "golang",
		Cmd:           []string{"go", "run", "/main.go"},
		EntryFileName: "main.go",
	},
	{
		Id:            2,
		Name:          "NodeJS",
		Version:       "latest",
		Image:         "node",
		Cmd:           []string{"node", "main.js"},
		EntryFileName: "main.js",
	},
	{
		Id:            3,
		Name:          "Python",
		Version:       "latest",
		Image:         "python",
		Cmd:           []string{"python", "main.py"},
		EntryFileName: "main.py",
	},
	{
		Id:            4,
		Name:          "C++",
		Version:       "latest",
		Image:         "cpp",
		Cmd:           []string{"sh", "-c", "g++ -o main main.cpp && ./main"},
		EntryFileName: "main.cpp",
	},
}
