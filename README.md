# Go_Todo_CLI_Tool

### To Use this Project
clone this repo and run following command in your terminal which will build the project and create an executable file out of it
```
go build .
```

## Add Task
following command is used to Add task
```
todo add -n <NAME_OF_THE_TASK> -p <PRIORITY_OF_THE_TASK> -d <DESCRIPTION>
```
here ```-n``` tag is used to specify name of the task and ```-p``` for priority.
The lowest priority a task can have is 0 and the highest is 10. For medium priority your can use 5.
<br/>
NOTE! The ```-d``` flag is optional
<br/>
An example of Add command
```
todo add -n "First task" -p 10 -d "Just an example"
```
this will store the json data into the db.json file like you can see in the following image
![Screenshot (93)](https://user-images.githubusercontent.com/70505181/188202962-48da8742-700d-4f41-a180-2b082d5c47b2.png)


## Delete Task
```
todo del -i 1
```
```-i``` is used to specify the index

## List All Tasks
```
todo list
```
![Screenshot (95)](https://user-images.githubusercontent.com/70505181/188274147-10806dd7-909a-4f29-82c1-81e392e6ae27.png)


## Expand a Task
```
todo exp -i 2
```
The ```-i``` is used to specify index 
it shows following output
<br/>
![Screenshot (97)](https://user-images.githubusercontent.com/70505181/188274248-574ad738-2b9c-4179-bf82-9e1d0f70e928.png)


## Mark Task as Complete
```
todo comp -i 1
```
```-i``` is used to specify the index of the task you want to mark as complete

## To Clear all Tasks
```
todo clear
```
