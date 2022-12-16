package mock_data

import (
	"github.com/recative/recative-backend/domain/admin_token/admin_token_model"
	"github.com/recative/recative-service-sdk/pkg/db/mock"
	"gorm.io/gorm"
	"time"
)

func InitAdminToken(db *gorm.DB) {
	err := mock.Mock(db, false,
		testAdminToken1,
		testAdminToken2,
	)
	if err != nil {
		panic(err)
	}
}

var testAdminToken1 = admin_token_model.Token{
	TokenParam: admin_token_model.TokenParam{
		Type:            admin_token_model.TokenTypeAdmin,
		Raw:             "1919810",
		AdminPermission: []string{"read", "write"},
		Comment:         "test1",
		ExpiredAt:       time.Now().Add(time.Hour * 24 * 30),
		IsValid:         true,
	},
}

var testAdminToken2 = admin_token_model.Token{
	TokenParam: admin_token_model.TokenParam{
		Type:            admin_token_model.TokenTypeAdmin,
		Raw:             "114514",
		AdminPermission: []string{"sudo"},
		Comment:         "test2",
		ExpiredAt:       time.Now().Add(time.Hour * 24 * 30),
		IsValid:         false,
	},
}
