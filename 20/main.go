package main

type Stack struct {
	items []rune
}

func (stk *Stack) peek() rune {
	return stk.items[len(stk.items)-1]
}

func (stk *Stack) push(val rune) {
	stk.items = append(stk.items, val)
}

func (stk *Stack) pop() {
	stk.items = stk.items[0 : len(stk.items)-1]
}

func (stk *Stack) isEmpty() bool {
	return len(stk.items) == 0
}

func isValid(s string) bool {
	characters := &Stack{}
	openCharacters := map[rune]rune{}
	openCharacters['('] = ')'
	openCharacters['['] = ']'
	openCharacters['{'] = '}'

	for _, character := range s {
		_, ok := openCharacters[character]
		if ok {
			characters.push(character)
		} else {
			if len(characters.items) == 0 {
				return false
			}
			if openCharacters[characters.peek()] == character {
				characters.pop()
			} else {
				return false
			}
		}
	}

	return characters.isEmpty()
}

func main() {
	println(isValid("{{}}"))
}
