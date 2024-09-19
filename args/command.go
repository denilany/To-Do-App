package args

import (
	"ToDo/todo"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit todo by index and specify new title")
	flag.IntVar(&cf.Del, "del", -1, "Specify the index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *todo.Todos) {
	switch {
	case cf.List:
		todos.Print()

	case cf.Add != "":
		todos.Add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.Edit(index, parts[1])

	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)

	case cf.Del != -1:
		todos.Delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
