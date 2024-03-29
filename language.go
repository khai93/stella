package stella

type Language struct {
	Id            int
	Name          string
	Version       string
	Image         string   // Docker Image
	Cmd           []string // Command to run the language's program for submission
	TestCmd       []string
	EntryFileName string
	TestFileName  string
	TestFramework TestFramework
}

var Languages = []Language{
	{
		Id:            1,
		Name:          "Go",
		Version:       "latest",
		Image:         "golang",
		Cmd:           []string{"go", "run", "/main.go"},
		TestCmd:       []string{"sh", "-c", "go mod init main;go test -json *.go"},
		EntryFileName: "main.go",
		TestFileName:  "main_test.go",
		TestFramework: GoTestFramework,
	},
	{
		Id:            2,
		Name:          "NodeJS",
		Version:       "latest",
		Image:         "node",
		Cmd:           []string{"node", "main.js"},
		TestCmd:       []string{"jest", "main.test.js", "-c", "{}", "--json"},
		EntryFileName: "main.js",
		TestFileName:  "main.test.js",
		TestFramework: JestTestFramework,
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
