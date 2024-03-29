package bb_pix_go

type BB struct {
	Environment Environment

	// Token básico para acesso do OAuth
	Basic string

	// É a credencial necessária para acionar as APIs do Banco do Brasil, independente
	// de estar em produção ou não.
	DeveloperKey string

	// É o identificador público e único no OAuth do Banco do Brasil.
	ClientID string

	// É conhecido apenas para sua aplicação e o servidor de autorização.
	ClientSecret string

	// Configuração do certificado mTLS (caso seja usado)
	Certificado mTLS
}

// New é a função responsável por instanciar toda a API e permitir o uso
// é recomendado que use em forma de singleton no seu programa
func New(options ...func(*BB)) *BB {
	bb := &BB{}

	for _, o := range options {
		o(bb)
	}

	return bb
}

// WithEnvironment é responsável por passar o ambiente
// que você deseja executar os pedidos
func WithEnvironment(env Environment) func(*BB) {
	return func(bb *BB) {
		bb.Environment = env
	}
}

// WithDeveloperKey é responsável por definir a chave
// do desenvolvedor do aplicativo
func WithDeveloperKey(devKey string) func(*BB) {
	return func(bb *BB) {
		bb.DeveloperKey = devKey
	}
}

// WithClientID é responsável por definir o Client ID usado
func WithClientID(clientId string) func(*BB) {
	return func(bb *BB) {
		bb.ClientID = clientId
	}
}

// WithClientSecret é responsável por definir o Client Secret
func WithClientSecret(clientSecret string) func(*BB) {
	return func(bb *BB) {
		bb.ClientSecret = clientSecret
	}
}

// WithBasic é responsável por definir o token básico do Oauth
func WithBasic(basic string) func(*BB) {
	return func(bb *BB) {
		bb.Basic = basic
	}
}
