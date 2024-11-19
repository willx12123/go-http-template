package validation

import (
	"github.com/dlclark/regexp2"
)

var validNameRegex = regexp2.MustCompile(`^(?![0-9-])[A-Za-z0-9\u4e00-\u9fa5-]+(?<!-)$`, regexp2.DefaultUnmarshalOptions)

// ValidNameChar 确保 name 不包含特殊字符
func ValidNameChar(name string) bool {
	if name == "" {
		return true
	}
	result, err := validNameRegex.MatchString(name)
	if err != nil {
		return false
	}
	return result
}
