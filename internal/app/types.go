package app

import "time"

type Task struct {
	ID   int
	Text string
	Done bool
} //структура показывает задачу в списке

type Event struct {
	ID        int
	Type      string
	TaskID    int
	TaskText  string
	Timestamp time.Time
} // собирает события в истории

type TaskManag struct {
	tasks       []Task
	events      []Event
	nextTaskID  int
	nextEventID int
} // структура задач и событий
