# Pacman Bubble Tea TUI

This project is a terminal user interface (TUI) application written in Go that provides a visual way to run `pacman` commands using your existing alias for `sudo pacman`. 

## Prerequisites

- **Go**: Make sure you have Go installed on your system. You can check if Go is installed by running:

    ```bash
    go version
    ```

    If it's not installed, follow the instructions on the [official Go website](https://golang.org/doc/install).

## Getting Started

### Step 1: Set Up Your Go Workspace

1. **Create a Directory**:

    ```bash
    mkdir -p ~/go/src/pacman-bubbletea
    cd ~/go/src/pacman-bubbletea
    ```

2. **Place Your Code**: Save your Go code in a file named `main.go` in the `pacman-bubbletea` directory.

### Step 2: Initialize the Go Module

To manage dependencies easily, initialize a Go module:

```bash
go mod init pacman-bubbletea

