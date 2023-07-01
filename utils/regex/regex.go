package regex

import "regexp"

func ExtractCustomID(customID string, regex *regexp.Regexp,
	expectedGroupSize int) ([]string, bool) {
	if regex.MatchString(customID) {
		groups := regex.FindStringSubmatch(customID)
		if len(groups) == expectedGroupSize {
			return groups, true
		}
	}

	return nil, false
}

func IsBelongTo(customID string, regexes ...*regexp.Regexp) bool {
	var isBelongTo bool
	for _, regex := range regexes {
		isBelongTo = isBelongTo || regex.MatchString(customID)
	}
	return isBelongTo
}
