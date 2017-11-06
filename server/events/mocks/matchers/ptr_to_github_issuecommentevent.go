package matchers

import (
	"reflect"

	github "github.com/google/go-github/github"
	"github.com/petergtz/pegomock"
)

func AnyPtrToGithubIssuecommentevent() *github.IssueCommentEvent {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*github.IssueCommentEvent))(nil)).Elem()))
	var nullValue *github.IssueCommentEvent
	return nullValue
}

func EqPtrToGithubIssuecommentevent(value *github.IssueCommentEvent) *github.IssueCommentEvent {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *github.IssueCommentEvent
	return nullValue
}