package main

import (
	"github.com/charmbracelet/bubbles/list"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

// Custom item
type Task struct {
	status      status
	title       string
	description string
}

// Implement the list item interface
func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

// Main model
type Model struct {
	list list.Model
	err  error
}

// TODO: call this on tea.WindowsSizeMsg
func (m *Model) initList(width, height) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "To Do"
	m.list.SetItems([]list.Item{
		Task{status: todo, title: "buy milk", description: "Strawberry milk"},
		Task{status: todo, title: "eat sushi", description: "miso soup too"},
		Task{status: todo, title: "laundry", description: "or not idc"},
	})
}
