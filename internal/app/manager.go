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
	// комментарии-документация пишется перед словом объявлением сущности, в
	// данном случае перед func ... (см. пример ниже). Исправить по всему коду.
}

// AddTask - добавляет новые задачи в список.
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

	// "add", "list", ... - упоминаются по всей программе как в модуле main, так
	// и в этом. Представь, что тебе надо будет переименовать "add" в "append"
	// (требование руководства, которого пока нет). Ты будешь это делать в
	// нескольких местах - а мог бы в одном. Вынеси эти строки в константы
	// (для этого хорошо подходит файл types.go).
	tm.addEvent("add", tm.nextTaskID, text) // событие

	id := tm.nextTaskID
	tm.nextTaskID++

	return id, nil
} // добавление новых задач

func (tm *TaskManag) GetTask(id int) (*Task, error) {
	for _, task := range tm.tasks {
		if task.ID == id {
			// Потенциальный баг - возвращаешь указатель на временную переменную
			// task, а не на элемент в массиве. Если кто-то захочет
			// воспользоваться этим API (либо библиотекой, как угодно), то
			// велик шанс, что разработчик, увидя результатом функции GetTask()
			// указатель, подумает, что это именно элемент массива.
			// Контринтуитивно. Исправить.
			return &task, nil
		}
	}
	return nil, fmt.Errorf("Задача %d. не найдена", id)
} // func GetTask возвращает задачу по ID

// Замечание под звездочкой для двух геттеров ниже (AllTasks() и AllEvents()).
// Возвращая то, что возвращается сейчас, ты теряешь контроль над внутренним
// состоянием содержимого tasks и events. То есть внешний обработчик (модуль)
// может при помощи этих методов получить внутренние поля и менять их там
// элементы как ему вздумается. Когда таких модулей нет или он один - это
// неплохо. Но на продакшн коде это потенциальная угроза состоянию данных.
func (tm *TaskManag) AllTasks() []Task {
	return tm.tasks
} // func AllTasks возвращает все задачи

func (tm *TaskManag) AllEvents() []Event {
	return tm.events
}

func (tm *TaskManag) DeleteTask(id int) error {
	// На подумать: почему для хранения задач был выбран слайс (динамический
	// массив)? Какие плюсы и минусы были бы от решения с мапой?
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
