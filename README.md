## Escola DevOps - Trilha DevOps

### ~~3. Aprender sobre integração e entrega contínua~~

### ~~3.1. Integração contínua: testes automatizados e pipeline no Github Actions~~

### ~~3.2. Integração contínua: Pipeline Docker no Github Actions~~


## Escola DevOps - Trilha/formação [Integração Contínua e Entrega Contínua](https://cursos.alura.com.br/formacao-integracao-continua-entrega-continua)

### ~~1. Criando uma rotina de CI~~

### ~~1.1. Integração contínua: testes automatizados e pipeline no Github Actions~~

### ~~1.2. Integração contínua: Pipeline Docker no Github Actions~~


### 2. ~~Começando com entrega contínua~~

### 2.1 ~~Integração Contínua: Pipeline de entrega e implementação contínua na EC2~~

### 2.2 ~~Integração Contínua: Automatize o deploy no Amazon ECS~~


### 3. Garantindo a aplicação e Kubernetes

### 3.1 Integração Contínua: rollback e teste de carga

### 3.2 Integração Contínua: automatizando a entrega no Kubernetes

# Requisitos para execução local:

- docker;
- docker compose;
- linguagem go;

```bash
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg lsb-release
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo apt install golang-go

# para testar a aplicação
docker compose up -d
go run main.go
go test -v main_test.go
```

Passos na criação do cluster ECS (2.2):

1- criação de cluster no Amazon ECS;

2- criação de 'task definition' no Amazon ECS;

3- criação de serviço através da task ou do cluster no Amazon ECS (faz ligação entre a task e o cluster);

4- para criar CI/CD no github para o ECS, pode requerer criar usuário IAM caso exista apenas user root na conta;

4.1- caso falhe criação do serviço no ECS, verificar se foi atribuído o grupo de segurança da API e não do banco em RDS, e se a permissão do grupo de segurança não tem restrições na porta da aplicação. 

Eu erroneamente restringi acessos na porta da aplicação para apenas meu IP, só que agora a aplicação estaria containerizada e por trás de um balanceador de carga que precisa ter permissão para acessar a API.

Também verificar se a imagem Docker está ok, ou seja, se está executando sem erros.
