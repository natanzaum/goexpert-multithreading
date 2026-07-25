## goexpert-multithreading  
###  Desafio 2 da Pós Go Expert - Trabalhando com Multithreading  

#### Como rodar o projeto:    
- Acesse a pasta do projeto  
- Digite o comando `go run main.go -cep=NUMERO DO CEP`

Exemplo :
```
go run main.go -cep=33200000
```

Exemplo de retorno BrasilCep:  
```
2026/07/25 20:18:50 Reposta recebida da BrasilCep!  
{"cep":"33203580","state":"MG","city":"Vespasiano","neighborhood":"Nova Pampulha 3ª e 4ª Seção","street":"Rua Trinta e Cinco","service":"open-cep"}
```
  
Exemplo de retorno ViaCep:  
```
2026/07/25 20:20:28 Reposta recebida da ViaCep!
{"cep":"33203-580","logradouro":"Rua Trinta e Cinco","complemento":"","unidade":"","bairro":"Nova Pampulha 3ª e 4ª Seção","localidade":"Vespasiano","uf":"MG","estado":"Minas Gerais","regiao":"Sudeste","ibge":"3171204","gia":"","ddd":"31","siafi":"5425"}
```