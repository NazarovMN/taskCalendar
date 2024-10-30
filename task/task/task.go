package task

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var FileName = "taskData" + ".json"

type Task struct {
	Name        string    `json:"name:"`
	Description string    `json:"description:"`
	DueDate     time.Time `json:"due_date:"`
	Priority    int       `json:"priority:"`
	Status      string    `json:"status:"`
}

func New(taskName, taskDescription, taskStatus string, taskDue time.Time, taskPriority int) (*Task, error) {
	if taskName == "" || taskDescription == "" {
		fmt.Println("Invalid Input. ")
		return &Task{}, errors.New("invalid input")
	}

	task := &Task{
		Name:        taskName,
		Description: taskDescription,
		Status:      taskStatus,
		DueDate:     taskDue,
		Priority:    taskPriority,
	}

	return task, nil
}

func (t *Task) UpdateName() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please Change task name: ")

	newName, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input", err)
		return
	}

	newName = strings.TrimSpace(newName)

	t.Name = newName

	err = t.Save()

	if err != nil {
		fmt.Println("Can't Save new info")
		return
	}

}

func (t *Task) UpdateDescription() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please Change description: ")

	newDescrioption, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reeading input", err)
		return
	}
	newDescrioption = strings.TrimSpace(newDescrioption)

	t.Description = newDescrioption

	err = t.Save()

	if err != nil {
		fmt.Println("Can't Save new info")
		return
	}
}

func (t *Task) UpdateDueDate() {
	var dateStr string
	fmt.Print("Please Change Due Date(DD-MM-YYYY): ")
	fmt.Scan(&dateStr)

	date, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		fmt.Println("Cant parse the time")
		return
	}

	t.DueDate = date

	err = t.Save()

	if err != nil {
		fmt.Println("Can't Save new info")
		return
	}
}

func (t *Task) UpdatePriority(priotirySt int) {
	fmt.Print("Change priority status(1-High, 2-Medium, 3-Low):", priotirySt)
	fmt.Scan(&priotirySt)
	if priotirySt >= 3 {
		priotirySt = 3
	}
	t.Priority = priotirySt

	err := t.Save()

	if err != nil {
		fmt.Println("Can't Save new info")
		return
	}
}

func (t Task) Save() error {
	fileName := FileName
	json, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (t *Task) ShowTask() (*Task, error) {
	fileName := FileName
	jsonchik, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("ty eblan failed to read from json")
		return nil, err
	}
	var task Task

	err = json.Unmarshal(jsonchik, &task)

	if err != nil {
		fmt.Println("Failed to read from json.")
		return nil, err
	}

	return &task, nil

}

func (t *Task) ShowAsJson() error {
	jsonData, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		fmt.Println("Can't output as json format")
		return err
	}
	fmt.Println(string(jsonData))

	return nil
}

func (u *Task) UpdateStatus() {

	var choice int
	fmt.Print("Choose your task status(1. Pending, 2. In Progress, 3. Completed) ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		u.Status = "Pending"
	case 2:
		u.Status = "In Progress"
	case 3:
		u.Status = "Completed"
	default:
		fmt.Println("Invalid choice.")
	}

	err := u.Save()

	if err != nil {
		fmt.Println("Can't Save new info")
		return
	}
}

func (t Task) OutputTaskDetails() {
	fmt.Printf("Task name: %v \nTask description: %v \nDate: %v \nPriority Status: %v\nStatus: %v\n", t.Name, t.Description, t.DueDate, t.Priority, t.Status)
}
