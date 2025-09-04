package main

func main() {
	tasks := Tasks{}
	storage := NewStorage[Tasks]("tasks.json")
	storage.Load(&tasks)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&tasks)

	// tasks.add("learn golang")
	// tasks.add("Buy bread")
	// tasks.toggle(1)
	// tasks.print()
	storage.Save(tasks)

	// fmt.Printf("%+v\n\n", tasks)
	// tasks.delete(1)
	// fmt.Printf("%+v", tasks)

}
