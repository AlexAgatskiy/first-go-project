package app

import "time"

type Task struct {
	// В чем смысл делать переменные экспортируемыми? (с Большой буквы). Как
	// будто не нужно давать контроль над внутренними полями внешним модулям.
	ID   int
	Text string
	Done bool
} //структура показывает задачу в списке

type Event struct {
	// Вопрос аналогичный верхнему.
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
