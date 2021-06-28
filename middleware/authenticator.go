package middleware

import (
	"context"
	"fmt"
	"github.com/auto-check/common-module/jwt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"runtime/debug"
)

func Authenticator(ctx context.Context) (context.Context, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	log.Println(headers)

	token := headers.Get("authorization")[1]

	id, err := jwt.ParseStudentIDFromToken(token)
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return nil, err
	}

	newCtx := context.WithValue(ctx, "student_id", id)

	return newCtx, nil
}
