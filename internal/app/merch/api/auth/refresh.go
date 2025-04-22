package auth

import (
	"context"
	"fmt"

	"merch-store/internal/app/mappers"
	"merch-store/internal/domain"
	desc "merch-store/pkg/merch"

	"github.com/opentracing/opentracing-go"
)

func (i *Implementation) Refresh(ctx context.Context, in *desc.RefreshRequest) (*desc.RefreshResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.auth.Refresh")
	defer span.Finish()

	if err := in.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %w", domain.ErrInvalidArgument, err)
	}

	tokens, err := i.service.Refresh(ctx, mappers.ProtoToTokens(in.Tokens))
	if err != nil {
		return nil, err
	}

	return &desc.RefreshResponse{
		Tokens: mappers.TokensToProto(tokens),
	}, nil
}
