package middleware

import (
	"github.com/kataras/iris/v12/context"
	"log"
)

func Token(ctx context.Context) {
	log.Printf("中间件 Token :%s", ctx.RouteName())
	ctx.Next()
}
