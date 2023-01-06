# CEP Retriever

Serviço que realiza busca por CEP em mais de um serviço disponível na internet, devolvendo o resultado do serviço que retornar mais rápido.

## Uso

```
curl --location --request GET 'http://localhost:8080/retrieve/04055-000'
```

Exemplo de retorno: 

```
{
    "cep": "04055-000",
    "address": "Rua Fagundes Dias",
    "district": "Saúde",
    "city": "São Paulo",
    "state": "SP"
}
```

**Swagger:** http://localhost:8080/docs/index.html

Para subir o serviço:

```
docker compose up -d
```