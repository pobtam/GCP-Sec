package utils

import "strings"

// ExtractPositional finds the first non-flag positional argument from an args list,
// returning it and a new args slice with the positional argument removed.
// knownValueFlags is the set of flag names (with dashes) that take a value argument
// (so their values are not mistaken for the positional argument).
//
// Example:
//
//	ExtractPositional(
//	    []string{"findings.csv", "--output", "report.md", "-v"},
//	    map[string]bool{"--output": true, "-o": true},
//	)
//	→ "findings.csv", ["--output", "report.md", "-v"]
func ExtractPositional(args []string, knownValueFlags map[string]bool) (positional string, flagArgs []string) {
	foundPositional := false
	skipNext := false

	for _, arg := range args {
		if skipNext {
			// This arg is a value for the previous flag
			flagArgs = append(flagArgs, arg)
			skipNext = false
			continue
		}

		if strings.HasPrefix(arg, "-") {
			flagArgs = append(flagArgs, arg)
			// If this flag takes a value and doesn't use = syntax, skip next arg
			name := strings.SplitN(arg, "=", 2)[0] // strip =value if present
			if !strings.Contains(arg, "=") && knownValueFlags[name] {
				skipNext = true
			}
			continue
		}

		// Non-flag argument
		if !foundPositional {
			positional = arg
			foundPositional = true
		} else {
			// Extra positional args — pass through as-is (shouldn't normally happen)
			flagArgs = append(flagArgs, arg)
		}
	}
	return positional, flagArgs
}
