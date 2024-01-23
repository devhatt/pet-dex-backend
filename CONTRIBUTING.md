# Diretrizes de Contribuição para o Repositório da PetDex

Bem-vindo à PetDex! Agradecemos pelo seu interesse em contribuir para este projeto open source. Suas contribuições são fundamentais para o sucesso e a melhoria contínua deste projeto. Antes de começar, por favor, leia e siga estas diretrizes para garantir um processo de contribuição harmonioso e eficaz.

## Como Contribuir

1. **Forquilhando o Repositório:** Faça um fork deste repositório para sua própria conta GitHub clicando no botão "Fork".
2. **Clonando o Repositório:** Clone o fork do repositório para a sua máquina local:

   ```bash
   git clone <https://github.com/devhatt/pet-dex-backend>
   cd pet-dex-backend
   ```

3. **Pegue uma Issue:** 

   Navegue pelas "Issues" e de preferência por uma Issue marcada como "Good First Issue". As "Issues" dessa forma são especialmente indicadas para novos colaboradores e são pontos de partida acessíveis no projeto. Mas, caso não haja nenhuma assim ou você se considere apto a pegar outra sem essa marcação, você possui total liberdade para pegá-la.

   Depois de escolher uma task, clique nela para obter mais detalhes.

   Lembre-se:
   Tarefas com fotos ao lado direito já foram selecionadas por outros colaboradores.
   Geralmente, você terá uma semana para concluir a tarefa, mas isso pode variar dependendo das políticas do projeto.


4. **Crie uma Branch:** Crie uma branch para trabalhar nas suas alterações.


5. **Faça Alterações:** Faça as alterações desejadas no código, documentação, ou outros recursos.
6. **Testes:** Certifique-se de que todas as mudanças são testadas e não introduzem erros.
7. **Commits Significativos:** Faça commits significativos e com mensagens claras. Utilizando comando abaixo e seguindo as instruções o commit ficara no padrão utilizado no projeto.

   ```bash
      git commit
   ```

   1. **Atualize a Documentação:** Se necessário, atualize a documentação relevante para refletir suas mudanças.
   2. **Envie as Alterações:** Envie suas alterações para o seu fork:

      ```bash
      git push origin nome-da-sua-branch

      ```

   3. **Criação de Pull Request (PR):** Abra um Pull Request pelo o seu fork para o repositorio da PetDex, descrevendo suas alterações e fornecendo contexto sobre o que foi feito.
   4. **Revisão de Código:** A equipe de mantenedores do projeto irá revisar o seu PR. Esteja disposto a fazer ajustes se necessário.
   5. **Merge e Fechamento:** Após a revisão bem-sucedida, suas alterações serão mescladas à branch principal. Seu PR será fechado.

## Diretrizes de Contribuição

- **Documentação:** Sempre atualize a documentação para refletir mudanças significativas.
- **Testes:** Certifique-se de que suas alterações não quebram testes existentes. Se necessário, adicione novos testes.
- **Tamanho das Pull Requests:** PRs menores são mais fáceis de revisar e mesclar. Tente manter o escopo de suas contribuições relativamente pequeno.
- **Mantenha a Cortesia:** Seja cortês e respeitoso ao discutir e revisar o trabalho de outros contribuidores.

## Reconhecimento

Agradecemos por ajudar a melhorar a PetDex! Sua dedicação à qualidade e inovação é fundamental para o sucesso contínuo deste projeto.

Se você tiver alguma dúvida ou precisar de ajuda em qualquer etapa do processo de contribuição, sinta-se à vontade para criar um problema (issue) ou entrar em contato com a equipe de mantenedores.[Discord](discord.gg/3gsMAEumEd)


## Banco de dados
Usamos MariaDB como Banco de Dados Relacional da aplicação. 
#### Migration
Uma migration é um script que é usado para alterar o esquema de um banco de dados, como a adição de novas tabelas, colunas ou índices.
Para criar uma migration, você deve criar um novo arquivo com a extensão .sql. O conteúdo do arquivo deve conter as alterações que você deseja fazer no esquema do banco de dados.
Os arquivos .sql deverão ser criados com o comando abaixo:
```bash
   make create-migrations title=titulo_da_migration
```
Ao executar o comando, serão criados dois arquivos na pasta /migrations com sufixos `.up` e `.down`. Os arquivos `.up` representam as alterações desejadas que serão aplicados no banco de dados, enquantos os arquivos `.down`, representa a ação de rollback referente ao que foi executado no arquivos `.up` de mesma versão.

Obs.: crie os script SQL de forma idempotente e caso o seu script tenha vários comando ou consultas, considere colocar isso em uma transação.