package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var todos []string

func addTodos(arg string) {
	todos = append(todos, arg)
}

func delTodos(arg string) error {
	// пробуем преобразовать в число
	num, err := strconv.Atoi(strings.TrimSpace(arg))
	if err != nil {
		return fmt.Errorf("некорректный номер заметки: '%s'. Введите число", arg)
	}

	// пользователь видит список начиная с 1, а не с 0
	if num < 1 || num > len(todos) {
		return fmt.Errorf("заметки с номером %d не существует", num)
	}

	// удаляем заметку
	todos = append(todos[:num-1], todos[num:]...)
	return nil
}

func showTodos() {
	if len(todos) == 0 {
		fmt.Println("Список заметок пуст")
		return
	}
	for i, td := range todos {
		fmt.Println(i+1, td)
	}
}

func validCommand(todo string) (string, string, error) {
	sep := strings.Index(todo, " ")
	var cmd, arg string

	if sep == -1 {
		cmd = todo
		arg = ""
	} else {
		cmd = todo[:sep]
		arg = strings.TrimSpace(todo[sep+1:]) // убираем лишние пробелы
	}

	switch cmd {
	case "/add":
		if arg == "" {
			return cmd, arg, fmt.Errorf("после команды /add нужно ввести текст заметки")
		}
		addTodos(arg)
		fmt.Println("Заметка была успешно добавлена")
	case "/del":
		if arg == "" {
			return cmd, arg, fmt.Errorf("после команды /del нужно указать номер заметки")
		}
		if err := delTodos(arg); err != nil {
			return cmd, arg, err
		}
		fmt.Println("Заметка была успешно удалена")
	case "/show":
		showTodos()
	case "/exit":
		fmt.Println("Выход из программы...")
		os.Exit(0)
	default:
		return cmd, arg, fmt.Errorf("неизвестная команда: %s", cmd)
	}
	return cmd, arg, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("----------------------ToDo----------------------")
	fmt.Println("1. Добавить заметку: /add [текст заметки]")
	fmt.Println("2. Удалить заметку: /del [номер заметки]")
	fmt.Println("3. Посмотреть заметки: /show")
	fmt.Println("4. Выйти из программы: /exit")
	fmt.Println("------------------------------------------------")
	for {
		if valid := scanner.Scan(); !valid {
			fmt.Println("Ошибка ввода. Завершение работы.")
			return
		}

		todo := scanner.Text()
		_, _, err := validCommand(todo)
		if err != nil {
			fmt.Println("Ошибка:", err)
		}
	}
}
