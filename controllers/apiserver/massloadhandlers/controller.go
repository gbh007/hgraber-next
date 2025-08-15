package massloadhandlers

import (
	"context"
	"log/slog"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"go.opentelemetry.io/otel/trace"
)

type MassloadController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	debug bool,
	ac *apiservercore.Controller,
) *MassloadController {
	c := &MassloadController{
		logger:  logger,
		tracer:  tracer,
		debug:   debug,
		apiCore: ac,
	}

	return c
}

// FIXME: реализовать
func (c *MassloadController) APIMassloadInfoAttributeCreatePost(ctx context.Context, req *serverapi.APIMassloadInfoAttributeCreatePostReq) (serverapi.APIMassloadInfoAttributeCreatePostRes, error) {
	return &serverapi.APIMassloadInfoAttributeCreatePostInternalServerError{
		InnerCode: apiservercore.MassloadUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

// FIXME: реализовать
func (c *MassloadController) APIMassloadInfoAttributeDeletePost(ctx context.Context, req *serverapi.APIMassloadInfoAttributeDeletePostReq) (serverapi.APIMassloadInfoAttributeDeletePostRes, error) {
	return &serverapi.APIMassloadInfoAttributeDeletePostInternalServerError{
		InnerCode: apiservercore.MassloadUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

// FIXME: реализовать
func (c *MassloadController) APIMassloadInfoCreatePost(ctx context.Context, req *serverapi.APIMassloadInfoCreatePostReq) (serverapi.APIMassloadInfoCreatePostRes, error) {
	return &serverapi.APIMassloadInfoCreatePostInternalServerError{
		InnerCode: apiservercore.MassloadUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

// FIXME: реализовать
func (c *MassloadController) APIMassloadInfoDeletePost(ctx context.Context, req *serverapi.APIMassloadInfoDeletePostReq) (serverapi.APIMassloadInfoDeletePostRes, error) {
	return &serverapi.APIMassloadInfoDeletePostInternalServerError{
		InnerCode: apiservercore.MassloadUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

// FIXME: реализовать
func (c *MassloadController) APIMassloadInfoExternalLinkCreatePost(ctx context.Context, req *serverapi.APIMassloadInfoExternalLinkCreatePostReq) (serverapi.APIMassloadInfoExternalLinkCreatePostRes, error) {
	return &serverapi.APIMassloadInfoExternalLinkCreatePostInternalServerError{
		InnerCode: apiservercore.MassloadUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

// FIXME: реализовать
func (c *MassloadController) APIMassloadInfoExternalLinkDeletePost(ctx context.Context, req *serverapi.APIMassloadInfoExternalLinkDeletePostReq) (serverapi.APIMassloadInfoExternalLinkDeletePostRes, error) {
	return &serverapi.APIMassloadInfoExternalLinkDeletePostInternalServerError{
		InnerCode: apiservercore.MassloadUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

// FIXME: реализовать
func (c *MassloadController) APIMassloadInfoGetPost(ctx context.Context, req *serverapi.APIMassloadInfoGetPostReq) (serverapi.APIMassloadInfoGetPostRes, error) {
	return &serverapi.APIMassloadInfoGetPostInternalServerError{
		InnerCode: apiservercore.MassloadUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

// FIXME: реализовать
func (c *MassloadController) APIMassloadInfoListGet(ctx context.Context) (serverapi.APIMassloadInfoListGetRes, error) {
	return &serverapi.APIMassloadInfoListGetInternalServerError{
		InnerCode: apiservercore.MassloadUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

// FIXME: реализовать
func (c *MassloadController) APIMassloadInfoUpdatePost(ctx context.Context, req *serverapi.APIMassloadInfoUpdatePostReq) (serverapi.APIMassloadInfoUpdatePostRes, error) {
	return &serverapi.APIMassloadInfoUpdatePostInternalServerError{
		InnerCode: apiservercore.MassloadUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}
