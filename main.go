# 1. Откройте файл main.go в редакторе
# 2. Удалите ВСЁ, кроме новой версии (верхней части):

# Оставить только это:
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AlexAgatskiy/go-todo-cli/internal/app"
)

type Task struct {
	id   int
	text string
	done bool
}
type TaskManag struct { // менеджер задач
	tasks  []Task
	nextID int
}

func main() {
	taskManager := app.TaskManager()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Привет вас в To-do list!")
	fmt.Println("Введите 'help' для вывода списка команд")

	for {
		fmt.Print("Введите команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода!")
			return
		}
		intext := scanner.Text()
		slova := strings.Fields(intext)

		if len(slova) == 0 {
			continue
		}

		cmd := slova[0]

		switch cmd {
		case "add":
			app.ProccesAdd(slova, taskManager)
		case "list":
			app.ProccesList(taskManager)
		case "done":
			app.ProccesDone(slova, taskManager)
		case "delete":
			app.ProccesDelete(slova, taskManager)
		case "events":
			app.ProccesEvents(taskManager)
		case "help":
			app.ProccesHelp()
		case "exit":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Printf("Неизвестная команда: '%s'\n", cmd)
			fmt.Println("Введите 'help' для списка команд")
		}
	}
}