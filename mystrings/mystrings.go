package mystrings

func Reverse(s string) string{      //we need to capitalize the first letter of the function name if we want to export it.
    result := ""
    for _, v := range s {
        result = string(v) + result
    }
    return result;
}
