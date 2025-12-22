package app

import (
	"fmt"
	"time"
)

func TaskManager() *TaskManag {
	return &TaskManag{
		tasks:       []Task{},
		events:      []Event{},
		nextTaskID:  1,
		nextEventID: 1,
	} // конструктор для TaskManag
}

func (tm *TaskManag) AddTask(text string) (int, error) {
	if text == "" {
		return 0, fmt.Errorf("текст задачи не может быть пустым")
	}

	task := Task{
		ID:   tm.nextTaskID,
		Text: text,
		Done: false,
	}
	tm.tasks = append(tm.tasks, task)

	tm.addEvent("add", tm.nextTaskID, text) // событие

	id := tm.nextTaskID
	tm.nextTaskID++

	return id, nil
} // добавление новых задач

func (tm *TaskManag) GetTask(id int) (*Task, error) {
	for _, task := range tm.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("Задача %d. не найдена", id)
} // func GetTask возвращает задачу по ID

func (tm *TaskManag) AllTasks() []Task {
	return tm.tasks
} // func AllTasks возвращает все задачи

func (tm *TaskManag) AllEvents() []Event {
	return tm.events
}

func (tm *TaskManag) DeleteTask(id int) error {
	for i, task := range tm.tasks {
		if task.ID == id {

			tm.addEvent("delete", id, task.Text)

			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("задача %d не найдена", id)
}
func (tm *TaskManag) MarkAsDone(id int) (bool, error) {
	for i, task := range tm.tasks {
		if task.ID == id {
			if task.Done {
				return false, fmt.Errorf("задача %d уже выполнена", id)
			}
			tm.tasks[i].Done = true

			tm.addEvent("done", id, task.Text)

			return true, nil
		}
	}
	return false, fmt.Errorf("задача %d не найдена", id)
} //отмечает задачу как выполненную

func (tm *TaskManag) addEvent(eventType string, taskID int, taskText string) {
	event := Event{
		ID:        tm.nextEventID,
		Type:      eventType,
		TaskID:    taskID,
		TaskText:  taskText,
		Timestamp: time.Now(), // создает и присваивает текущие дату и время полю.
	}
	tm.events = append(tm.events, event)
	tm.nextEventID++
} //добавляет события в историю
