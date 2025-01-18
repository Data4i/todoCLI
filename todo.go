package main

import (
	"fmt"
	"github.com/aquasecurity/table"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CompletedAt *time.Time
	CreatedAt   time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := fmt.Errorf("index out of range")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	todoTable := table.New(os.Stdout)
	todoTable.SetRowLines(false)
	todoTable.SetHeaders("#", "Title", "CreatedAt", "Completed", "CompletedAt")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
		}
		if t.CompletedAt != nil {
			completedAt = t.CompletedAt.Format(time.RFC1123)
		}

		todoTable.AddRow(strconv.Itoa(index), t.Title, t.CreatedAt.Format(time.RFC1123), completed, completedAt)
	}
	todoTable.Render()
}
