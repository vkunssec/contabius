---
title: "Assinatura de Rotas"
type: "blank"
date: 2024-04-23
---


### Criar Conta
#### `/account/create`, método `POST`
Método de criação de uma conta

Corpo da Requisição:
```json
{
    "account": "account name",
    "color": "#cc0000"
}
```

Retorno esperado:
```json
{
    "id": "66286938afa1be771b54814e",
    "account": "account name",
    "color": "#cc0000",
    "created_at": "2024-04-23T23:06:48.0387898-03:00",
    "updated_at": "2024-04-23T23:06:48.0387898-03:00"
}
```

### Retornar Contas
#### `/account?ids=`, método `GET`
Rota de retorno de uma ou mais contas

Caso não seja informado nenhum ID, serão retornados todos os valores cadastrados.
Query Params:
```
ids[] (opcional): IDS das contas salvas

```

Retorno esperado:
```json
[
    {
        "id": "6621c29060c1b212a97c99ef",
        "account": "conta",
        "color": "#121212",
        "created_at": "2024-04-19T01:02:08.611Z",
        "updated_at": "2024-04-19T01:02:08.611Z"
    },
    {
        "id": "6621c32d60c1b212a97c99f1",
        "account": "account name",
        "color": "#820ad1",
        "created_at": "2024-04-19T01:04:45.876Z",
        "updated_at": "2024-04-19T01:04:45.876Z"
    }
]
```

### Atualizar Conta
#### `/account/update/:id`, método `PUT`
Método de atualização de uma conta

Path Variável:
```
id: ID da conta

```

Corpo da Requisição
```json
{
    "account": "renew name"
}
```

Retorno esperado:
```json
{
    "account": "renew name",
    "color": "#000000",
    "created_at": "2024-04-23T02:06:48.038Z",
    "updated_at": "2024-04-24T23:14:59.038Z"
}
```

### Excluir Conta
#### `/account/remove/:id`, método `DELETE`
Método para exclusão de uma conta

Path Variável:
```
id: ID da conta

```

Retorno esperado:
```json
{
    "success": true
}
```

