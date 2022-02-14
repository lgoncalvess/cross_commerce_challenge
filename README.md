# Cross Commerce Store Desafio

Simples Aplicação ETL

- Todo o código fonte está contido no arquivo main.go
- Para executar basta acessar a API que está hospedada no Heroku no seguinte endpoint: https://crossommerceapi.herokuapp.com/ 
- Utilizei concorrência, através das Go Routines, para lidar com as requisições http (Que poderiam ser muitas) na etapa de Extract, assim possibilitando uma maior velocidade.
- Para ordenação utilizei o algoritmo Merge Sort
- Para criar a Api HTTP utilizei a biblioteca Echo
