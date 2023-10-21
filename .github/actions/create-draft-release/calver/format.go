package main

import (
	"regexp"
	"strconv"
	"time"
)

/**
* GetNextVersion gets next version according to the tag of the format.
* @ref https://github.com/Connehito/gdp/blob/5a5e9ccec14b9d7a0462cfc6d85667add4ba5a46/format.go#L30-L47
**/
func GetNextVersion(tag string) (string, error) {

	// date version(e.g. 20180525.1 or release_20180525.1)
	const layout = "20060102"
	today := time.Now().Format(layout)

	dateRe := regexp.MustCompile(`(.*)(\d{8})\.(.+)`)
	if m := dateRe.FindStringSubmatch(tag); m != nil {
		if m[2] == today {
			minor, err := strconv.Atoi(m[3])
			if err != nil {
				return "", err
			}

			next := strconv.Itoa(minor + 1)
			return m[1] + today + "." + next, nil
		}
		return m[1] + today + "." + "1", nil
	}
	return today + ".1", nil
}
