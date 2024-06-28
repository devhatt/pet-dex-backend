<p align="center">
  <a href="#1-catálogo-de-pets-personalizado">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://github.com/devhatt/pet-dex-backend/assets/76929097/e1521917-2177-44eb-a267-e46e89ca6459">
      <img src="https://github.com/devhatt/pet-dex-backend/assets/76929097/e1521917-2177-44eb-a267-e46e89ca6459" height="128">
    </picture>
    <h1 align="center">PetDex - Seu Catálogo de Pets Virtual</h1>
  </a>
</p>

[![Node.js](https://img.shields.io/badge/Node.js-18.x-green.svg)](https://nodejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.x-blue.svg)](https://www.typescriptlang.org/)
[![Prisma](https://img.shields.io/badge/Prisma-5.x-purple.svg)](https://www.prisma.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-blue.svg)](https://www.postgresql.org/)
[![Swagger](https://img.shields.io/badge/Swagger-OpenAPI-85EA2D.svg)](https://swagger.io/)

Bem-vindo ao PetDex, o aplicativo que transforma a experiência de ser tutor de pets em algo único e interativo. Com o PetDex, os tutores podem catalogar e compartilhar informações sobre seus pets, semelhante à famosa Pokedex, mas para animais de estimação.

## Funcionalidades Principais

### 1. **Catálogo de Pets Personalizado**

- Adicione informações sobre seus pets, incluindo nome, raça, idade e peculiaridades.
- Faça o upload de fotos adoráveis dos seus companheiros peludos.

### 2. **Exploração de Raças**

- Descubra novas raças de animais que você ainda não tem.
- Explore informações detalhadas sobre cada raça, como características físicas, temperamento e cuidados específicos.

## Como Contribuir

Se você é um entusiasta de pets, desenvolvedor em ascensão ou simplesmente quer fazer parte da comunidade PetDex, aqui estão algumas maneiras de contribuir:

1. **Desenvolvimento:**
   - Faça um fork do repositório e trabalhe em novas funcionalidades.
   - Resolva problemas existentes ou proponha melhorias.
2. **Documentação:**
   - Aprimore a documentação existente ou crie tutoriais para ajudar outros desenvolvedores.
3. **Testes:**
   - Ajude a garantir a estabilidade do aplicativo testando as novas funcionalidades e relatar problemas.

## Executando o projeto

### Manualmente

Requisitos:

- Go `1.21.4`

Todos os comandos devem ser executados na raiz do projeto. Não esqueça de adaptar as variáveis de ambiente e configurar as conexões com os serviços que a aplicação depende.

```bash
make run

# ou

go run ./api/main.go
```

### Com docker compose

Requisitos:

- Go `1.21.4`
- Docker

O projeto possui um arquivo `docker-compose.yml` que irá subir containers com todas as dependências do projeto e executará o programa com live reload ativado. Com um simples comando você vai ter um ambiente de desenvolvimento completamente configurado onde só vai se preocupar em codificar e salvar o código.

Executar o projeto com Docker Compose proporciona muitas vantagens, facilitando a colaboratividade e também executando o programa em um ambiente o mais parecido possível com o ambiente de produção.

Antes de iniciar, certifique-se de ter o [Docker](https://docs.docker.com/get-docker/) instalado e configurado corretamente em sua máquina.

Na raiz do projeto, copie o `.env.example` e nomeie o novo arquivo com `.env`:

```bash
cp .env.example .env
```

Por fim, execute o projeto com:

```bash
make dev

# ou

docker compose --profile development --env-file .env up # use -d para executar os containers em background
```

_Subir todos os containers pode demorar um tempo dependendo do seu setup ou internet._

## Contato

Se precisar de ajuda, tiver sugestões ou quiser se envolver mais profundamente com a comunidade PetDex, entre em contato conosco:

- Discord: [https://discord.gg/9f5BZ7yD](https://discord.gg/9f5BZ7yD)
- Twitter: [Devhat (@DevHatt) / X (twitter.com)](https://twitter.com/DevHatt)

Junte-se a nós nesta jornada emocionante de tornar o PetDex a melhor experiência para tutores de pets em todo o mundo!
