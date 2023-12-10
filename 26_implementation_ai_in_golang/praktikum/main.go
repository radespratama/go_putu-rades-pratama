package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Laptop struct {
	Name          string `json:"name"`
	Budget        int    `json:"budget"`
	Category      string `json:"category"`
	Specification string `json:"specification"`
}

var laptops []Laptop

func init() {
	laptops = []Laptop{
		{
			Name:     "Asus ROG Zephyrus G14",
			Budget:   15000000,
			Category: "Gaming",
			Specification: `
                Processor: AMD Ryzen 9 5900HS
                Graphics: NVIDIA GeForce RTX 3070
                RAM: 16GB DDR4
                Storage: 512GB SSD
            `,
		},
		{
			Name:     "Dell XPS 13",
			Budget:   12000000,
			Category: "Kreatif",
			Specification: `
                Processor: Intel Core i7-12500H
                Graphics: NVIDIA GeForce RTX 3050
                RAM: 16GB LPDDR5
                Storage: 512GB SSD
            `,
		},
		{
			Name:     "Acer Aspire 5",
			Budget:   7000000,
			Category: "Menengah",
			Specification: `
                Processor: Intel Core i5-12500H
                Graphics: Intel Iris Xe
                RAM: 8GB DDR4
                Storage: 256GB SSD
            `,
		},
	}
}

func main() {
	http.HandleFunc("/recommend", recommend)
	http.ListenAndServe(":8080", nil)
}

func recommend(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Budget  int    `json:"budget"`
		Purpose string `json:"purpose"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	results := []Laptop{}
	for _, laptop := range laptops {
		if laptop.Category == req.Purpose && laptop.Budget <= req.Budget {
			results = append(results, laptop)
		}
	}

	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}{
		Status: "success",
		Data:   fmt.Sprintf("With a budget of Rp %d, you can get a high-performance %s laptop that can handle the latest games and provide a smooth gaming experience. Here's a recommendation for a %s laptop within your budget: %s", req.Budget, req.Purpose, req.Purpose, results[0].Name),
	})
}
