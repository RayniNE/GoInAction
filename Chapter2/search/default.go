package search

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}
