<h1 align="center" style="font-weight: bold;">Decisões</h1>

## Pensamentos

Esse arquivo é referente aos pensamentos, ideias e decisões. Como cheguei a alguma decisão, o desenvolvimento de funcionalidades e etc. Digamos que é um segundo cerebro, quem sabe essa anotação vira um artigo no medium.


## Decisões Iniciais
Para inicio decidir a arquitertura é importante para termos organização e após pensar e ler alguns artigos decidi que a melhor escolha seria a junção de clean arquiterture e hexagonal 

https://dev.to/espigah/go-estrutura-de-projetos-1j0k<br>
https://johnfercher.medium.com/go-arquitetura-hexagonal-dbcd2e986b55<br>
https://dev.to/booscaaa/implementando-clean-architecture-com-golang-4n0a

## Nível 1
Um ponto relevante foi a validação, tive que pesquisar e o giroscópio pode retornar 0 que seria o norte além da latitude e longitude então a utilização do * no float64 para comparar se é nulo

## Nível 2
Comecei Fazendo os primeiros testes e tem um artigo que é ótimo para se basear, interessante adicionar o -cover para ser possível analisar a cobertura de testes, é interessante que no primeiro momento os testes sejam o fluxo básico (fluxo feliz) e os erros mais comnus, 100% não é o número mágico, não é garantia de nada. por isso que 60% já é ótimo para este MVP e 80% num sistema real.
A importância de testes unitários, por conta de um teste encontrei um erro que deixei, estava verificando longitude duas vezes ao invés de latitude.
Falando um pouquinho sobre mocks, eles são essênciais mas não podem ser tudo, devem ser utilizados com moderação para casos atipícos, horários e rotas diferente mas mockar todos os testes não é recomendável. Existe uma diferença entre mock e dados de testes. 

https://medium.com/@habbema/golang-testes-86da3e5e0602

Já estava convicto de fazer em com o postgres porém se a vaga é para dev cloud seria importante mostrar essas compentêcias então a melhor escolha era utilizar DynamoDB. Mas utilizar os produtos aws sempre causam temor de ter gastos exorbitantes então utilizarei o dynamo local neste primeiro momento, para isso o artigo abaixo mostra como começar.

https://medium.com/@snassr/dynamodb-with-go-golang-quickstart-e0e005b88e8f

## Nível 3
Referencias sobre o assunto de docker e golang, sempre ter um foco que sua imagem docker precisa ser a mais otimizada possível 
https://dev.to/rflpazini/go-docker-como-criar-as-melhores-imagens-docker-para-aplicacoes-golang-ikj

O .env não setava nos testes passei alguns bons minutos, como é test coloquei as variáveis no código e funcionou, como é test e são apenas vatiáveis fakes não tem grande impacto.

## Nível 4
Artigo utilizado para estudo
https://docs.aws.amazon.com/rekognition/latest/dg/faces-comparefaces.html

## Nível 5
Artigos utilizados para estudo
https://www.synadia.com/blog/building-a-job-queue-with-nats-io-and-go

https://dev.to/aleksk1ng/go-nats-grpc-and-postgresql-clean-architecture-microservice-with-monitoring-and-tracing-2kka

Poderia utilizar design baseado a eventos mais precisamente baseado em mensagens, porém eu não pensei nem conhecia mas sempre é bom conversar com alguém mais experiente, tirar ideias que você desconhece, quem sabe eu não refaço esse teste focando nessa arquitetura.

## Nível 6
Artigo baseado para swagger
https://dev.to/booscaaa/documentanto-uma-api-go-com-swagger-2k05

Ótimo vídeo sobre c4 model
https://www.youtube.com/watch?v=iWkXd0RJ2FA&t=3783s