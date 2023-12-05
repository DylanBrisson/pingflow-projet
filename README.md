# Projet PingFlow

Bienvenue dans le projet PingFlow ! Ce projet consiste en une application avec un worker en Go, une API NodeJS, et un frontend VueJS. L'objectif est de démontrer une intégration complète utilisant Redis et pub/sub.

## Configuration requise

- Go
- Node.js
- Vue CLI
- Redis

## Installation

1. **Cloner le dépôt GitHub :**

   ```bash
   git clone https://github.com/DylanBrisson/pingflow-projet.git
   cd pingflow-projet
## API Go

2. ** Accéder au répertoire api-go :**

   ```bash
   cd api-go

3. ** Installer les dépendances Go :**
   ```bash
   go get

4. ** Exécuter l'API Go :**
   ```bash
   go run main.go

## API Node.js

1. **Accéder au répertoire api-node :**
   ```bash
   cd api-node

2. **Installer les dépendances Node.js :**
   ```bash
   npm install

3. **Exécuter l'API NodeJS :**
   ```bash
   npm start

## Front-end - Vue.js

1. **Accéder au répertoire frontend :**
   ```bash
   cd frontend

2. **Installer les dépendances Vue.js :**
   ```bash
   npm install

3. **Exécuter l'application VueJS :**
   ```bash
   npm run serve

  L'application devrait être accessible à l'adresse http://localhost:8080 dans votre navigateur.

## Utilisation

Accéder à l'application VueJS à l'adresse http://localhost:8080.
Interagir avec l'application en demandant des jobs, utilisant la pagination, les filtres, etc.
Explorer le code pour comprendre le fonctionnement interne de l'application.

Merci à vous, enjoy code !
