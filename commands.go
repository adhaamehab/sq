package sq

import "strings"

type Command string

const (
	PINGCMD    Command = "PING"
	ENQUEUECMD Command = "ENQUEUE"
	DEQUEUECMD Command = "DEQUEUE"
	PEEKCMD    Command = "PEEK"
	UnknownCMD Command = "UNKNOWN"
)

// parseCommand will parse the input string and return the command and arguments
func parseCommand(input string) (Command, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", nil
	}

	cmd := strToCommand(parts[0])
	args := parts[1:]

	return cmd, args
}

func strToCommand(str string) Command {
	switch str {
	case "PING":
		return PINGCMD
	case "ENQUEUE":
		return ENQUEUECMD
	case "DEQUEUE":
		return DEQUEUECMD
	case "PEEK":
		return PEEKCMD
	default:
		return UnknownCMD
	}
}
