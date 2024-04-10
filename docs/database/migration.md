
# Migration

## Sumário

1. [Dependências](#dependências)

2. [Novas Migrations](#criando-novas-migrations)

3. [Executando Migrations](#executando-as-migrations)

### Dependências

As migrations utilizam a CLI do Golang Migrate, sendo assim é necessário instalar em seu ambiente
de desenvolvimento:

* [Golang Migrate Instalação](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

Também é necessário ter o GNU Make instalado para executar os comandos das migrations:

* [Make](https://www.gnu.org/software/make/)

### Criando novas migrations

Para criar novos arquivos de migrations, execute o comando abaixo:

```bash
make create-migrations 
```

Ele criará na pasta `migrations` dois arquivos, sendo um para **UP** e outro para **DOWN**.
No primeiro você digitará o código `sql` para executar a mudança necessária no banco (incluir uma nova tabela, 
uma coluna nova e etc), já para o DOWN você escreverá o necessário para reverter essas mudanças caso seja preciso.

### Executando as migrations

Para executar as migrations previamente criadas é necessário rodar um dos dois comandos abaixo:

```bash
make run-migrations-up
```
```bash
make run-migrations-down
```

Como já deve ter percebido o primeiro comando executa as mudanças desejadas e o segundo reverte essas mudanças.