# INSTALL.md

# 🛠 Installing Todo CLI

Follow these instructions to install **Todo CLI** on your system.

---

## 1. Build from Source

Make sure you have **Go 1.20+** installed.

```bash
# Clone the repository
git clone https://github.com/kodwo-essel/todo-cli.git
cd todo-cli

# Build the binary for your platform
go build -o todo
```

**Optional:** Move the binary to a directory in your `PATH` to use globally:

```bash
sudo mv todo /usr/local/bin/
```

Verify installation:

```bash
todo --version
```

---

## 2. Prebuilt Binaries (Optional)

You can download prebuilt binaries for your platform from the [Releases](https://github.com/kodwo-essel/todo-cli/releases) page:

| Platform | File                                                    |
| -------- | ------------------------------------------------------- |
| Linux    | `todo-linux-amd64.tar.gz`                               |
| macOS    | `todo-darwin-amd64.tar.gz` / `todo-darwin-arm64.tar.gz` |
| Windows  | `todo-windows-amd64.zip`                                |

Extract the archive and move the executable to a folder in your `PATH`:

```bash
# Example for Linux
tar -xzf todo-linux-amd64.tar.gz
sudo mv todo /usr/local/bin/
```

---

## 3. Verify Installation

Run:

```bash
todo
```

You should see the welcome screen:

```
📝 Welcome to todo!
Type 'todo help' to see available commands.
```

---
