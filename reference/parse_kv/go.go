import "strings"

type Entry struct{ Key, Value string }

func ParseKV(s string) ([]Entry, bool) {
	if s == "" {
		return []Entry{}, true
	}
	out := []Entry{}
	for _, e := range strings.Split(s, ";") {
		parts := strings.Split(e, "=")
		if len(parts) != 2 || parts[0] == "" {
			return nil, false
		}
		out = append(out, Entry{Key: parts[0], Value: parts[1]})
	}
	return out, true
}
