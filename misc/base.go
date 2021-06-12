package misc

import "unicode"

func VerifyPassword(s string) (length, number, lower, upper, special bool) {
    letters := 0
    for _, c := range s {
        switch {
        case unicode.IsNumber(c):
            number = true
        case unicode.IsUpper(c):
            upper = true
        case unicode.IsLower(c):
            lower = true
        case unicode.IsPunct(c) || unicode.IsSymbol(c):
            special = true
        default:
            //return false, false, false, false
        }
        letters++
    }
    length = letters >= 8 || letters <= 21
    return
}


