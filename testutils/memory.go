package testutils

import (
	"context"

	"github.com/athenianco/metadata-retrieval/github/graphql"

	"gopkg.in/src-d/go-log.v1"
)

// Memory implements the storer interface
type Memory struct {
	Organization     *graphql.Organization
	Repository       *graphql.RepositoryFields
	Topics           []string
	Users            []*graphql.UserExtended
	Issues           []*graphql.Issue
	IssueComments    []*graphql.IssueComment
	PRs              []*graphql.PullRequest
	PRComments       []*graphql.IssueComment
	PRReviews        []*graphql.PullRequestReview
	PRReviewComments []*graphql.PullRequestReviewComment
}

// SaveOrganization stores an organization in memory,
// it also initializes the list of users
func (s *Memory) SaveOrganization(ctx context.Context, organization *graphql.Organization) error {
	log.Infof("organization data fetched for %s\n", organization.Login)
	s.Organization = organization
	// Initialize users to 0 for each repo
	s.Users = make([]*graphql.UserExtended, 0)
	return nil
}

// SaveUser appends a user to the user list in memory
func (s *Memory) SaveUser(ctx context.Context, orgID int, orgLogin string, user *graphql.UserExtended) error {
	log.Infof("user data fetched for %s\n", user.Login)
	s.Users = append(s.Users, user)
	return nil
}

// SaveRepository stores a repository and its topics in memory and
// initializes PRs and PR comments
func (s *Memory) SaveRepository(ctx context.Context, repository *graphql.RepositoryFields, topics []string) error {
	log.Infof("repository data fetched for %s/%s\n", repository.Owner.Login, repository.Name)
	s.Repository = repository
	s.Topics = topics
	// Initialize prs and comments to 0 for each repo
	s.Issues = make([]*graphql.Issue, 0)
	s.IssueComments = make([]*graphql.IssueComment, 0)
	s.PRs = make([]*graphql.PullRequest, 0)
	s.PRComments = make([]*graphql.IssueComment, 0)
	s.PRReviews = make([]*graphql.PullRequestReview, 0)
	s.PRReviewComments = make([]*graphql.PullRequestReviewComment, 0)
	return nil
}

// SaveIssue appends an issue to the issue list in memory
func (s *Memory) SaveIssue(ctx context.Context, repositoryOwner, repositoryName string, issue *graphql.Issue, assignees []string, labels []string) error {
	log.Infof("issue data fetched for #%v %s\n", issue.Number, issue.Title)
	s.Issues = append(s.Issues, issue)
	return nil
}

// SaveIssueComment appends an issue comment to the issue comments list in memory
func (s *Memory) SaveIssueComment(ctx context.Context, repositoryOwner, repositoryName string, issueNumber int, comment *graphql.IssueComment) error {
	log.Infof("\tissue comment data fetched by %s at %v: %q\n", comment.Author.Login, comment.CreatedAt, trim(comment.Body))
	s.IssueComments = append(s.IssueComments, comment)
	return nil
}

// SavePullRequest appends an PR to the PR list in memory, also initializes the map for the review commnets for that particular PR
func (s *Memory) SavePullRequest(ctx context.Context, repositoryOwner, repositoryName string, pr *graphql.PullRequest, assignees []string, labels []string) error {
	log.Infof("PR data fetched for #%v %s\n", pr.Number, pr.Title)
	s.PRs = append(s.PRs, pr)
	return nil
}

// SavePullRequestComment appends an PR comment to the PR comment list in memory
func (s *Memory) SavePullRequestComment(ctx context.Context, repositoryOwner, repositoryName string, pullRequestNumber int, comment *graphql.IssueComment) error {
	log.Infof("\tpr comment data fetched by %s at %v: %q\n", comment.Author.Login, comment.CreatedAt, trim(comment.Body))
	s.PRComments = append(s.PRComments, comment)
	return nil
}

// SavePullRequestReview appends a PR review to the PR review list in memory
func (s *Memory) SavePullRequestReview(ctx context.Context, repositoryOwner, repositoryName string, pullRequestNumber int, review *graphql.PullRequestReview) error {
	log.Infof("\tPR Review data fetched by %s at %v: %q\n", review.Author.Login, review.SubmittedAt, trim(review.Body))
	s.PRReviews = append(s.PRReviews, review)
	return nil
}

// SavePullRequestReviewComment appends a PR review comment to the PR review comments list in memory
func (s *Memory) SavePullRequestReviewComment(ctx context.Context, repositoryOwner, repositoryName string, pullRequestNumber int, pullRequestReviewID int, comment *graphql.PullRequestReviewComment) error {
	log.Infof("\t\tPR review comment data fetched by %s at %v: %q\n", comment.Author.Login, comment.CreatedAt, trim(comment.Body))
	s.PRReviewComments = append(s.PRReviewComments, comment)
	return nil
}

// Begin is a noop method at the moment
func (s *Memory) Begin() error {
	return nil
}

// Commit is a noop method at the moment
func (s *Memory) Commit() error {
	return nil
}

// Rollback is a noop method at the moment
func (s *Memory) Rollback() error {
	return nil
}

// Version is a noop method at the moment
func (s *Memory) Version(v int) {
}

// SetActiveVersion is a noop method at the moment
func (s *Memory) SetActiveVersion(ctx context.Context, v int) error {
	return nil
}

// Cleanup is a noop method at the moment
func (s *Memory) Cleanup(ctx context.Context, currentVersion int) error {
	return nil
}

func trim(s string) string {
	if len(s) > 40 {
		return s[0:39] + "..."
	}
	return s
}
