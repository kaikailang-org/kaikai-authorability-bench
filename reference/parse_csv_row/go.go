func ParseCsvRow(s string) ([]string, bool) {
	fields := []string{}
	i, n := 0, len(s)
	for {
		var f []byte
		if i < n && s[i] == '"' {
			i++
			closed := false
			for i < n {
				if s[i] == '"' {
					if i+1 < n && s[i+1] == '"' {
						f = append(f, '"')
						i += 2
					} else {
						i++
						closed = true
						break
					}
				} else {
					f = append(f, s[i])
					i++
				}
			}
			if !closed {
				return nil, false
			}
			if i < n && s[i] != ',' {
				return nil, false
			}
		} else {
			for i < n && s[i] != ',' {
				if s[i] == '"' {
					return nil, false
				}
				f = append(f, s[i])
				i++
			}
		}
		fields = append(fields, string(f))
		if i == n {
			break
		}
		i++
	}
	return fields, true
}
