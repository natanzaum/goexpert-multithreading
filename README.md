## goexpert-multithreading  
###  Desafio 2 da Pós Go Expert - Trabalhando com Multithreading  

#### Como rodar o projeto:    
- Acesse a pasta do projeto  
- Digite o comando `go run main.go -cep=NUMERO_DO_CEP`

Exemplo :
```
go run main.go -cep=33200000
```

Exemplo de retorno BrasilCep:  
```
2026/07/25 20:22:36 Reposta recebida da BrasilCep!
{"cep":"33200000","state":"MG","city":"Vespasiano","neighborhood":"","street":"","service":"widenet"}
```
  
Exemplo de retorno ViaCep:  
```
2026/07/25 20:20:28 Reposta recebida da ViaCep!
{"cep":"31565-400","logradouro":"Rua Martinica","complemento":"","unidade":"","bairro":"Santa Branca","localidade":"Belo Horizonte","uf":"MG","estado":"Minas Gerais","regiao":"Sudeste","ibge":"3106200","gia":"","ddd":"31","siafi":"4123"}
```