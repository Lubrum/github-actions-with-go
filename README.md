## Escola DevOps - Trilha DevOps

### ~~3. Aprender sobre integração e entrega contínua~~

### ~~3.1. Integração contínua: testes automatizados e pipeline no Github Actions~~

### ~~3.2. Integração contínua: Pipeline Docker no Github Actions~~


## Escola DevOps - Trilha/formação [Integração Contínua e Entrega Contínua](https://cursos.alura.com.br/formacao-integracao-continua-entrega-continua)

### ~~1. Criando uma rotina de CI~~

### ~~1.1. Integração contínua: testes automatizados e pipeline no Github Actions~~

### ~~1.2. Integração contínua: Pipeline Docker no Github Actions~~


### 2. ~~Começando com entrega contínua~~

### 2.1 Integração Contínua: Pipeline de entrega e implementação contínua na EC2

### 2.2 Integração Contínua: Automatize o deploy no Amazon ECS


### 3. Garantindo a aplicação e Kubernetes

### 3.1 Integração Contínua: rollback e teste de carga

### 3.2 Integração Contínua: automatizando a entrega no Kubernetes


## comandos utilizados no curso

```bash
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg lsb-release
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo apt install golang-go
git clone https://github.com/alura-cursos/api_rest_gin_go_2-validacoes-e-testes.git
cd api_rest_gin_go_2-validacoes-e-testes

docker-compose up -d
go run main.go
go test -v main_test.go
```
