package pattern

import (
	"net/url"
	"strings"
)

var (
	MultiLevelWildCard byte = '*'
	OneLevelWildCard   byte = ':'
)

// Pattern used to match multiple URLs in a single string
//
// Pattern must not starts with '/'
//
// Pattern allows suffix "/". Use strings.TrimSuffix if you need.
type Pattern string

func (p Pattern) MatchPattern(v Pattern) bool {
	plist := strings.Split(string(p), "/")
	vlist := strings.Split(string(v), "/")

	sameLength := len(plist) == len(vlist)

	for i := 0; i < minInt(len(plist), len(vlist)); i++ {
		if isMultiLevelWildCard(plist[i]) || isMultiLevelWildCard(vlist[i]) {
			return true
		}

		if isOneLevelWildCard(plist[i]) || isOneLevelWildCard(vlist[i]) {
			continue
		}

		if plist[i] == vlist[i] {
			continue
		}

		return false
	}

	return sameLength
}

func (p Pattern) MatchString(s string) bool {
	plist := strings.Split(string(p), "/")
	slist := strings.Split(strings.TrimPrefix(s, "/"), "/")

	sameLength := len(plist) == len(slist)

	for i := 0; i < minInt(len(plist), len(slist)); i++ {
		if isMultiLevelWildCard(plist[i]) {
			return true
		}

		if isOneLevelWildCard(plist[i]) {
			continue
		}

		if plist[i] == slist[i] {
			continue
		}

		return false
	}

	return sameLength
}

// GetMatched find match value from s and return as Value
//
// Pattern.MatchString(s) must true
func (p Pattern) GetMatched(s string) (v url.Values) {
	v = make(url.Values)

	plist := strings.Split(string(p), "/")
	slist := strings.Split(strings.TrimPrefix(s, "/"), "/")

	for i := 0; i < minInt(len(plist), len(slist)); i++ {
		if isMultiLevelWildCard(plist[i]) {
			v.Add(plist[i][1:], strings.Join(slist[i:], "/"))
			return
		}

		if isOneLevelWildCard(plist[i]) {
			v.Add(plist[i][1:], slist[i])
			continue
		}
	}
	return
}

func (p Pattern) IsInvalid() error {
	if strings.HasPrefix(string(p), "/") {
		return ErrInvalidPattern
	}

	plist := strings.Split(strings.TrimSuffix(string(p), "/"), "/")
	for i, s := range plist {
		if len(s) == 0 {
			return ErrInvalidPattern
		}

		if isOneLevelWildCard(s) && len(s) == 1 {
			return ErrOneLevelWileCard
		}

		if isMultiLevelWildCard(s) && (len(s) == 1 || i != len(plist)) {
			return ErrMultiLevelWildCard
		}
	}
	return nil
}

func isMultiLevelWildCard(s string) bool { return len(s) > 0 && s[0] == MultiLevelWildCard }

func isOneLevelWildCard(s string) bool { return len(s) > 0 && s[0] == OneLevelWildCard }

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
