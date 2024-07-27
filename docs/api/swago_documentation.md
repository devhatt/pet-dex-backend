
# Swag 

Pacote responsável pela geração automática de documentação API RESTful com Swagger 2.0 para Go. 

## Sumário

1. [Dependências](#dependências)

2. [Nova documentação](#executando-o-Swago)

3. [Acesso](#acessando-documentação)

### Dependências

Para habilitar a geração automática de documentação via Swago, é necessário a sua instalação. Execute o comando a seguir para disponibilizar os recursos dessa feramenta pré configurados via Makefile.

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
Em caso de dúvidas na instalação e na sintaxe da documentação, consulte o repositório oficial da documentaçãopode ser acessado
* [swag](https://github.com/swaggo/swag)

Também é necessário ter o GNU Make instalado para executar os comandos das migrations:

* [Make](https://www.gnu.org/software/make/)

### Executando o swago

Para gerar a documentação da API, depois documentado o endpoint, execute o comando a seguir:

```bash
make swag
```

Esse comando irá atualizar os arquivos swagger disponíveis no diretório `docs/api`


### Acessando Documentação

A documentação estará disponívei para acesso através da rota `api/swagger/`
