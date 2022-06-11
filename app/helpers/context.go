package helpers

import (
	"context"
	"log"
)

func SetUserContext(ctx context.Context, userId uint) context.Context {
	return context.WithValue(ctx, "user_id", userId)
}

func CallUserContext(ctx context.Context) {
	userId := ctx.Value("user_id")

	log.Printf("User Id %v", userId)
}
