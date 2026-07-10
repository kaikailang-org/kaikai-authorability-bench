import "fmt"
import "strings"

type LogLine struct{ Level, Message string }

func ParseLine(s string) (LogLine, bool) {
	for _, lvl := range []string{"INFO", "WARN", "ERROR"} {
		if rest, found := strings.CutPrefix(s, lvl+": "); found {
			return LogLine{lvl, rest}, true
		}
	}
	return LogLine{}, false
}

func CountLevel(lines []string, level string) (int64, bool) {
	if level != "INFO" && level != "WARN" && level != "ERROR" {
		return 0, false
	}
	var n int64 = 0
	for _, s := range lines {
		l, ok := ParseLine(s)
		if !ok {
			return 0, false
		}
		if l.Level == level {
			n++
		}
	}
	return n, true
}

func Summary(lines []string) (string, bool) {
	i, ok1 := CountLevel(lines, "INFO")
	w, ok2 := CountLevel(lines, "WARN")
	e, ok3 := CountLevel(lines, "ERROR")
	if !ok1 || !ok2 || !ok3 {
		return "", false
	}
	return fmt.Sprintf("INFO=%d WARN=%d ERROR=%d", i, w, e), true
}
