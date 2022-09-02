package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct{
	Task string
	Status bool
	Priority int64
	Description string
}

type TODO []Task

func ListAll(){
	byte,err:=os.ReadFile("DB/db.json")
	if err!=nil{
		log.Fatal("Cannot read data")
	}
	var tasks []Task
	json.Unmarshal(byte,&tasks)
	for i := 0; i < len(tasks); i++ {
		var prio string
		if tasks[i].Priority<5{
			prio="low"
		} else if tasks[i].Priority>5 {
			prio="high"
		} else{
			prio="medium"
		}
		fmt.Println(i+1,"> Task:",tasks[i].Task)
		fmt.Println("\tPriority:",prio)
		fmt.Println("\tStatus:",tasks[i].Status)
	}
}

func getTasks() ([]Task,error){
	byte,err:=os.ReadFile("DB/db.json")
	var data []Task
	if err!=nil{
		log.Fatal("Error accessing data")
	} else {
		json.Unmarshal(byte,&data)
	}
	return data,err
}
	
func UpdateDb(task []Task){
	input,err:=json.Marshal(task)
	if err!=nil{
		log.Fatal("Error writing data")
	}
	os.WriteFile("DB/db.json",input,0644)
}

func AddTask(task string,status bool,priority int64,Description string){
	var newTask Task
	newTask.Task=task
	newTask.Priority=priority
	newTask.Status=status
	newTask.Description=Description
	tasks,err:=getTasks()
	if err!=nil{
		log.Fatal("Error accessing data")
	}
	tasks = append(tasks, newTask)
	UpdateDb(tasks)
}

func Delete(index int64){
	tasks,err:=getTasks()
	if err!=nil{
		log.Fatal("Unable to access data")
	}
	if index<0 || index>int64(len(tasks)){
		log.Fatal("Index out of range")
	}
	tasks = append(tasks[0:index-1],tasks[index:]... )
	UpdateDb(tasks)
}

func Expand(index int64){
	tasks,err:=getTasks()
	if err!=nil{
		log.Fatal("Unable to access data")
	}
	if index<0 || index>int64(len(tasks)){
		log.Fatal("Index out of range")
	}
	var prio string
		if tasks[index-1].Priority<5{
			prio="low"
		} else if tasks[index-1].Priority>5 {
			prio="High"
		} else{
			prio="Medium"
		}
	task:=tasks[index-1]
	fmt.Println("Title:",task.Task)
	fmt.Println("Priority:",prio)
	fmt.Println("Status:",task.Status)
	fmt.Println("Description:",task.Description)
}