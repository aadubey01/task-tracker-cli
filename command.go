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

	flag.StringVar(&cf.Add, "add", "", "Add a new task specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task by index and specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify task by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify task by index to toggle complete true/false")
	flag.BoolVar(&cf.List, "list", false, "List all tasks")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(tasks *Tasks) {
	switch {
	case cf.List:
		tasks.print()
	case cf.Add != "":
		tasks.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for Edit. Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		tasks.edit(index, parts[1])

	case cf.Toggle != -1:
		tasks.toggle(cf.Toggle)
	case cf.Del != -1:
		tasks.delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
