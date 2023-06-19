package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// IsInteractive contains the interactive terminal state set by
// DetectInteractive at init() time.
var IsInteractive bool

func init() {
	IsInteractive = DetectInteractive()
}

// DetectInteractive returns true if the output is to an interactive
// terminal (not piped in any way).
func DetectInteractive() bool {
	if f, _ := os.Stdout.Stat(); (f.Mode() & os.ModeCharDevice) != 0 {
		return true
	}
	return false
}

// Prompt prints the given message reads the string by calling Read. The
// argument signature is identical as that passed to fmt.Printf().
func Prompt(form string, args ...any) string {
	fmt.Printf(form, args...)
	return Read()
}

// PromptHidden prints the given message if the terminal IsInteractive
// and reads the string by calling ReadHidden (which does not echo to
// the screen). The argument signature is identical and passed to to
// fmt.Printf().
func PromptHidden(form string, args ...any) string {
	fmt.Printf(form, args...)
	return ReadHidden()
}

// Read reads a single line of input and chomps the \r?\n. Also see
// ReadHidden.
func Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// ReadHidden disables the cursor and echoing to the screen and reads
// a single line of input. Leading and trailing whitespace are removed.
// Also see Read.
func ReadHidden() string {
	byt, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(byt))
}
