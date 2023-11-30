package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()

func main() {
	// Initialiser le routeur
	r := mux.NewRouter()

	// Configuration Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Exposer une route API en GET
	r.HandleFunc("/miniprojet", func(w http.ResponseWriter, r *http.Request) {
		// Appel à l'API externe
		apiData, err := callExternalAPI()
		if err != nil {
			log.Printf("Erreur lors de l'appel à l'API externe: %v", err)
			http.Error(w, "Erreur lors de l'appel à l'API externe", http.StatusInternalServerError)
			return
		}

		// Traitement des données si nécessaire

		// Stockage dans Redis avec TTL de 1 minute
		err = storeInRedis(redisClient, "miniprojet:data", apiData, time.Minute)
		if err != nil {
			log.Printf("Erreur lors du stockage dans Redis: %v", err)
			http.Error(w, "Erreur lors du stockage dans Redis", http.StatusInternalServerError)
			return
		}

		// Retourner les données
		json.NewEncoder(w).Encode(apiData)
	}).Methods("GET")

	// Écouter les jobs via pub/sub Redis
	pubsub := redisClient.Subscribe(ctx, "miniprojet:jobs")
	defer pubsub.Close()

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Println(err)
				continue
			}

			// Gérer les jobs reçus
			jobData := // Extraire les données du job depuis msg
			apiData, err := processJob(jobData)
			if err != nil {
				log.Printf("Erreur lors du traitement du job: %v", err)
				continue
			}

			// Stockage dans Redis avec TTL de 1 minute
			err = storeInRedis(redisClient, "miniprojet:data", apiData, time.Minute)
			if err != nil {
				log.Printf("Erreur lors du stockage dans Redis: %v", err)
				continue
			}

			// Publier un message via pub/sub pour indiquer que les données sont prêtes
			pubsub.Publish(ctx, "miniprojet:data_ready", "miniprojet:data")
		}
	}()

	// Démarrer le serveur HTTP
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Fonction pour appeler l'API externe
func callExternalAPI() (interface{}, error) {
	// Implémenter la logique d'appel à l'API externe
	return nil, nil
}

// Fonction pour traiter un job
func processJob(jobData string) (interface{}, error) {
	// Implémenter la logique de traitement du job
	return nil, nil
}

// Fonction pour stocker dans Redis avec TTL
func storeInRedis(client *redis.Client, key string, data interface{}, ttl time.Duration) error {
	// Convertir les données en JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Stocker dans Redis avec TTL
	err = client.Set(ctx, key, jsonData, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}
