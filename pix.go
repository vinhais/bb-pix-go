package bb_pix_go

import "encoding/json"

type Pix struct {
	TxID       string         `json:"txid"`
	Calendario *PixCalendario `json:"calendario,omitempty"`
	Devedor    *PixDevedor    `json:"devedor,omitempty"`
	Loc        *PixLoc        `json:"loc,omitempty"`
	Valor      PixValor       `json:"valor"`
	Chave      string         `json:"chave"`

	// Opcional, representa um texto, a ser apresentado ao
	// usuário pagador para que ele possa digital uma
	// informação correlata, em formato livre, a ser enviada ao
	// usuário recebedor. Máximo 140 caracteres.
	SolicitacaoPagador string `json:"solicitacaoPagador,omitempty"`

	// Cada informação adicional será apresentada ao usuário
	// pagador. Máximo 50 itens.
	InfoAdicionais []PixInfoAdicional `json:"infoAdicionais,omitempty"`
}

type PixCalendario struct {
	// Tempo de vida da cobrança, em segundos, a partir da
	// data de criação. Deve maior do que zero. Se não
	// enviado (em branco), assume 86.400 segundos, ou 24
	// horas (default). Máximo: 2.147.483.648 segundos.
	Expiracao string `json:"expiracao,omitempty"`
}

type PixDevedor struct {
	// Nome do devedor, máximo 200 caracteres. Se
	// preenchido, o CPF ou o CNPJ deve ser informado.
	Nome string `json:"nome,omitempty"`

	*PixDevedorPF
	*PixDevedorPJ
}

type PixDevedorPJ struct {
	// CNPJ do devedor. Não deve ser preenchido se o CPF
	// for informado. Se o nome for informado, o CPF ou o
	// CNPJ deve ser preenchido.
	CNPJ string `json:"cnpj,omitempty"`
}

type PixDevedorPF struct {
	// CPF do devedor. Não deve ser preenchido se o CNPJ
	// for informado. Se o nome for informado, o CPF ou o
	// CNPJ deve ser preenchido.
	CPF string `json:"cpf,omitempty"`
}

type PixLoc struct {
	// ID do location. Deve ser informado se o usuário
	// recebedor desejar utilizar um location previamente
	// reservado, do tipo 'cob'.
	ID string `json:"id,omitempty"`
}

type PixValor struct {
	// Valor original da cobrança. Deve ser informado, com
	// casas decimais, mesmo que seja 0. Máximo: 13 dígitos.
	Original json.Number `json:"original"` // o uso do json.Number foi necessário, pois float64 ele removeria o (.) e causaria problemas com a API

	// Determina se o valor final da cobrança pode ser alterado
	// pelo usuário pagador. Se omitido, assume valor 0 (false).
	ModalidadeAlteracao bool `json:"modalidadeAlteracao,omitempty"`

	// Estrutura opcional; se utilizada, a cobrança deixa de
	// considerada Pix comum e passa à categoria de Pix
	// Saque e Pix Troco.
	Retirada *PixValorRetirada `json:"retirada,omitempty"`
}

type PixValorRetirada struct {
	// Informações referentes ao saque. Não pode ser utilizada
	// juntamente com troco. Se utilizado, o PixValor.Original
	// deve ser 0, e o PixValor.ModalidadeAlteracao deve ser
	// falso.
	Saque bool `json:"saque,omitempty"`

	// Informações referentes ao troco. Não pode ser utilizada
	// juntamente com saque. Se utilizado, o PixValor.Original
	// deve ser maior que 0, e o PixValor.ModalidadeAlteracao
	// deveer falso
	Troco bool `json:"troco,omitempty"`

	// Valor do saque/troco a ser realizado.
	Valor json.Number `json:"valor,omitempty"`

	// Determina se o valor de saque/troco pode ser alterado
	// pelo usuário pagador. Se omitido, assume como
	// não pode alterar.
	ModalidadeAlteracao bool `json:"modalidadeAlteracao,omitempty"`

	// Indica a modalidade do agente por meio da qual se dá a
	// facilitação do serviço de saque.
	// Domínio:
	// AGTEC – Agente Estabelecimento Comercial
	// AGTOT – Agente Outra Espécie de Pessoa Jurídica ou Correspondente no País
	// AGPSS – Agente Facilitador de Serviço de Saque
	ModalidadeAgente string `json:"modalidadeAgente,omitempty"`

	// ISPB do Facilitador de Serviço de Saque
	PrestadorDoServicoDeSaque string `json:"prestadorDoServicoDeSaque,omitempty"`
}

type PixInfoAdicional struct {
	// Nome do campo. Máximo 50 caracteres.
	Nome string `json:"nome"`

	// Valor do campo. Máximo 200 caracteres.
	Valor string `json:"valor"`
}

type PixResposta struct{}
