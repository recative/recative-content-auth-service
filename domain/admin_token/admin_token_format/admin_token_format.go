package admin_token_format

import (
	"github.com/recative/recative-backend/domain/admin_token/admin_token_model"
	"github.com/recative/recative-backend/spec"
	"github.com/recative/recative-service-sdk/util/ref"
	"time"
)

func TokenToResponse(token *admin_token_model.Token) spec.TokenResponse {
	return spec.TokenResponse{
		AdminPermission: token.AdminPermission,
		Comment:         token.Comment,
		ExpiredAt:       ref.T(token.ExpiredAt.Format(time.RFC3339)),
		IsValid:         ref.T(token.IsValid()),
		Token:           token.Raw,
		Type:            spec.TokenResponseType(token.Type),
	}
}

func TokensToResponses(tokens []admin_token_model.Token) spec.TokensResponse {
	var res spec.TokensResponse
	for _, token := range tokens {
		res = append(res, TokenToResponse(&token))
	}
	return res
}

func TokenRequestToTokenParam(req spec.TokenRequest) admin_token_model.TokenParam {
	var token = ""
	if req.Token != nil {
		token = *req.Token
	}

	var comment = ""
	if req.Comment != nil {
		comment = *req.Comment
	}

	var expiredAt = time.Time{}
	if req.ExpiredAt != nil {
		expiredAt, _ = time.Parse(time.RFC3339, *req.ExpiredAt)
	}

	return admin_token_model.TokenParam{
		Raw:             token,
		AdminPermission: AdminPermissionArrayToStringArray(req.AdminPermission),
		Comment:         comment,
		ExpiredAt:       expiredAt,
	}
}

func AdminPermissionArrayToStringArray(array []spec.AdminPermission) []string {
	var res []string
	for _, v := range array {
		res = append(res, string(v))
	}
	return res
}
