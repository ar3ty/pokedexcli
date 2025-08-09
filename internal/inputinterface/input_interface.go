package inputinterface

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type Reader struct {
	keywords      []string
	history       []string
	current       int
	historyBuffer []rune
}

type Input struct {
	buffer []rune
	cursor int
}

func Init(words []string) Reader {
	return Reader{
		keywords:      words,
		history:       []string{""},
		current:       0,
		historyBuffer: nil,
	}
}

func commonPrefix(strs []string) string {
	if len(strs) < 2 {
		return ""
	}

	shortest := strs[0]
	for _, str := range strs {
		if len(shortest) > len(str) {
			shortest = str
		}
	}

	for i, char := range shortest {
		for _, str := range strs {
			if rune(str[i]) != char {
				return shortest[:i]
			}
		}
	}

	return shortest
}

func (r *Reader) autoComplete(prefix string, input *Input) {
	matches := []string{}
	for _, option := range r.keywords {
		if strings.HasPrefix(option, prefix) {
			matches = append(matches, option)
		}
	}

	if len(matches) == 0 {
		return
	}
	if len(matches) == 1 {
		fmt.Print(strings.Repeat("\b", len(prefix)))
		input.cursor -= len(prefix) - len(matches[0])
		fmt.Printf("%s", matches[0])
		input.buffer = []rune(matches[0])
	}

	if common := commonPrefix(matches); common != "" {
		fmt.Print(strings.Repeat("\b", len(prefix)))
		input.cursor -= len(prefix) - len(common)
		fmt.Printf("%s", common)
		input.buffer = []rune(common)
	}
}

func (r *Reader) lineRedraw(toDraw string, input *Input) {
	if input.cursor > 0 {
		fmt.Printf("\x1b[%dD", input.cursor)
	}
	fmt.Print("\x1b[K")
	fmt.Printf("%s", toDraw)
	input.cursor = len(toDraw)
}

func (r *Reader) Read() (string, error) {
	oldTerminal, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldTerminal)

	input := Input{
		buffer: []rune{},
		cursor: 0,
	}

	for {
		buf := make([]byte, 4)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return "", err
		}

		char := rune(buf[0])
		switch char {
		case '\r':
			fmt.Print("\n\r")
			value := string(input.buffer)
			if input.cursor > 0 && len(strings.Fields(value)) > 0 {
				r.history[r.current] = value
				if r.current == len(r.history)-1 {
					r.history = append(r.history, "")
				}
				r.current = len(r.history) - 1
			}
			return value, nil
		case '\n':
			fmt.Print("\n\r")
			value := string(input.buffer)
			if input.cursor > 0 && len(strings.Fields(value)) > 0 {
				r.history[r.current] = value
				if r.current == len(r.history)-1 {
					r.history = append(r.history, "")
				}
				r.current = len(r.history) - 1
			}
			return value, nil
		//tab is assigned to autocompletion
		case '\t':
			if len(input.buffer) == 0 {
				continue
			}
			if len(strings.Fields(string(input.buffer))) == 1 {
				prefix := string(input.buffer[:input.cursor])
				r.autoComplete(prefix, &input)
			}
		//control sequence
		case '\x1b':
			if n >= 3 && buf[1] == '[' {
				last := buf[2]
				switch last {
				//arrow up
				case 'A':
					if r.current > 0 {
						r.history[r.current] = string(input.buffer)
						r.current--
						recovered := r.history[r.current]
						r.lineRedraw(recovered, &input)
						input.buffer = []rune(recovered)
					}
				//arrow down
				case 'B':
					if r.current < len(r.history)-1 {
						r.history[r.current] = string(input.buffer)
						r.current++
						recovered := r.history[r.current]
						r.lineRedraw(recovered, &input)
						input.buffer = []rune(recovered)
					}
				//arrow right
				case 'C':
					if input.cursor < len(input.buffer) {
						fmt.Printf("\x1b[1C")
						input.cursor++
					}
				//arrow left
				case 'D':
					if input.cursor > 0 {
						fmt.Printf("\b")
						input.cursor--
					}
				//delete key
				case '3':
					if buf[3] == '~' && input.cursor < len(input.buffer) {
						fmt.Printf("%s ", string(input.buffer[input.cursor+1:]))
						fmt.Print(strings.Repeat("\b", len(input.buffer[input.cursor+1:])+1))
						input.buffer = append(input.buffer[:input.cursor], input.buffer[input.cursor+1:]...)
					}
				}
			}
		//backspace
		case '\x7f':
			if len(input.buffer) > 0 && input.cursor > 0 {
				fmt.Printf("\b%s ", string(input.buffer[input.cursor:]))
				fmt.Print(strings.Repeat("\b", len(input.buffer[input.cursor:])+1))
				input.buffer = append(input.buffer[:input.cursor-1], input.buffer[input.cursor:]...)
				input.cursor--
			}
		default:
			//copypaste handling cycle
			for i := 0; i < 4; i++ {
				if buf[i] == 0 {
					break
				}
				char := rune(buf[i])
				fmt.Print(string(char))
				fmt.Printf("%s", string(input.buffer[input.cursor:]))
				fmt.Print(strings.Repeat("\b", len(input.buffer[input.cursor:])))
				input.buffer = append(input.buffer[:input.cursor], append([]rune{char}, input.buffer[input.cursor:]...)...)
				input.cursor++
			}
		}
	}
}
