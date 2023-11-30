const express = require('express');
const redis = require('redis');
const { promisify } = require('util');

const app = express();
const port = 3000;

// Création du client Redis
const redisClient = redis.createClient();

// Convertir les fonctions Redis en Promises pour une utilisation asynchrone
const getAsync = promisify(redisClient.get).bind(redisClient);
const publishAsync = promisify(redisClient.publish).bind(redisClient);

// Endpoint pour demander un job
app.get('/request-job/:athleteId', async (req, res) => {
  const athleteId = req.params.athleteId;

  // Publier le job via pub/sub Redis
  await publishAsync('miniprojet:jobs', athleteId);

  res.json({ message: 'Job demandé avec succès' });
});

// Endpoint pour récupérer les résultats du job
app.get('/get-result', async (req, res) => {
  try {
    // Récupérer les résultats depuis Redis
    const result = await getAsync('miniprojet:data');
    const parsedResult = JSON.parse(result);

    res.json(parsedResult);
  } catch (error) {
    console.error(error);
    res.status(500).json({ error: 'Erreur lors de la récupération des résultats' });
  }
});

// Démarrer le serveur
app.listen(port, () => {
  console.log(`API NodeJS démarrée sur le port ${port}`);
});
