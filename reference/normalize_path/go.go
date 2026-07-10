import "strings"

func NormalizePath(s string) (string, bool) {
	if !strings.HasPrefix(s, "/") {
		return "", false
	}
	stack := []string{}
	for _, c := range strings.Split(s, "/") {
		switch c {
		case "", ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, c)
		}
	}
	return "/" + strings.Join(stack, "/"), true
}
