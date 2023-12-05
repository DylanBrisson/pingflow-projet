<template>
  <div>
    <h1>Mini-Projet VueJS</h1>

    <button @click="requestJob">Demander un job</button>

    <div v-if="result">
      <h2>Résultat du job :</h2>
      <pre>{{ result }}</pre>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      result: null,
    };
  },
  methods: {
    async requestJob() {
      try {
        // Demander un job à l'API NodeJS
        const response = await fetch('http://localhost:3000/request-job/1', {
          method: 'GET',
        });

        // Récupérer le résultat du job depuis l'API NodeJS
        const result = await response.json();
        this.result = result.message;

        // Attendre quelques secondes avant de réinitialiser le résultat
        setTimeout(() => {
          this.result = null;
        }, 5000);
      } catch (error) {
        console.error('Erreur lors de la demande de job :', error);
      }
    },
  },
};
</script>

<style scoped>
h1 {
  text-align: center;
  margin-bottom: 20px;
}

button {
  font-size: 16px;
  padding: 10px;
  cursor: pointer;
}

pre {
  background-color: #f0f0f0;
  padding: 10px;
  border-radius: 5px;
  overflow-x: auto;
}
</style>