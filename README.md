# System-Delivery
Web app to track delivey


Aplicacao simple de deliverym que vai permitir visaluzar em tempo real o entregador.
Gerenciando multiplo entregadores de forma simulanea com cada rotda de forma facil identificação.


Similuação com posições do usuáio final. 


A posições e dados serao armazenado nos Elastic Search. 


Para evitar a perda de dados, nao será utilizado o REST, para isso nós vamos usar o Apache Kafka.


Websocket para conexao do back com front, é um forma de fazer conexao TCP com meu browser.
É um canal de comunicação direto.



Tools:
- Golang (Simulador)
- Backend: Nest.js & Mongo
- Frontend: React
- Kafka
- Kafka Connect 
    Vai pegar os dados do Kafka e enviara para o ElasticSearch
- Elastic Search
- Kibana 
- Docker e Kubernetes
- Istio, Kiali, Prometheus e Grafana

Topic:
    - É um stream que atua como banco de dados
    - Tem diversas partições = Topic
        - Cada partição tem um numero
    - 

Kafka Cluster:
    - Tem varias maquinas rodando
    - Conjunto de broker ( broker sao maquinas)
    - Cada partição do Topic está distribuida em diferentes brokers
    - Zookeper (É um serviço de server discovery, fica monitorando se a maquina está no ar faz o balanceamento)
    - 




