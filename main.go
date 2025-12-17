package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	id   int
	text string
	done bool
}

func main() {
	tasks := []Task{} // динамический массив для записи задач
	inId := 1         //
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Приветствую вас в Ту-Ду лист!")
	fmt.Println("Вееди 'Help' для вывод списка команд ")

	for {
		fmt.Print("Введите команду:")

		if ok := scanner.Scan(); !ok {
			// проверяем на ошибку
			fmt.Println("Ошибка ввода!")
			return
		}
		intext := scanner.Text() // текст ввода пользователя

		slova := strings.Fields(intext) // разибиваем ввод отдельные слова по пробелам

		if len(slova) == 0 { // пропуск после пустого ввода
			continue
		}
		cmd := slova[0] // первое слово=команда

		switch cmd {
		case "add":

			if len(slova) > 1 { // новая задача
				taskText := strings.Join(slova[1:], " ")

				nwTask := Task{
					id:   inId,
					text: taskText,
					done: false,
				}
				tasks = append(tasks, nwTask) // добавляем задачу в список

				fmt.Printf("Задача #%d доабвлена\n", inId) // пишем что добавили задачу

				inId++ // инкрементируем номер следущей задачи
			} else {
				fmt.Println("Задача не добавлена")
				fmt.Println("Ошибка ввода: нужен текст задачи")
			}

		case "list":
			if len(tasks) == 0 { // проверка и показатель задач
				fmt.Println("Нет задач")
			} else {
				fmt.Println("--Список задач--")
				fmt.Printf("%-4s %-6s %s\n", "ID", "Status", "Task") // %-4s(добавить в поле 4 символа) %-6s(добавить в поле 6 символов [x])
				fmt.Println("")

				for _, task := range tasks {
					status := " "
					if task.done {
						status = "x"
					}
					fmt.Printf("%d.  [%s] %s\n", task.id, status, task.text)
				}
			}
		case "events":
			fmt.Println("--Доступные команды--")
			fmt.Println("add <текст> Добавляет команды")
			fmt.Println("lits        Показывает добавленные команды")
			fmt.Println("done        Показывает выполненные задачи")
			fmt.Println("delete      Удаляет выбраннные задачи")
			fmt.Println("events      Показывает все команды")
			fmt.Println("exit        Выход из программы")

		case "help":
			fmt.Println("TO-DO list справка")
			fmt.Println("")
			fmt.Println("У каждой задачи свой ID")
			fmt.Println("")
			fmt.Println("Примеры:")
			fmt.Println(" add Добавить Лук ")
			fmt.Println(" list")
			fmt.Println(" done 1")
			fmt.Println(" delete 3")
		case "delete":
			if len(slova) > 1 { // Удаляем задачу
				taskID, conOk := strconv.Atoi(slova[1]) // конвектируем второй индекс в число
				if conOk != nil {                       // Ошибка конвектора
					fmt.Println("Error: номер задачи должен быть числом")
				} else {
					fnd := false // ищем задачу по номеру
					for i, task := range tasks {
						if task.id == taskID {
							fnd = true                                // удаляем задачу из списка
							tasks = append(tasks[:i], tasks[i+1:]...) // (...) оператор распоковки среза в отдельные элементы
							fmt.Printf("Задача %d. удалена\n", taskID)
							break
						}
					}
					if !fnd {
						fmt.Printf("Задача %d. не найдена\n", taskID)
					}
				}
			} else {
				fmt.Printf("Error: укажите номер задачи")
				fmt.Println("Например: delete 3")
			}
		case "done":
			if len(slova) > 1 { // берем и помечаем таску как выполненную
				taskID, conOk := strconv.Atoi(slova[1]) // Второй индекс переобразовать в число
				if conOk != nil {                       // если не конвертнуло в число
					fmt.Println("Error: номер задачи должен быть числом")
				} else {
					fnd := false
					for i := range tasks {
						if tasks[i].id == taskID {
							fnd = true
							if tasks[i].done {
								fmt.Printf("Задача %d. уже выполненная\n", taskID)
							} else {
								tasks[i].done = true
								fmt.Printf("Задача %d. отмечена как выполненная\n", taskID)
							}
							break
						}
					}
					if !fnd {
						fmt.Printf("Задача %d. не найдена\n", taskID)
					}
				}
			} else {
				fmt.Println("Error: укажите номер задачи")
				fmt.Println("Пример: done 3")
			}
		case "exit":
			fmt.Println("До встречи!")
			return
		default:
			fmt.Printf("Неизвестная команда: '%s'\n", cmd)
			fmt.Println("Введите 'help' для списка команд")
		}
	}
}
