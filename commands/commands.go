package commands

type Command struct {
	Name        string
	Usage       string
	Description string
}

var Commands []Command

func PutCommand(name string, usage string, description string) {
	Commands[0] = Command{Name: name, Usage: usage, Description: description}
}
