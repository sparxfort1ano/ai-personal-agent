package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ai-personal-agent/internal/storage"
)

func main() {
	store := storage.NewStorage()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Добро пожаловать в ИИ-агент!")
	fmt.Println("Введите 'help' для списка команд или 'exit' для выхода.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]

		switch command {
		case "add":
			// Ожидаем формат: add YYYY-MM-DD Название события
			if len(parts) < 3 {
				fmt.Println("Ошибка: недостаточно аргументов. Использование: add <Дата> <Название>")
				continue
			}
			date := parts[1]
			title := strings.Join(parts[2:], " ")
			id := store.AddEvent(title, date)
			fmt.Printf("Событие '%s' на дату %s успешно добавлено (ID: %d)\n", title, date, id)

		case "list":
			events := store.ListEvents()
			if len(events) == 0 {
				fmt.Println("Список событий пуст.")
			} else {
				fmt.Println("Ваши события:")
				for _, ev := range events {
					fmt.Printf("[%d] %s — %s\n", ev.ID, ev.Date, ev.Title)
				}
			}

		case "delete":
			if len(parts) != 2 {
				fmt.Println("Ошибка: неверное число аргументов. Использование: delete <ID>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Ошибка: ID должен быть числом.")
				continue
			}
			err = store.DeleteEvent(id)
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Событие %d успешно удалено.\n", id)
			}

		case "help":
			fmt.Println("Доступные команды:")
			fmt.Println("  add <Дата> <Название> - Добавить новое событие")
			fmt.Println("  list                  - Показать все события")
			fmt.Println("  delete <ID>           - Удалить событие по его ID")
			fmt.Println("  help                  - Показать это меню")
			fmt.Println("  exit                  - Выйти из программы")

		case "exit":
			fmt.Println("Завершение работы программы.")
			os.Exit(0)

		default:
			fmt.Printf("Неизвестная команда: %s. Введите 'help' для справки.\n", command)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения ввода: %v\n", err)
		os.Exit(1)
	}
}
