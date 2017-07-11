package path

import (
    "regexp"
)

var (
    unCommonStringRegexp = regexp.MustCompile("[^A-Za-z0-9\\.\\-\u4e00-\u9fa5]+")
)

func HasUnCommonString(src string) bool {
    return unCommonStringRegexp.MatchString(src)
}

func ReplaceUnCommonString(src string) string {
    return unCommonStringRegexp.ReplaceAllString(src, "_")
}

