package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const FILE_NAME = "tasks.json"

type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"createdAt"`
	IsComplete bool `json:"isComplete"`
	DueDate *time.Time `json:"dueDate"`
}

func ListTasks(filename string) ([]Task, error) {
	var tasks []Task

    fileData, err := os.ReadFile(filename)

    if err != nil {
        return tasks, fmt.Errorf("failed to read file: %w", err)
    }

    err = json.Unmarshal(fileData, &tasks) 

    if err != nil {
        return tasks, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    return tasks, nil
}

func ListPendingTasks(filename string) ([]Task, error) {
	var tasks []Task

    fileData, err := os.ReadFile(filename)

    if err != nil {
        return tasks, fmt.Errorf("failed to read file: %w", err)
    }

    err = json.Unmarshal(fileData, &tasks) 

    if err != nil {
        return tasks, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    var pendingTasks []Task

	for _ , task := range tasks {
		if(task.IsComplete == false) {
			pendingTasks = append(pendingTasks,task)
		}
	}

    return pendingTasks, nil
}

func ListCompleteTasks(filename string) ([]Task, error) {
	var tasks []Task

    fileData, err := os.ReadFile(filename)

    if err != nil {
        return tasks, fmt.Errorf("failed to read file: %w", err)
    }

    err = json.Unmarshal(fileData, &tasks) 

    if err != nil {
        return tasks, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

	var completedTasks []Task

	for _ , task := range tasks {
		if(task.IsComplete == true) {
			completedTasks = append(completedTasks,task)
		}
	}

    return completedTasks, nil
}

func AddTask(description string,due_date *time.Time) (bool, error) {
	tasks,err := ListTasks(FILE_NAME);
	var due *time.Time
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false, fmt.Errorf("failed to read file: %w", err)
	}
	if(due_date != nil){
		due = due_date
	}else{
		due = nil
	}
	newTask := Task{tasks[len(tasks) - 1].ID + 1,description,time.Now(),false,due}
	tasks = append(tasks,newTask)

	jsonData, err := json.MarshalIndent(tasks, "", "  ")

	err = os.WriteFile(FILE_NAME, jsonData, 0644) 
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false , fmt.Errorf("Error writing to file: %w", err)
	}
	return true,nil
}

func DeleteTask(id int)(bool,error){
	tasks,err := ListTasks(FILE_NAME);
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false, fmt.Errorf("failed to read file: %w", err)
	}
	var newTasks []Task
	for _ , task :=range tasks{
		if task.ID != id{
			newTasks = append(newTasks, task)
		}
	}

	jsonData, err := json.MarshalIndent(newTasks, "", "  ")

	err = os.WriteFile(FILE_NAME, jsonData, 0644) 
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false , fmt.Errorf("Error writing to file: %w", err)
	}
	return true,nil
}

func UpdateTask(id int,description string)(bool,error){
	tasks,err := ListTasks(FILE_NAME);
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false, fmt.Errorf("failed to read file: %w", err)
	}
	for index , task :=range tasks{
		if task.ID == id{
			tasks[index].Description = description
		}
	}

	jsonData, err := json.MarshalIndent(tasks, "", "  ")

	err = os.WriteFile(FILE_NAME, jsonData, 0644) 
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false , fmt.Errorf("Error writing to file: %w", err)
	}
	return true,nil
}

func CompleteTask(id int,)(bool,error){
	tasks,err := ListTasks(FILE_NAME);
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false, fmt.Errorf("failed to read file: %w", err)
	}
	for index , task :=range tasks{
		if task.ID == id{
			tasks[index].IsComplete = true
		}
	}

	jsonData, err := json.MarshalIndent(tasks, "", "  ")

	err = os.WriteFile(FILE_NAME, jsonData, 0644) 
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return false , fmt.Errorf("Error writing to file: %w", err)
	}
	return true,nil
}