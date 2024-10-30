package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"example.com/taskerville/task"
)

func main() {
	fmt.Println("Welcome to the best task app ever!!!")

	for {
		fmt.Println("What do you want to do: ")
		fmt.Println("1. Add the task.")
		fmt.Println("2. Check your tasks.")
		fmt.Println("3. Update task")
		fmt.Println("4.Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			name := getUserInput("Give task name: ")
			description := getUserInput("Give the Description: ")
			status := getUserInput("Give the status: ")
			date := getUserInput("Name the date(DD-MM-YYYY): ")

			dueDate, err := time.Parse(("02-01-2006"), date)

			if err != nil {
				fmt.Println("cannot parse date format")
				return
			}

			priorityStr := getUserInput("Priority Status(1-High, 2-Medium, 3-Low): ")

			priority, err := strconv.Atoi(priorityStr)

			if err != nil {
				fmt.Println("can't parse to int")
				return
			}

			task, err := task.New(name, description, status, dueDate, priority)

			if err != nil {
				fmt.Println(err)
				return
			}

			task.UpdateStatus()
			task.OutputTaskDetails()

			err = task.Save()
			if err != nil {
				fmt.Println(err)
				return
			}

		case 2:
			var task task.Task
			taski, err := task.ShowTask()

			if err != nil {
				fmt.Println(err)
				return
			}

			err = taski.ShowAsJson()

			if err != nil {
				fmt.Println(err)
				return
			}

		case 3:
			task, err := (&task.Task{}).ShowTask()

			if err != nil {
				fmt.Println(err)
				return
			}

			var choice int
			fmt.Printf("choose task to update:\n 1)%v\n 2)%v\n 3)%v\n 4)%v\n 5)%v \n", task.Name, task.Description, task.DueDate, task.Priority, task.Status)
			fmt.Scan(&choice)
			if choice == 1 {
				task.UpdateName()

			} else if choice == 2 {
				task.UpdateDescription()
			} else if choice == 3 {
				task.UpdateDueDate()
			} else if choice == 4 {
				var priority int
				task.UpdatePriority(priority)

			} else if choice == 5 {
				task.UpdateStatus()
			} else {
				fmt.Println("please choose correct one ")
			}
		default:
			fmt.Println("Bye-Bye")
			return

		}
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
