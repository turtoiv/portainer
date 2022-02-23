package validator

import (
	"container/list"
)

var pairs = map[int32]int32{
	'n': 'n',
	'{': '}',
	'(': ')',
	'[': ']',
}

/* We use a stack to determine if the sequence is valid.
** If at the end of string parsing stack is not empty the string is not a valid representation
** If a closing tag is encountered and the previous element is not the matching tag the string is invalid
** and we exit the iteration
 */
func Validate(buffer string) bool {
	if len(buffer) == 0 {
		return false
	}

	var stack list.List
	prevValue := 'n'

	for _, value := range buffer {
		if pairs[prevValue] != value {
			//if we have a closing tab which does not match the previous tag
			//we return because at this point the string is not a valid representation
			if isClosingTag(value) {
				return false
			}
			stack.PushFront(value)
			prevValue = value
		} else if stack.Len() > 0 {
			e := stack.Front()
			stack.Remove(e)
			if stack.Len() > 0 {
				elem := stack.Front()
				prevValue = elem.Value.(int32)
			}
		}
	}

	return stack.Len() == 0

}

func isClosingTag(tag int32) bool {
	return tag == '}' || tag == ']' || tag == ')'
}
