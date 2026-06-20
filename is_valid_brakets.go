package main

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.

// Open brackets must be closed in the correct order.

// Every close bracket has a corresponding open bracket of the same type.

// "[]({})" - valid

// "[]({)}" - invalid
// "[]({}" - invalid

// "{[()]}" valid

func isValidString(s string) bool{
	var stack []rune 
	brackets := map[rune]rune{
		'}' : '{',
		']' : '[',
		')' : '(',
	}
	
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			last_elem := stack[len(stack) -1]
			stack = stack[:len(stack)-1]
			if brackets[char] != last_elem {
				return false
			}
		}
	}
	return len(stack) < 1
}