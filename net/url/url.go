package url

import (
    "net/url"
    "path"
    "strings"
)

func ResolveReference(u string, ref string) (string, error) {
    refUrl, err := url.Parse(ref)
    if err != nil {
        return "", err
    }
    if refUrl.IsAbs() {
        return refUrl.String(), nil
    }
    uUrl, err := url.Parse(u)
    if err != nil {
        return "", err
    }
    return uUrl.ResolveReference(refUrl).String(), nil
}

func GetFileName(u string) string {
    _, s := path.Split(u)
    return strings.Split(s, "?")[0]
}

