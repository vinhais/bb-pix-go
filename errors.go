package bb_pix_go

import "errors"

var (
	ErrHttpRequestFoundation = errors.New("falha ao criar a requisição HTTP")
	ErrMTLSNotConfigured     = errors.New("mTLS não configurado")
	ErrEnvironmentNotSet     = errors.New("ambiente não configurado")
	ErrClientIDNotSet        = errors.New("client id não configurado")
	ErrClientSecretNotSet    = errors.New("client secret não configurado")
	ErrDeveloperKeyNotSet    = errors.New("developer key não configurado")
	ErrBasicNotSet           = errors.New("basic não configurado")
	ErrAccessTokenNotSet     = errors.New("access token não configurado")
	ErrAccessTokenExpired    = errors.New("access token expirado")
	ErrOauthFailed           = errors.New("falha ao obter o token de acesso")
)
