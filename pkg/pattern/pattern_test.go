package pattern

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/isbang/simple/pkg/testutil"
)

func TestIsOneLevelWildCard(t *testing.T) {
	tests := []struct {
		s      string
		output bool
	}{{
		s:      ":foo",
		output: true,
	}, {
		s:      "foo",
		output: false,
	}}

	for _, test := range tests {
		if b := isOneLevelWildCard(test.s); b != test.output {
			testutil.TestFailed(t, b, test.output)
		}
	}
}

func TestIsMultiLevelWildCard(t *testing.T) {
	tests := []struct {
		s      string
		output bool
	}{{
		s:      "*foo",
		output: true,
	}, {
		s:      "foo",
		output: false,
	}}

	for _, test := range tests {
		if b := isMultiLevelWildCard(test.s); b != test.output {
			testutil.TestFailed(t, test.output, b)
		}
	}
}

func TestMatchPattern(t *testing.T) {
	tests := []struct {
		p, v   Pattern
		output bool
	}{{
		p:      Pattern("foo/bar"),
		v:      Pattern("foo/bar"),
		output: true,
	}, {
		p:      Pattern("foo/:key"),
		v:      Pattern("foo/bar"),
		output: true,
	}, {
		p:      Pattern(":key/bar"),
		v:      Pattern("foo/bar"),
		output: true,
	}, {
		p:      Pattern(":key/bar"),
		v:      Pattern("foo/:bar"),
		output: true,
	}, {
		p:      Pattern(":key/bar/"),
		v:      Pattern("foo/:bar/"),
		output: true,
	}, {
		p:      Pattern(":key/bar/:some"),
		v:      Pattern("foo/:bar/"),
		output: true,
	}, {
		p:      Pattern("*key"),
		v:      Pattern("foo/:bar"),
		output: true,
	}, {
		p:      Pattern("*key"),
		v:      Pattern("foo/:bar/"),
		output: true,
	}, {
		p:      Pattern("foo/*key"),
		v:      Pattern("foo/:bar"),
		output: true,
	}, {
		p:      Pattern("foo/*key"),
		v:      Pattern("foo/:bar/"),
		output: true,
	}, {
		p:      Pattern("foo/bar"),
		v:      Pattern("foo/par"),
		output: false,
	}, {
		p:      Pattern("foo/bar"),
		v:      Pattern("foo/bar/"),
		output: false,
	}, {
		p:      Pattern(":key/bar"),
		v:      Pattern("foo/:bar/some"),
		output: false,
	}, {
		p:      Pattern("key/:bar"),
		v:      Pattern("foo/:bar/some"),
		output: false,
	}, {
		p:      Pattern("foo/bar/*key"),
		v:      Pattern("foo/:bar"),
		output: false,
	}}

	for _, test := range tests {
		if b := test.p.MatchPattern(test.v); b != test.output {
			testutil.TestFailed(t, test.output, b)
		}
	}
}

func TestMatchString(t *testing.T) {
	tests := []struct {
		p      Pattern
		s      string
		output bool
	}{{
		p:      Pattern("foo/bar"),
		s:      "/foo/bar",
		output: true,
	}, {
		p:      Pattern(":key/bar"),
		s:      "/foo/bar",
		output: true,
	}, {
		p:      Pattern("foo/:key"),
		s:      "/foo/bar",
		output: true,
	}, {
		p:      Pattern(":key/:key"),
		s:      "/foo/bar",
		output: true,
	}, {
		p:      Pattern("foo/*some"),
		s:      "/foo/bar",
		output: true,
	}, {
		p:      Pattern("*some"),
		s:      "/foo/bar",
		output: true,
	}, {
		p:      Pattern("foo/bar"),
		s:      "/foo/bar/",
		output: false,
	}, {
		p:      Pattern("foo/bar"),
		s:      "/foo/bar/som",
		output: false,
	}, {
		p:      Pattern("foo/*key"),
		s:      "/foo",
		output: false,
	}, {
		p:      Pattern(":key/bar"),
		s:      "/as/df",
		output: false,
	}, {
		p:      Pattern("foo/:bar"),
		s:      "/foo/bar/som",
		output: false,
	}}

	for _, test := range tests {
		if b := test.p.MatchString(test.s); b != test.output {
			testutil.TestFailed(t, test.output, b)
		}
	}
}

func TestGetMatched(t *testing.T) {
	tests := []struct {
		p      Pattern
		s      string
		output url.Values
	}{{
		p:      Pattern("*key"),
		s:      "/foo/bar",
		output: url.Values{"key": []string{"foo/bar"}},
	}, {
		p:      Pattern("*key"),
		s:      "/foo/bar/",
		output: url.Values{"key": []string{"foo/bar/"}},
	}, {
		p:      Pattern("foo/*key"),
		s:      "/foo/bar",
		output: url.Values{"key": []string{"bar"}},
	}, {
		p:      Pattern("foo/*key"),
		s:      "/foo/bar/",
		output: url.Values{"key": []string{"bar/"}},
	}, {
		p:      Pattern(":key/bar"),
		s:      "/foo/bar",
		output: url.Values{"key": []string{"foo"}},
	}, {
		p:      Pattern(":key/bar/"),
		s:      "/foo/bar/",
		output: url.Values{"key": []string{"foo"}},
	}, {
		p:      Pattern("foo/:key"),
		s:      "/foo/bar",
		output: url.Values{"key": []string{"bar"}},
	}, {
		p:      Pattern("foo/:key/"),
		s:      "/foo/bar/",
		output: url.Values{"key": []string{"bar"}},
	}}

	for _, test := range tests {
		matched := test.p.GetMatched(test.s)
		if !reflect.DeepEqual(matched, test.output) {
			testutil.TestFailed(t, test.output, matched)
		}
	}
}

func TestIsInvalid(t *testing.T) {
	tests := []struct {
		p       Pattern
		invalid bool
	}{{
		p:       Pattern("foo/bar"),
		invalid: false,
	}, {
		p:       Pattern("foo:/bar"),
		invalid: false,
	}, {
		p:       Pattern(":foo/bar"),
		invalid: false,
	}, {
		p:       Pattern("foo/:bar"),
		invalid: false,
	}, {
		p:       Pattern("foo:/:bar:"),
		invalid: false,
	}, {
		p:       Pattern("foo:/bar/"),
		invalid: false,
	}, {
		p:       Pattern("/wrong"),
		invalid: true,
	}, {
		p:       Pattern(""),
		invalid: true,
	}, {
		p:       Pattern("wow//awesome"),
		invalid: true,
	}, {
		p:       Pattern(":"),
		invalid: true,
	}, {
		p:       Pattern("*"),
		invalid: true,
	}, {
		p:       Pattern("*foo/bar"),
		invalid: true,
	}}

	for _, test := range tests {
		if err := test.p.IsInvalid(); (err != nil) != test.invalid {
			testutil.TestFailed(t, test.invalid, err)
		}
	}
}
