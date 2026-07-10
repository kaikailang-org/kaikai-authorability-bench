import "strings"

type Param struct{ Key, Value string }

func ParseQuery(s string) ([]Param, bool) {
	if s == "" {
		return []Param{}, true
	}
	out := []Param{}
	for _, piece := range strings.Split(s, "&") {
		k, v, found := strings.Cut(piece, "=")
		if !found || k == "" {
			return nil, false
		}
		out = append(out, Param{k, v})
	}
	return out, true
}

func GetParam(ps []Param, k string) (string, bool) {
	for _, p := range ps {
		if p.Key == k {
			return p.Value, true
		}
	}
	return "", false
}

func SetParam(ps []Param, k, v string) []Param {
	out := append([]Param{}, ps...)
	for i, p := range out {
		if p.Key == k {
			out[i].Value = v
			return out
		}
	}
	return append(out, Param{k, v})
}

func RenderQuery(ps []Param) string {
	parts := make([]string, len(ps))
	for i, p := range ps {
		parts[i] = p.Key + "=" + p.Value
	}
	return strings.Join(parts, "&")
}
