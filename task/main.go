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

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			name := getUserInput("Give task name: ")
			description := getUserInput("Give the Description: ")
			status := getUserInput("Give the status: ")
			date := getUserInput("Name the date(MM-DD-YYYY): ")

			dueDate, err := time.Parse(("01-02-2006"), date)

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

			fmt.Println(taski)

		case 3:
			var t *task.Task

			task, err := t.ShowTask()

			if err != nil {
				fmt.Println(err)
				return
			}

			var choice int
			fmt.Printf("choose task to update:\n 1)%v\n 2)%v\n 3)%v\n 4)%v\n 5)%v \n", task.Name, task.Description, task.DueDate, task.Priority, task.Status)
			fmt.Scan(&choice)
			if choice == 1 {
				var name string
				t.UpdateName(name)
			} else if choice == 2 {
				var description string
				t.UpdateDescription(description)
			} else if choice == 3 {
				var date time.Time
				t.UpdateDueDate(date)
			} else if choice == 4 {
				var priority int
				t.UpdatePriority(priority)

			} else if choice == 5 {
				t.UpdateStatus()
			} else {
				fmt.Println("please choose correct one ")
			}

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
