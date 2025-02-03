package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type contextKey int

const (
	subjectKey contextKey = iota

	HeaderUserID        = "X-Social-User-Id"
	HeaderAuthUserID    = "X-Social-Auth-User-Id"
	HeaderAuthorization = "Authorization"
)

type Subject struct {
	ID         string
	AuthUserID string
}

func SubjectFromContext(ctx context.Context) *Subject {
	if subject, ok := ctx.Value(subjectKey).(*Subject); ok {
		return subject
	}
	return nil
}

func ContextWithSubject(ctx context.Context, subject *Subject) context.Context {
	return context.WithValue(ctx, subjectKey, subject)
}

func ParseSubjectFromMetadata(ctx context.Context) (*Subject, error) {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("metadata not found")
		return nil, nil
	}
	if len(metadata.Get(HeaderUserID)) > 0 || len(metadata.Get(HeaderAuthUserID)) > 0 {
		return parseSubjectFromHeaders(metadata)
	}
	return nil, nil
}

func parseSubjectFromHeaders(metadata metadata.MD) (*Subject, error) {
	subject := new(Subject)
	if value := metadata.Get(HeaderUserID); len(value) > 0 {
		subject.ID = value[0]
	}
	if value := metadata.Get(HeaderAuthUserID); len(value) > 0 {
		subject.AuthUserID = value[0]
	}

	return subject, nil
}
