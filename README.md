# 📝 Todo CLI

A minimal, beautiful, and fast command-line task manager. Manage your tasks from your terminal with a simple, elegant interface and local SQLite storage.

✨ Features:

* Add, list, complete, and delete tasks
* Track pending, completed, and overdue tasks
* Colorful and easy-to-read output
* Lightweight SQLite backend — no external dependencies

---

## Installation

### 1. Build from source

Make sure you have **Go 1.20+** installed.

```bash
# Clone the repository
git clone https://github.com/kodwo-essel/todo-cli.git
cd todo-cli

# Build the binary for your platform
go build -o todo
```

**Optional:** Move it to your system path to use globally:

```bash
sudo mv todo /usr/local/bin/
```

### 2. Prebuilt binaries

*(Optional — you can provide prebuilt binaries for Linux, macOS, Windows here if you want.)*

---

## Usage

Run `todo` to see the welcome screen:

```bash
$ todo
📝 Welcome to todo!
Type 'todo help' to see available commands.
Running on linux
```

See all commands:

```bash
$ todo --help
```

### Common Commands

* **Add a task**

```bash
todo add "Buy groceries" --priority high
todo add "Finish report" --due tomorrow
```

* **List tasks**

```bash
todo list
todo list --status pending
todo list --status completed
```

* **Mark a task as complete**

```bash
todo complete 3
```

* **Delete a task**

```bash
todo delete 5
```

* **Generate shell completions**

```bash
todo completion bash   # or zsh, fish, powershell
```

---

## Task Status

Your tasks are displayed with **status and color**:

* 🟡 Pending
* 🔴 Overdue
* ✅ Completed

---

## Development

To contribute:

```bash
git clone https://github.com/kodwo-essel/todo-cli.git
cd todo-cli
go build -o todo
```

Then make your changes, build, and test.

---

## License

MIT License © 2025 [Your Name]


