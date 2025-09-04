package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Task struct {
	Title       string
	Description string
	Status      string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Tasks []Task

func (tasks *Tasks) add(title string) {
	task := Task{

		Title: title,
		// Description: description,
		// Status:      status,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*tasks = append(*tasks, task)
}

func (tasks *Tasks) validateIndex(index int) error {
	if index < 0 || index >= len(*tasks) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (tasks *Tasks) delete(index int) error {
	t := *tasks
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*tasks = append(t[:index], t[index+1:]...)
	return nil
}

func (tasks *Tasks) toggle(index int) error {
	t := *tasks

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

// func (tasks *Tasks) toggle(index int) error {
// 	t := *tasks

// 	if err := t.validateIndex(index); err != nil {
// 		return err
// 	}

// 	isCompleted := t[index].Completed

// 	if !isCompleted {
// 		// Marking as completed → set completion time
// 		completionTime := time.Now()
// 		t[index].CompletedAt = &completionTime
// 	} else {
// 		// Marking as incomplete → clear completion time
// 		t[index].CompletedAt = nil
// 	}

// 	// Flip the status
// 	t[index].Completed = !isCompleted

// 	return nil
// }

func (tasks *Tasks) edit(index int, title string) error {
	t := *tasks

	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Title = title

	return nil
}

func (tasks *Tasks) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *tasks {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index+1), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
