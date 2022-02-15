# Cross Commerce Store Desafio

Simples Aplicação ETL
- Para executar basta acessar a API que está hospedada no Heroku no seguinte endpoint: https://crossommerceapi.herokuapp.com/ 
![Captura de tela de 2022-02-14 12-00-55](https://user-images.githubusercontent.com/41243909/153888636-8c70f834-962c-430d-be23-8f7d6724f3a2.png)
- Utilizei concorrência, através das Go Routines, para lidar com as requisições http (Que poderiam ser muitas) na etapa de Extract, assim possibilitando uma maior velocidade.
- Para ordenação utilizei o algoritmo Merge Sort
- Para criar a Api HTTP utilizei a biblioteca Echo

Caso queira executar em sua máquina local
- Instale o Go https://go.dev/doc/install
- Entre no diretório do projeto
- compile o projeto
- execute o arquivo compilado
-  Basta fazer um request na porta 3000 que o programa irá retornar um Slice ordenado
