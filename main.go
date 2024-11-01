package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	commands []string
	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
}

var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle           = lipgloss.NewStyle().Margin(1, 2)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

func newModel() model {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return model{
		commands: []string{"update", "install package1", "install package2"}, // Example commands, modify as needed
		spinner:  s,
		progress: p,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(runPacman(m.commands[m.index]), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case installedPkgMsg:
		cmd := m.commands[m.index]
		if m.index >= len(m.commands)-1 {
			// Everything's done. We're finished!
			m.done = true
			return m, tea.Sequence(
				tea.Printf("%s %s", checkMark, cmd), // print the last success message
				tea.Quit,                            // exit the program
			)
		}

		// Update progress bar
		m.index++
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.commands)))

		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s %s", checkMark, cmd),     // print success message above our program
			runPacman(m.commands[m.index]), // run the next command
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
	}
	return m, nil
}

func (m model) View() string {
	n := len(m.commands)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Done! Processed %d commands.\n", n))
	}

	cmdCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+cmdCount))

	cmd := currentPkgNameStyle.Render(m.commands[m.index])
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Executing " + cmd)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+cmdCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + cmdCount
}

type installedPkgMsg string

// This function runs pacman commands using the alias 'sp'
func runPacman(command string) tea.Cmd {
	cmd := exec.Command("bash", "-c", "sp "+command) // Use your alias 'sp'
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	go func() {
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			fmt.Printf("Error output: %s\n", errOut.String())
			return
		}
		// Send a message after command completion
		tea.Send(installedPkgMsg(command))
	}()

	// Simulate a progress wait time (for demo purposes)
	time.Sleep(1 * time.Second) // Adjust time as necessary
	return nil
}

// Example package commands; modify as needed for your use case
func getCommands() []string {
	return []string{"-Syu", "-U package1", "-S package2"} // Replace with actual commands if needed
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	if _, err := tea.NewProgram(newModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

