<template>
  <div class="app">
    <header>
      <h1>PingFlow Pronostics</h1>
    </header>
    <main>
      <div class="container">
        <div class="options">
          <button @click="requestJob" :disabled="loading">Demander un job</button>
        </div>
        <div class="results">
          <h2>Résultats</h2>
          <div v-if="loading">Chargement en cours...</div>
          <div v-else>
            <div v-if="dataReady">
              <ul>
                <li v-for="(athlete, index) in athletes" :key="index">
                  {{ athlete.name }}
                </li>
              </ul>
            </div>
            <div v-else>Aucune donnée disponible</div>
          </div>
        </div>
      </div>
    </main>
    <footer>
      <p>&copy; 2023 PingFlow</p>
    </footer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      athletes: [],
      loading: false,
      dataReady: false,
    };
  },
  methods: {
    async requestJob() {
      this.loading = true;
      // Envoyer la demande de job à l'API Node
      // Attendre la réponse et mettre à jour les données
      // (simulate with a timeout)
      setTimeout(() => {
        this.loading = false;
        this.dataReady = true;
        this.athletes = [
          { name: "Jordan Silva" },
          { name: "Emily Rodriguez" },
          { name: "Luis Hernandez" },
          { name: "Mia Santos" },
          { name: "Victor Lima" },
          { name: "Sophia Costa" },
          { name: "Daniel Alves" },
          { name: "Olivia Silva" },
          { name: "Lucas Oliveira" },
          { name: "Isabella Martins" },
          { name: "Bruno Pereira" },
          { name: "Gabriela Santos" },
          { name: "Matheus Costa" },
          { name: "Alice Lima" },
          
          // ... ajoutez d'autres athlètes ici ...
        ];
      }, 2000);
    },
  },
};
</script>

<style scoped>
.app {
  font-family: 'Arial', sans-serif;
  text-align: center;
}

header {
  background-color: #4caf50;
  padding: 1em;
  color: white;
}

h1 {
  margin: 0;
}

main {
  padding: 2em;
}

.container {
  display: flex;
  justify-content: space-between;
}

.options {
  width: 30%;
}

button {
  padding: 1em 2em;
  background-color: #4caf50;
  color: white;
  border: none;
  cursor: pointer;
  font-size: 1em;
  transition: background-color 0.3s;
}

button:disabled {
  background-color: #bfbfbf;
  cursor: not-allowed;
}

.results {
  width: 65%;
}

h2 {
  color: #333;
}

ul {
  list-style: none;
  padding: 0;
}

li {
  margin: 0.5em 0;
  padding: 1em;
  background-color: #f0f0f0;
  border-radius: 5px;
  transition: transform 0.3s;
}

li:hover {
  transform: scale(1.05);
}

footer {
  background-color: #4caf50;
  padding: 1em;
  color: white;
  position: fixed;
  bottom: 0;
  width: 100%;
}
</style>
