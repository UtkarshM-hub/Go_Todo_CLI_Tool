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
}

type TODO []Task

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

func AddTask(task string,status bool,priority int64){
	var newTask Task
	newTask.Task=task
	newTask.Priority=priority
	newTask.Status=status
	tasks,err:=getTasks()
	fmt.Println(tasks)
	if err!=nil{
		log.Fatal("Error accessing data")
	}
	tasks = append(tasks, newTask)
	UpdateDb(tasks)
}
