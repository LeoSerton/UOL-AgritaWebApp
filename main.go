package main

import (
	"avianrms/db"
	handler "avianrms/handlers"
	"avianrms/models"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create a new database connection
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Uncoment to populate the database with sample AdUnits and Creatives
	// err = populateDB(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Set up your HTTP routes

	// Define a file server handler for serving static files from the "frontend" directory.
	fs := http.FileServer(http.Dir("frontend"))

	// Register the file server handler for the root route ("/").
	http.Handle("/", fs)

	http.HandleFunc("/create-bird", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateBirdHandler(w, r, db) // Pass db instance to the handler
	})

	http.HandleFunc("/get-bird", func(w http.ResponseWriter, r *http.Request) {
		handler.GetBirdHandler(w, r, db) // Pass db instance to the handler
	})

	http.HandleFunc("/update-bird", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateBirdHandler(w, r, db) // Pass db instance to the handler
	})

	http.HandleFunc("/delete-bird", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteBirdHandler(w, r, db) // Pass db instance to the handler
	})

	http.HandleFunc("/create-cage", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateCageHandler(w, r, db) // Pass db instance to the handler
	})

	http.HandleFunc("/get-cage", func(w http.ResponseWriter, r *http.Request) {
		handler.GetCageHandler(w, r, db) // Pass db instance to the handler
	})

	http.HandleFunc("/get-all-cages", func(w http.ResponseWriter, r *http.Request) {
		handler.GetAllCagesHandler(w, r, db) // Pass db instance to the handler
	})

	http.HandleFunc("/update-cage", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateCageHandler(w, r, db) // Pass db instance to the handler
	})

	http.HandleFunc("/delete-cage", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteCageHandler(w, r, db) // Pass db instance to the handler
	})

	// Start the HTTP server
	addr := "localhost:8080"
	log.Printf("Starting server on %s...", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func populateDB(db *db.DB) error {
	// Generate sample Bird data
	bird1 := models.Bird{
		CageID:       1,
		Specie:       "Parrot",
		Gender:       "Male",
		YearAcquired: 2023,
		FatherBirdID: 1,
		MotherBirdID: 2,
	}

	// Insert the Bird data into the database
	if err := db.InsertBird(bird1); err != nil {
		return err
	}

	// Generate sample Cage data
	cage1 := models.Cage{
		Status:           "BP",
		BreedingBird1ID:  1,
		BreedingBird2ID:  2,
		Timestamp:        time.Now(),
		CleaningRequired: false,
		EggsCurr:         3,
		ChicksCurr:       1,
		Notes:            "Some notes here",
	}

	// Insert the Cage data into the database
	if err := db.InsertCage(cage1); err != nil {
		return err
	}

	return nil
}
