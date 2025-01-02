# RPGda0 Starter CLI

A simple, flexible CLI for bootstraping full-stack web applications using react, postgreSQL, go, docker, auth0. <br/>
This CLI provides a boilerplate setup for quickly starting your next project with these technologies. <br/>
It includes ready-to-use configurations for backend &amp; frontend integration, database setup, user authentication, &amp; docker.
- R - react
- P - postgreSQL
- G - golang
- d - docker
- a0 - auth0
<br/>

![chrome-capture-2024-12-28](https://github.com/user-attachments/assets/3806750c-780e-4db3-907f-7d4380ce4519)

## Prerequisites
- Docker
- Node
- Auth0 Account with required values in .env
- Git

## Installation steps

### Installation using .tar.gz in Linux and Mac
- [Download](https://github.com/gorvk/RPGda0-starter-cli/releases) the suitable .tar.gz
- Move the download to the preffered installation directory.
- Run following command to extract and make the executable.
```bash
tar xf RPGda0-starter-cli_Linux_x86_64.tar.gz && sudo chmod +x ./RPGda0-starter-cli
```
- Export path of directory in .profile and source the .profile.
```bash
export PATH=$PATH:</path/to/installation/directory>
```
### Installation using .zip in Windows
- [Download](https://github.com/gorvk/RPGda0-starter-cli/releases) the suitable .zip
- Move the download to the preffered installation directory.
- extract the zip
- Export path of directory in Environment variables.

## How to run the CLI
- After the successfull installation just run following command to generate the starter project in the preferred directory.
 ```bash
RPGda0-starter-cli
```
- Fill all the required information to generate the .env

## Steps to Run DB, API and Client

### Steps to run PSQL DB
- Open new terminal
- In root directory, run database container.
```bash
docker compose up starter_db
```

### Steps to run Migration
- Open new terminal
- Uncomment following line in docker-compose.yml
```bash
# command: go run ./cmd/migrate
```
- Run following command.
```bash
docker compose up starter_api
```

### Steps to run API Server
- Open new terminal
- Comment back following line in docker-compose.yml
```bash
command: go run ./cmd/migrate
```
- Run following command to start api server.
```bash
docker compose up starter_api
```

### Steps to run Web Client
- Open new terminal
- Change directory to web.
- Install Dependencies.
```bash
npm i
```
- run client.
```bash
npm run start
```

## Running PORTS
- API will be running at port 9090
- Web app will be running at port 5173
