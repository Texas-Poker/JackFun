package server_http

import (
	"context"
	"server/pb/pb_http"
)

func (this *ComponentHttp) TestRPC(ctx context.Context, req *pb_http.ReqEntry) (*pb_http.RespEntry, error) {
	return nil, nil
}