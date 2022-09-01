# Go_Todo_CLI_Tool

### To Use this Project
clone this repo and run following command in your terminal which will build the project and create an executable file out of it
```
go build .
```

## Add Task
following command is used to Add task
```
todo add -n <NAME_OF_THE_TASK> -p <PRIORITY_OF_THE_TASK> 
```
for example
```
todo add -n "First task" -p 10
```
here ```-n``` tag represents name of the task and ```-p``` represents priority.

this will store the json data into the db.json file like you can see in the following image
![Screenshot (90)](https://user-images.githubusercontent.com/70505181/187959847-f50b244e-0df3-4d96-9a18-af988fd97e0d.png)
