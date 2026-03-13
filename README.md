## Todo CLI App

### Goal

Create a CLI application for managing tasks in the terminal.

```
$ tasks
```

This application allows users to manage tasks using command-line commands.
Tasks are stored in a JSON file and can be created, listed, completed, and deleted.

---

## Features

* Add new tasks
* List tasks
* Mark tasks as complete
* Delete tasks
* Store tasks in JSON file
* CLI built using Cobra
* Friendly CLI output using tabwriter
* Relative time display using timediff

---

## Commands

### Add Task

Add a new task with description.

```
$ tasks add <description>
```

Example:

```
$ tasks add "Tidy my desk"
```

Add a new task with description and due date. 

```
$ tasks add <description> --due <due-date>
```

Example: 

```
$ tasks add "Start new go project" --due "23-03-2026" 
```

---

### List Tasks

Show all uncompleted tasks.

```
$ tasks list
```

Example output:

```
ID    Task                Created
1     Tidy my desk        a minute ago
3     Buy groceries       a few seconds ago
```

Show all completed tasks:

```
$ tasks completed
```

Example output:

```
ID    Task                Created
2     Tidy my desk        a minute ago
34    Buy groceries       a few seconds ago
```

Show all tasks including completed:

```
$ tasks list -a
```

Example:

```
ID    Task                Created          Status
1     Tidy my desk        2 minutes ago    false
2     Write docs          1 minute ago     true
3     Buy groceries       few sec ago      false
```

---

### Complete Task

Mark a task as complete.

```
$ tasks complete <taskid>
```

Example:

```
$ tasks complete 2
```

---
### Edit Task 

Allow users to modify a task description.
```
$ tasks edit <id> "New description"
```
Example:

```
$ tasks edit 3 "Buy groceries and fruits"
```
---

### Delete Task

Delete a task from storage.

```
$ tasks delete <taskid>
```

Example:

```
$ tasks delete 3
```

---

## Data Storage

Tasks are stored in a JSON file.

Example:

```json
[
  {
    "id": 1,
    "description": "My new task",
    "createdAt": "2024-07-27T16:45:19Z",
    "isComplete": false
  }
]
```

---

## Packages Used

* `github.com/spf13/cobra` → CLI commands
* `encoding/json` → storing tasks
* `strconv` → type conversion
* `text/tabwriter` → formatted output
* `os` → file handling
* `github.com/mergestat/timediff` → friendly time output

---

## How to Run

```
go build -o tasks
./tasks add "My task"
./tasks list
./tasks complete 1
./tasks delete 1
```

---

## Conclusion

This project demonstrates building a CLI application in Go using Cobra, file storage, and command-based task management.

