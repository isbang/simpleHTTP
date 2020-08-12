package pattern

import "errors"

var (
	ErrInvalidPattern     = errors.New("pattern: invalid pattern string")
	ErrOneLevelWileCard   = errors.New("pattern: invalid one level wildcard")
	ErrMultiLevelWildCard = errors.New("pattern: invalid multi level wildcard")
)
