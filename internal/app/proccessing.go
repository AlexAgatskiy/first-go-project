package app

import (
	"fmt"
	"strconv"
	"strings"
)

func ProccesAdd(slova []string, tm *TaskManag) {
	if len(slova) < 2 {
		fmt.Println("Ошибка: укажите текст задачи")
		fmt.Println("Пример: add Купить молоко")
		return
	}

	taskText := strings.Join(slova[1:], " ")
	id, err := tm.AddTask(taskText)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	fmt.Printf("Задача %d. добавлена\n", id)
}

func ProccesList(tm *TaskManag) {
	tasks := tm.AllTasks()

	if len(tasks) == 0 {
		fmt.Println("Нет задач")
		return
	}
	fmt.Println(".. Список задач ..")
	fmt.Printf("%-4s %-6s %s\n", "ID", "Status", "Задача")
	fmt.Println()

	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("%d.   [%s] %s\n", task.ID, status, task.Text)
	}
}

func ProccesEvents(tm *TaskManag) {
	events := tm.AllEvents()

	if len(events) == 0 {
		fmt.Println("История событий не указана")
		return
	}

	fmt.Println(".. Список всех событий ..")
	fmt.Printf("%-4s %-12s %-10s %-8s %-25s %s\n",
		"ID", "Time", "Действие", "Задача", "Текст", "Описание")
	fmt.Println(strings.Repeat("=", 90)) // пакет работы со строками и функция повтора знака "=" 90 раз

	for _, event := range events {
		timeStr := event.Timestamp.Format("15:04:05")

		actionName := event.Type
		description := ""

		switch event.Type {
		case "add":
			actionName = "добавление"
			description = fmt.Sprintf("Добавлена задача %d", event.TaskID) // string print format
		case "delete":
			actionName = "удаление"
			description = fmt.Sprintf("Удалена задача %d", event.TaskID)
		case "done":
			actionName = "выполнение"
			description = fmt.Sprintf("Задача %d выполнена", event.TaskID)
		}

		taskText := event.TaskText //обрезание задачи когда он длинный
		if len(taskText) > 20 {
			taskText = taskText[:17] + "..."
		}
		fmt.Printf("%-4d %-12s %-10s %-8d %-25s %s\n",
			event.ID,
			timeStr,
			actionName,
			event.TaskID,
			taskText,
			description)
	}
	fmt.Printf("\nВсего событий: %d\n", len(events))
}

func ProccesDone(slova []string, tm *TaskManag) {
	if len(slova) < 2 {
		fmt.Println("Ошибка: укажите номер задачи")
		fmt.Println("Пример: done 1")
		return
	}

	taskID, err := strconv.Atoi(slova[1])
	if err != nil {
		fmt.Println("Ошибка: номер задачи должен быть числом")
		return
	}
	success, err := tm.MarkAsDone(taskID)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	if success {
		fmt.Printf("Задача %d отмечена как выполненная \n", taskID)
	}
}

func ProccesDelete(slova []string, tm *TaskManag) {
	if len(slova) < 2 {
		fmt.Println("Ошибка: укажите номер задачи")
		fmt.Println("Пример: delete 1")
		return
	}
	taskID, err := strconv.Atoi(slova[1])
	if err != nil {
		fmt.Println("Ошибка: номер задачи должен быть числом")
		return
	}
	err = tm.DeleteTask(taskID)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Задача %d удалена\n", taskID)
}

func ProccesHelp() {
	fmt.Println("=== TO-DO List Справка ===")
	fmt.Println("")
	fmt.Println("Доступные команды:")
	fmt.Println("  add <текст>    - Добавить новую задачу")
	fmt.Println("  list           - Показать все задачи")
	fmt.Println("  done <id>      - Отметить задачу как выполненную")
	fmt.Println("  delete <id>    - Удалить задачу")
	fmt.Println("  events         - Показать список всех событий")
	fmt.Println("  help           - Показать эту справку")
	fmt.Println("  exit           - Выйти из программы")
	fmt.Println("")
	fmt.Println("Примеры использования:")
	fmt.Println("  add Купить хлеб и молоко, оплатить кредит")
	fmt.Println("  add Купить мороженое девушке")
	fmt.Println("  list           # Показать все задачи")
	fmt.Println("  done 1         # Отметить задачу #1 как выполненную")
	fmt.Println("  delete 2       # Удалить задачу #2")
	fmt.Println("  events         # Показать все события")
	fmt.Println("  help           # Показать справку")
	fmt.Println("  exit           # Выйти из программы")
	fmt.Println("")
	fmt.Println("Особенности:")
	fmt.Println("- Каждое действие сохраняется в историю событий")
	fmt.Println("- Команда 'events' показывает всю историю")
	fmt.Println("- События нумеруются и имеют временные метки")
}
