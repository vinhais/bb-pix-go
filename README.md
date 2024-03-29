# SDK Go API Pix v2 Banco do Brasil

🚧 Atualmente o projeto está em desenvolvimento e não é recomendado para uso em produção.

## Instalação

Para instalar o SDK, execute o seguinte comando na raiz do seu projeto:
```bash
go get github.com/vinhais/bb-pix-go
```

## Acesso a API

Para acessar a API Pix do Banco do Brasil, é necessário ter uma conta PJ e ter o aplicativo criado na plataforma de desenvolvedores do Banco do Brasil.

Para mais informações, acesse o [Portal de Desenvolvedores do Banco do Brasil](https://developers.bb.com.br/).

## Exemplos de uso

```go
package main

import (
    bb "github.com/vinhais/bb-pix-go"
)

func main() {
    api := bb.New(
        bb.WithEnvironment(bb.EnvHolomogacaoSemMTLS),
        bb.WithClientID("XXXXX"),
        bb.WithClientSecret("XXXXX"),
        bb.WithDeveloperKey("XXXXXX"),
        bb.WithBasic("Basic xxxxx"),
    )
}
```

## A fazer

- [ ] Finalizar o suporte a Cobrança Imediata (cob)
- [ ] Implementar o suporte a Cobrança com Vencimento (cobv)
- [ ] Implementar o suporte a mTLS e certificados (necessário para produção)
- [ ] Implementar o suporte a Webhooks
- [ ] Implementar todos os endpoints da API Pix
- [ ] Implementar testes unitários
- [ ] Adicionar exemplos de uso

Se você tiver alguma sugestão ou feedback, adoraríamos ouvir de você! 🚀✨
