package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var redisClient *redis.Client

func main() {
	// Initialiser le client Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Adresse de votre serveur Redis
		DB:   0,
	})

	// Créer le routeur HTTP avec Gorilla Mux
	router := mux.NewRouter()
	router.HandleFunc("/miniprojet", handleMiniprojet).Methods("GET")

	// Démarrer le serveur HTTP
	port := 8080
	log.Printf("Serveur HTTP écoutant sur le port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func handleMiniprojet(w http.ResponseWriter, r *http.Request) {
	// Effectuer l'appel à l'API externe pour récupérer la liste des athlètes
	athletes, err := fetchAthletes()
	if err != nil {
		log.Printf("Erreur lors de la récupération des athlètes: %v\n", err)
		http.Error(w, "Erreur lors de la récupération des athlètes", http.StatusInternalServerError)
		return
	}

	// Stocker la liste des athlètes dans Redis avec un TTL de 1 minute
	err = storeInRedis(redisClient, "miniprojet:athletes", athletes, time.Minute)
	if err != nil {
		log.Printf("Erreur lors du stockage dans Redis: %v\n", err)
		http.Error(w, "Erreur lors du stockage dans Redis", http.StatusInternalServerError)
		return
	}

	// Renvoyer la liste des athlètes au client
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(athletes); err != nil {
		log.Printf("Erreur lors de l'encodage JSON: %v\n", err)
		http.Error(w, "Erreur lors de l'encodage JSON", http.StatusInternalServerError)
		return
	}
}

func fetchAthletes() ([]Athlete, error) {
	// Ici, vous devriez effectuer l'appel à l'API externe pour récupérer la liste des athlètes
	// Remplacez cela par votre logique réelle.
	// Vous pouvez utiliser la liste générée aléatoirement pour le moment.

	return generateRandomAthletes(), nil
}

func storeInRedis(client *redis.Client, key string, value interface{}, ttl time.Duration) error {
	// Convertir la valeur en JSON
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

// Stocker dans Redis avec un TTL
err = client.Set(context.Background(), key, jsonValue, ttl).Err()
if err != nil {
    return err
}

}

func generateRandomAthletes() []Athlete {

	var athletes []Athlete
	for i := 1; i <= 25; i++ {
		athlete := Athlete{
			ID:   i,
			Name: fmt.Sprintf("Athlete %d", i),
		}
		athletes = append(athletes, athlete)
	}
	return athletes
}

// Définition de la structure Athlete
type Athlete struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
