package main

import (
	"errors"
	"fmt"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}
type TodoList struct {
	Tasks  []*Task
	NextId int
}

func (l *TodoList) Add(title string) *Task {
	l.NextId++
	task := &Task{ID: l.NextId, Title: title, Completed: false}
	l.Tasks = append(l.Tasks, task)
	return task
}
func (l *TodoList) Complete(id int) error {
	if len(l.Tasks) == 0 {
		return errors.New("no tasks found")
	}
	for _, task := range l.Tasks {
		if task.ID == id {
			task.Completed = true
			return nil
		}
	}
	return errors.New("task not found")
}
func (l *TodoList) Remove(id int) error {
	if len(l.Tasks) == 0 {
		return errors.New("no tasks found")
	}
	for index, task := range l.Tasks {
		if task.ID == id {
			l.Tasks = append(l.Tasks[:index], l.Tasks[index+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
func (l *TodoList) Print() {
	if len(l.Tasks) == 0 {
		return
	}
	for _, task := range l.Tasks {
		if task.Completed {
			fmt.Printf("[X] %d. %s\n", task.ID, task.Title)
		} else {
			fmt.Printf("[] %d. %s\n", task.ID, task.Title)
		}
	}
}

func main() {

}
