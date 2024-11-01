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

This creates a go.mod file in your project directory.
Step 3: Install Required Dependencies

Install the necessary packages used in the application:

bash

go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/bubbles/progress
go get github.com/charmbracelet/bubbles/spinner
go get github.com/charmbracelet/lipgloss

Step 4: Build the Program

Compile your Go program with the following command:

bash

go build -o pacman-bubbletea

This creates an executable named pacman-bubbletea in the current directory.
Step 5: Make It Callable from Zsh

    Move the Executable to a Directory in Your PATH:

    bash

mkdir -p ~/bin
cp pacman-bubbletea ~/bin/

Update Your PATH: If you added ~/bin, ensure it is in your PATH. Add the following line to your ~/.zshrc file:

bash

export PATH="$HOME/bin:$PATH"

After adding this, reload your Zsh configuration:

bash

    source ~/.zshrc

Step 6: Run the Program

Now you can run your program directly from the terminal using:

bash

pacman-bubbletea

Step 7: Customize Usage (Optional)

To invoke your program using your alias sp, create a simple script:

    Create the sp Script:

    bash

touch ~/bin/sp
chmod +x ~/bin/sp

Edit the sp File: Add the following lines to the sp file:

bash

    #!/bin/bash
    # Call the Go TUI program with any arguments
    ~/bin/pacman-bubbletea "$@"

Now you can use sp in your terminal to run your Go program.
