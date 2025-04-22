package mappers

import (
	"merch-store/internal/domain"
	desc "merch-store/pkg/merch"
)

func TokensToProto(tokens domain.Tokens) *desc.TokensPair {
	return &desc.TokensPair{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}

func ProtoToTokens(pb *desc.TokensPair) domain.Tokens {
	return domain.Tokens{
		AccessToken:  pb.AccessToken,
		RefreshToken: pb.RefreshToken,
	}
}
