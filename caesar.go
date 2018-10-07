package main

func caesar(in input) output {

	return output{}
}

func shift(in string, n int) string {
	const power = 26
	out := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		ch := in[i]
		switch {
		case byte('a') <= ch && ch <= byte('z'):
			out[i] = byte((int(ch)-'a'+n)%power + 'a')
		case byte('A') <= ch && ch <= byte('Z'):
			out[i] = byte((int(ch)-'A'+n)%power + 'A')
		default:
			out[i] = ch
		}
	}
	return string(out)
}