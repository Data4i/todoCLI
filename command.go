package main

import (
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
	flag.StringVar(&cf.Add, "add", "", "Add a new TODO title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a TODO title by index. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify TODO to delete by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify TODO to be toggle Completed by index")
	flag.BoolVar(&cf.List, "list", false, "List all TODOS")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, Invalid format for edit. Use id:new_title")
			os.Exit(1)
		}
		idx, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error, Invalid index for edit")
			os.Exit(1)
		}
		err = todos.edit(idx, parts[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case cf.Toggle != -1:
		err := todos.toggle(cf.Toggle)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case cf.Del != -1:
		err := todos.delete(cf.Del)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Println("Invalid command")
	}
}
