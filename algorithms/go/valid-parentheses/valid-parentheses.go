package isValid

func isValid(s string) bool {
   charactersMap := map[byte]byte{
        '(' : ')',
        '{' : '}',
        '[' : ']',
    }

    size := len(s)

    if size % 2 > 0{
        return false
    }

    stackcap :=  len(s)/2
    stack := make([]byte, 0, stackcap)
    for _, k := range s{
        if closingCharacter, ok := charactersMap[byte(k)]; ok{
            stack = append(stack, closingCharacter)
        }else if len(stack) > 0{
            last := stack[len(stack)-1]
            if last != byte(k) {
                return false
            }
            stack = stack[:len(stack)-1]
        }else{
            return false
        }

    }

    return len(stack) == 0
}
