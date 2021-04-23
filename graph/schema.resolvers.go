package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/cadazab/gqlgenServer/graph/generated"
	"github.com/cadazab/gqlgenServer/graph/model"
)

func (r *mutationResolver) CreatePoll(ctx context.Context, question string, answerA string, answerB string) (*model.Poll, error) {
	votesA := 0
	votesB := 0
	poll := &model.Poll{
		Question: question,
		AnswerA:  &model.VotedAnswer{Answer: answerA, Votes: &votesA},
		AnswerB:  &model.VotedAnswer{Answer: answerB, Votes: &votesB},
	}
	r.polls = append(r.polls, poll)
	return poll, nil
}

func (r *queryResolver) Polls(ctx context.Context) ([]*model.Poll, error) {
	return r.polls, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
