package balance

func Balance(s string) bool {

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		switch s[i] {
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
		default:
			continue
		}
		stack = stack[:len(stack)-1]
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
