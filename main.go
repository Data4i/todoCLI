package main

func main() {
	todos := Todos{}

	storage := NewStorage[Todos]("todos.json")
	err := storage.Load(&todos)
	if err != nil {
		panic(err)
	}

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	_ = storage.Save(todos)
}
