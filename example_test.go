package bb_pix_go

import (
	"testing"
)

func TestBBInstance(t *testing.T) {
	bb := New(
		WithEnvironment(EnvHolomogacaoSemMTLS),
		WithClientID("XXXXX"),
		WithClientSecret("XXXXX"),
		WithDeveloperKey("XXXXXX"),
		WithBasic("Basic xxxxx"),
	)

	// gerando um txid com 26 caracteres
	txid := RandomString(26)

	// instanciando a estrutura do pix
	pix := Pix{
		TxID:  txid,
		Chave: "hmtestes2@bb.com.br",
		Valor: PixValor{
			Original: "10.00",
		},
		Calendario: &PixCalendario{},
	}

	// testando a criação de cobrança
	t.Run("deve gerar um código de qr code", func(t *testing.T) {
		i, err := bb.CriarCob(pix)
		if err != nil {
			t.Error(err)
		}

		t.Log(i)
	})
}
