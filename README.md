# Symphony

**Symphony** é uma rede social de músicas desenvolvida como parte de um projeto de laboratório de banco de dados. O objetivo deste projeto é proporcionar uma plataforma para usuários interagirem, compartilharem e descobrirem músicas, utilizando um banco de dados híbrido com tecnologias modernas para garantir escalabilidade e eficiência.

## Tecnologias Utilizadas

- **Neo4j**: Banco de dados gráfico para modelar e explorar relacionamentos entre músicas, artistas e usuários.
- **MongoDB**: Banco de dados NoSQL para armazenar informações não relacionais, como dados de usuários e metadados de músicas.
- **PostgreSQL**: Banco de dados relacional para armazenar dados estruturados e garantir integridade transacional.

## Funcionalidades

- **Perfis de Usuário**: Criação e personalização de perfis, onde usuários podem adicionar suas músicas favoritas e acompanhar suas playlists.
- **Descoberta de Música**: Sistema de recomendação baseado nos gostos e nas interações dos usuários.
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
   git clone https://github.com/felipetressmann/symphony-api.git
   cd symphony
   docker-compose up --build
   ```

## Autores
- Thiago Duvanel Ferreira
- Filipe Tressmann Velozo
- Guilherme Wallace Ferreira