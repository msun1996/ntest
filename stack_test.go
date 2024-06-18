package ntest

/**
1. 有效的括号 "()[]{}". 可以{[ ]}但是不能([ )]
*/

func isValid(s string) bool {

	if len(s)%2 == 1 {
		return false
	}
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if m[s[i]] > 0 {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || m[stack[len(stack)-1]] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
