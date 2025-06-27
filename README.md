# Symphony

**Symphony** é uma rede social de músicas desenvolvida como parte de um projeto de laboratório de banco de dados. O objetivo deste projeto é proporcionar uma plataforma para usuários interagirem, compartilharem e descobrirem músicas, utilizando um banco de dados híbrido com tecnologias modernas para garantir escalabilidade e eficiência.

## Tecnologias Utilizadas

- **Neo4j**: Banco de dados gráfico para modelar e explorar relacionamentos entre músicas, artistas e usuários.
- **MongoDB**: Banco de dados NoSQL para armazenar informações não relacionais, como dados de usuários e metadados de músicas.
- **PostgreSQL**: Banco de dados relacional para armazenar dados estruturados e garantir integridade transacional.

## Funcionalidades

- **Perfis de Usuário**: Criação e personalização de perfis, onde usuários podem adicionar suas músicas favoritas e acompanhar suas playlists.
- **Relacionamentos Sociais**: Usuários podem seguir outros usuários, compartilhar músicas e playlists.
- **Listas de Reproduções**: Criação de playlists personalizadas e compartilháveis.
- **Interações**: Curtir, comentar e compartilhar músicas e playlists.

## Como Rodar o Projeto

### Pré-requisitos

1. [Neo4j](https://neo4j.com/download/)
2. [MongoDB](https://www.mongodb.com/try/download/community)
3. [PostgreSQL](https://www.postgresql.org/download/)

### Instalação

1. Clone este repositório:
   ```bash
   git clone git@github.com:filipetressmann/symphony-api.git
   cd symphony
   ```
2. Crie um arquivo `.env` com os conteúdos presentes no arquivo `.env.example`
3. Execute o comando abaixo para subir a aplicação:
   ```bash
   docker-compose up --build
   ```

### Como popular a aplicação com dados aleatórios?

Para popular a aplicação com dados aleatórios você pode executar o script `populateDB.py` presente na raiz do projeto:
   ```bash
   python3 populateDB.py
   ```

Além disso, existem os scripts `integration_tests.sh` e `mongo_integration_test.sh`. Esses scripts também podem ser usados para popular o banco mas servem mais como um sanity check ao desenvolver.

### Como executar requests?

Para interagir com a API, suba a aplicação com `docker-compose` e vá até `http://localhost:8080/swagger/index.html` no seu navegador para abrir a UI do swagger.

## Autores
- Thiago Duvanel Ferreira
- Filipe Tressmann Velozo
- Guilherme Wallace Ferreira