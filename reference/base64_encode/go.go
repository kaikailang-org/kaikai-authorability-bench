const b64abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Base64Encode(bytes []int64) (string, bool) {
	for _, b := range bytes {
		if b < 0 || b > 255 {
			return "", false
		}
	}
	out := []byte{}
	i := 0
	for ; i+2 < len(bytes); i += 3 {
		v := bytes[i]<<16 | bytes[i+1]<<8 | bytes[i+2]
		out = append(out, b64abc[v>>18], b64abc[(v>>12)&63], b64abc[(v>>6)&63], b64abc[v&63])
	}
	switch len(bytes) - i {
	case 1:
		v := bytes[i] << 16
		out = append(out, b64abc[v>>18], b64abc[(v>>12)&63], '=', '=')
	case 2:
		v := bytes[i]<<16 | bytes[i+1]<<8
		out = append(out, b64abc[v>>18], b64abc[(v>>12)&63], b64abc[(v>>6)&63], '=')
	}
	return string(out), true
}
