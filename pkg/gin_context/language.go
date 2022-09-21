package gin_context

import (
	"golang.org/x/text/language"
)

func MatchLanguage(header string, tags []string) []string {
	var t []language.Tag
	for _, ts := range tags {
		t = append(t, language.MustParse(ts))
	}
	m := language.NewMatcher(t)
	acceptLanguage, _, err := language.ParseAcceptLanguage(header)
	if err != nil {
		return []string{}
	}
	var rs []string
	set := map[string]struct{}{}
	for _, al := range acceptLanguage {
		_, index, _ := m.Match(al)
		if _, ok := set[tags[index]]; !ok {
			set[tags[index]] = struct{}{}
			rs = append(rs, tags[index])
		}
	}
	return rs
}
