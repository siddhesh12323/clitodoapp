# Todo CLI Application

A simple command-line interface (CLI) application to manage your tasks. This app allows you to add, view, delete, and complete tasks.

## Features

- Add tasks to your list
- View incomplete or all tasks
- Mark tasks as complete
- Delete tasks
- Display human-readable timestamps (e.g., "2 days ago")

## How to Run the App Without Using `./todo`  

By default, after building the app, you run it using `./todo`. However, to make it executable system-wide using just `todo`, follow these steps:

### Step 1: Build the App

Run the following command in the root directory of the project to build the application:

```bash
go build -o todo
```

This creates an executable file named `todo`.

### Step 2: Move the Executable to a System-Wide Path

To run the app with `todo`, move the built executable to a directory included in your system's `$PATH`. Common directories include `/usr/local/bin` or `/usr/bin`.

```bash
sudo mv todo /usr/local/bin/
```

### Step 3: Test the Application

Once the executable is in the `$PATH`, you can run the application directly by typing:

```bash
todo
```

---

## Example Usage

### Display the Welcome Message
```bash
todo
```
Output:
```
Welcome Back to todo app! Run 'todo help' to get list of commands
```

### View Help
```bash
todo help
```

### Add a Task
```bash
todo add
```

### Mark a Task as Complete
```bash
todo complete
```

### Delete a Task
```bash
todo delete
```

---

## Development

To customize or add new features, edit the `cmd` folder where the commands are implemented. Add your custom commands and link them to the root command in `root.go`.

---

## License

This project is open-source. Feel free to modify and use it as needed!
