package salamoonder

import (
	"fmt"
	"regexp"
)

var chromeVersionRe = regexp.MustCompile(`Chrome/(\d+)`)

func ExtractSecChUa(userAgent string) string {
	v := "122"
	if match := chromeVersionRe.FindStringSubmatch(userAgent); match != nil {
		v = match[1]
	}
	return fmt.Sprintf(`"Google Chrome";v="%s", "Chromium";v="%s", "Not(A:Brand";v="24"`, v, v)
}
