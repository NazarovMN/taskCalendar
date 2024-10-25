package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

func (t *Task) UpdateName(newName string) {
	fmt.Print("Please Change task name: ", newName)
	fmt.Scan(&newName)
	t.Name = newName

}

func (t *Task) UpdateDescription(newDescrioption string) {
	fmt.Print("Please Change description: ", newDescrioption)
	fmt.Scan(&newDescrioption)
	t.Description = newDescrioption
}

func (t *Task) UpdateDueDate(date time.Time) {
	fmt.Print("Please Change Due Date(MM-DD-YYYY): ", date)
	fmt.Scan(&date)
	t.DueDate = date
}

func (t *Task) UpdatePriority(priotirySt int) {
	fmt.Print("Change priority status: ", priotirySt)
	fmt.Scan(&priotirySt)
	t.Priority = priotirySt
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
	var task *Task

	err = json.Unmarshal(jsonchik, task)

	if err != nil {
		fmt.Println("Failed to read from json.")
		return nil, err
	}

	return task, nil

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

}

func (t Task) OutputTaskDetails() {
	fmt.Printf("Task name: %v \nTask description: %v \nDate: %v \nPriority Status: %v\nStatus: %v\n", t.Name, t.Description, t.DueDate, t.Priority, t.Status)
}
