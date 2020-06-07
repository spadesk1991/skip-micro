package wrappers

import (
	"context"

	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro/errors"

	"github.com/micro/go-micro/server"
)

// logWrapper is a handler wrapper
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Infof("[LogWrapper] server request: %v", req.Endpoint())
		err := fn(ctx, req, rsp)
		if err != nil {
			log.Errorf("[LogWrapper]%s", err)
			return errors.BadRequest("111", `%s`, err.Error())
		}
		return err
	}
}
