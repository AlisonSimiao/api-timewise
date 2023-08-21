# Timewise API

O Timewise API é uma aplicação em Go (Golang) que utiliza o framework Gin Tonic e a biblioteca GORM para fornecer uma API para gerenciar informações temporais.

    A API  permite que os usuários cadastrem e gerenciem projetos, registrando o que foi feito em cada projeto e o tempo decorrido durante sua execução. Os usuários podem criar projetos, adicionar atividades realizadas e obter informações sobre o progresso de cada projeto.

## Funcionalidades Principais

### Cadastro de Projetos

Os usuários podem criar novos projetos fornecendo informações como nome do projeto, descrição e data de início.

### Adicionar Atividades

 Para cada projeto, os usuários podem adicionar atividades realizadas, registrando o que foi feito em cada etapa do projeto. Cada atividade terá uma descrição e a data/hora de inicio e termino em que foi realizada.

### Tempo Decorrido

 A API calcula automaticamente o tempo decorrido desde o início do projeto até a data/hora de cada atividade registrada, fornecendo uma visão clara do progresso do projeto.

### Listar Projetos

 Os usuários podem visualizar uma lista de todos os projetos cadastrados, incluindo detalhes como o nome do projeto, data de início e a quantidade de atividades registradas.

### Detalhes do Projeto

Ao selecionar um projeto específico, os usuários podem ver todas as atividades associadas a ele, incluindo suas descrições e datas.

## Tecnologias utilizadas

### Linguagem

- <span style="font-size: 20px; font-weight: bold;  color: #fff;text-shadow: 2px 2px #100;">Go (Golang)</span>: [https://golang.org/](https://golang.org/)


Golang, ou Go, é uma linguagem de programação de código aberto desenvolvida pelo Google. Conhecida por sua simplicidade, eficiência e gerenciamento de concorrência. É amplamente usada para desenvolvimento web, serviços e aplicativos escaláveis.

biblioteca | docs | descrição
--- | --- | ---
<span style="font-size: 20px; font-weight: bold;  color: #4CAF50;text-shadow: 2px 2px #100;">Gin Tonic<span> | [link docs](https://github.com/gin-gonic/gin) | Gin Tonic é um framework web em Go (Golang) que oferece um conjunto de funcionalidades para criação rápida de APIs e serviços web com alta performance e baixa latência. Ele utiliza uma sintaxe simples e possui uma estrutura flexível, tornando-o uma ótima escolha para desenvolvimento web em Go.
<span style="font-size: 20px; font-weight: bold; color: #0e83cd;font-family: sans-serif; text-transform:uppercase; text-shadow: 2px 2px #100;">GORM</span> | [link docs](https://gorm.io/docs/) | GORM é uma biblioteca ORM (Object-Relational Mapping) em Go (Golang), que oferece uma interface fácil de usar para interagir com bancos de dados relacionais. Com o GORM, é possível realizar operações de criação, leitura, atualização e exclusão (CRUD) em registros do banco de dados, além de oferecer suporte para relacionamentos e migração de esquema.
<span style="font-size: 20px; font-weight: bold;  color: #F58300; text-shadow: 2px 2px #100; font-family: sans-serif;"> cloudnary </span> | [link docs](https://cloudinary.com/documentation/go_integration )| Cloudinary é uma plataforma de gerenciamento e entrega de mídia baseada em nuvem, que oferece serviços para armazenar, otimizar e entregar imagens e vídeos em aplicativos web e móveis. Ela fornece APIs poderosas para manipulação de mídia e integração fácil com várias linguagens de programação, incluindo Go, tornando-a uma escolha popular para gerenciamento de ativos de mídia em escala.

### Database

- <span style="font-size: 20px; font-weight: bold;  color: #336791; text-shadow: 2px 2px #100; font-family: sans-serif;"> Postgres </span> - [https://www.postgresql.org/]( https://www.postgresql.org/)

O PostgreSQL é um poderoso sistema de gerenciamento de banco de dados relacional de código aberto e gratuito. Ele é conhecido por sua confiabilidade, recursos avançados, suporte a ACID (Atomicidade, Consistência, Isolamento e Durabilidade) e sua capacidade de lidar com grandes volumes de dados.

## Instalação

1. `Clone este repositório` para o diretório de sua preferência:

```bash
git clone git@github.com:AlisonSimiao/api-timewise.git
```

2. Entre na pasta criada e faça `download das dependencias`

```bash
    cd /api-timewise

    go mod tidy
```

atente-se as variaveis de ambiente disponobilizadas no `.env.example` e preencha-as em um arquivo `.env` antes do proximo passo

3. `rode` ou `build` o projeto

```golang
    go run main.go
```

ou

```golang
    go build

    ./time-wise
```

# Apendices

## Type_photo

index | description
--- | ---
1 | profile
2 | cover project
3 | default

## Status

index | description
--- | ---
1 | Em Andamento:
2 | Concluído
3 | Pendente
4 | Atrasado
5 | Cancelada
6 | Bloqueado
7 | Reaberto
8 | Em Revisão
9 | Aguardando Aprovação
10| Em Teste