package properties

import "strings"

func shouldSkip(line string) bool {
	return startsWith("#", line) || isEmpty(line)
}

func isEmpty(line string) bool {
	return len(strings.TrimSpace(line)) == 0
}

func startsWith(prefix string, line string) bool {
	return strings.HasPrefix(line, prefix)
}

