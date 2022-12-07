package util

import (
	"unicode"

	"regexp"
)

func IsChinese(str string) bool {
    var count int
    for _, v := range str {
        if unicode.Is(unicode.Han, v) {
            count++
            break
        }
    }
    return count > 0
}

func IsChineseChar(str string) bool {
    for _, r := range str {
        if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
            return true
        }
    }
    return false
}