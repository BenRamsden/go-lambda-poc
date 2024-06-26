package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
)

// CreatePanic is the resolver for the createPanic field.
func (r *mutationResolver) CreatePanic(ctx context.Context, message string) (bool, error) {
	// We want to trigger sentry
	panic(message)
}

// GetPanic is the resolver for the getPanic field.
func (r *queryResolver) GetPanic(ctx context.Context, message string) (bool, error) {
	// We want to trigger sentry
	panic(message)
}
