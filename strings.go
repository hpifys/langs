package langs

import "strings"

// StringContainsParts 检测字符串str中是否包含切片parts中的字符串，caseSensitive标示是否大小写敏感
func StringContainsParts(str string, parts []string, caseSensitive bool) bool {
    target := str
    if !caseSensitive {
        target = strings.ToUpper(target)
    }

    for _, part := range parts {
        _part := part
        if !caseSensitive {
            _part = strings.ToUpper(_part)
        }
        if strings.Contains(target, _part) {
            return true
        }
    }
    return false
}
