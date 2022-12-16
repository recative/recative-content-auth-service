package admin_token_model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/recative/recative-service-sdk/pkg/db/db_err"
	"gorm.io/gorm"
	"time"
)

type TokenModel interface {
	ReadTokenInfoByRaw(tokenRaw string) (*Token, error)
	UpdateTokenInfoByRaw(tokenRaw string, param TokenParam) error
	DeleteTokenByRaw(tokenRaw string) error
	CreateToken(param TokenParam) (token Token, err error)
	ReadAllTokens() (token []Token, err error)
	ReadSelectTokens(tokenRaws []string) (token []Token, err error)
	IsTokenExist(token string) bool
	GenerateSudoToken(sudoToken string) (token Token)
}

type Token struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	Id uuid.UUID `gorm:"primaryKey;not null;uniqueIndex"`

	TokenParam
}

func (t *Token) IsValid() bool {
	return t.TokenParam.IsValid && t.ExpiredAt.After(time.Now())
}

type TokenType string

const (
	TokenTypeAdmin TokenType = "admin"
	TokenTypeSudo  TokenType = "sudo"
)

type AdminPermission = string

const (
	AdminPermissionSudo  AdminPermission = "sudo"
	AdminPermissionRead  AdminPermission = "read"
	AdminPermissionWrite AdminPermission = "write"
)

type TokenParam struct {
	Type            TokenType      `gorm:"default:'admin'"`
	Raw             string         `gorm:"unique;index"`
	AdminPermission pq.StringArray `gorm:"type:varchar[]"`
	Comment         string
	ExpiredAt       time.Time
	IsValid         bool
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	t.Id = uuid.New()
	if t.Raw == "" {
		t.Raw = uuid.New().String()
	}
	return nil
}

func (m *model) ReadTokenInfoByRaw(tokenRaw string) (*Token, error) {
	var res Token
	err := m.db.Where("raw = ?", tokenRaw).First(&res).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return &res, nil
}

func (m *model) UpdateTokenInfoByRaw(tokenRaw string, param TokenParam) error {
	err := m.db.Where("raw = ?", tokenRaw).Updates(Token{TokenParam: param}).Error
	if err != nil {
		return db_err.Wrap(err)
	}
	return nil
}

func (m *model) DeleteTokenByRaw(tokenRaw string) error {
	err := m.db.Where("raw = ?", tokenRaw).Delete(&Token{}).Error
	if err != nil {
		return db_err.Wrap(err)
	}
	return nil
}

func (m *model) CreateToken(param TokenParam) (token Token, err error) {
	token = Token{
		TokenParam: param,
	}
	err = m.db.Create(&token).Error
	if err != nil {
		return Token{}, db_err.Wrap(err)
	}
	return token, nil
}

func (m *model) ReadAllTokens() (token []Token, err error) {
	err = m.db.Find(&token).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return token, nil
}

func (m *model) ReadSelectTokens(tokenRaws []string) (token []Token, err error) {
	err = m.db.Where("raw IN ?", tokenRaws).Find(&token).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return token, nil
}

func (m *model) IsTokenExist(token string) bool {
	var count int64
	m.db.Model(&Token{}).Where("raw = ?", token).Count(&count)
	return count > 0
}

func (m *model) GenerateSudoToken(sudoToken string) (token Token) {
	return Token{
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Id:        uuid.UUID{},
		TokenParam: TokenParam{
			Type:            TokenTypeSudo,
			Raw:             sudoToken,
			AdminPermission: []string{AdminPermissionSudo},
			Comment:         "This is a sudo token",
			ExpiredAt:       time.Time{},
		},
	}
}
