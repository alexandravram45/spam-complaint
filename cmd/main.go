package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/alexandravram45/spam-complaint/internal/api"
	"github.com/alexandravram45/spam-complaint/internal/config"
	"github.com/alexandravram45/spam-complaint/internal/detector"
	"github.com/alexandravram45/spam-complaint/internal/orchestrator"
)

func main() {
	fmt.Println("Automatic Spam Complaint System")

	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	det := detector.NewSpamDetector(cfg.Spam.Keywords)
	system := orchestrator.NewSystem(det)

	handler := api.NewHandler(
		system.ComplaintService(),
		system.Detector(),
	)

	http.HandleFunc("/complaints", handler.Complaints)

	fmt.Println("API running on :8080")
	http.ListenAndServe(":8080", nil)

	system.Run()
}
