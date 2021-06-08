package middleware

import (
	"context"
	"fmt"
	"github.com/auto-check/common-module/jwt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

func Authenticator(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return nil, err
	}

	id, err := jwt.ParseStudentIDFromToken(token)
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return nil, err
	}

	newCtx := context.WithValue(ctx, "student_id", id)

	return newCtx, nil
}
