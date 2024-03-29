package bb_pix_go

type Environment string
type ModalidadeAgente string

const (
	OauthHomologacaoUrl = "https://oauth.hm.bb.com.br/oauth/token"
	OauthProdUrl        = "https://oauth.bb.com.br"

	EnvHolomogacaoSemMTLS Environment = "https://api.hm.bb.com.br/pix/v2"
	EnvHomologacaoComMTLS Environment = "https://api-pix.hm.bb.com.br/pix/v2"
	EnvProducao           Environment = "https://api-pix.bb.com.br/pix/v2"

	// DevAppKey é o nome da query param que deve ser enviada junto a key do aplicativo
	DevAppKey = "gw-dev-app-key"

	// ModalidadeAgenteAGTEC significa Agente Estabelecimento Comercial
	ModalidadeAgenteAGTEC = "AGTEC"

	// ModalidadeAgenteAGTOT significa Agente Outra Espécie de Pessoa Jurídica ou Correspondente no País
	ModalidadeAgenteAGTOT = "AGTOT"

	// ModalidadeAgenteAGPSS significa Agente Facilitador de Serviço de Saque
	ModalidadeAgenteAGPSS = "AGPSS"

	// PixChaveHomologacao é a chave Pix para homologação que o BB disponibiliza
	PixChaveHomologacao = "hmtestes2@bb.com.br"

	// Scopes é o escopo de permissão que o aplicativo deve ter para acessar a API Pix
	Scopes = "pix.write pix.read cob.write cob.read cobv.write cobv.read"
)
