package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Item struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Rating float64 `json:"rating"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type System struct {
	Users []User `json:"users"`
	Items []Item `json:"items"`
}

func (s *System) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range s.Users {
		if u.Username == user.Username {
			http.Error(w, "username already exists", http.StatusBadRequest)
			return
		}
	}

	s.Users = append(s.Users, user)
	w.WriteHeader(http.StatusCreated)
}

func (s *System) AuthoriseHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range s.Users {
		if u.Username == user.Username && u.Password == user.Password {
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "authorization failed", http.StatusUnauthorized)
}

func (s *System) SearchByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	var items []Item
	for _, item := range s.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(name)) {
			items = append(items, item)
		}
	}

	json.NewEncoder(w).Encode(items)
}

func (s *System) FilterByPriceHandler(w http.ResponseWriter, r *http.Request) {
	minStr := r.FormValue("min")
	maxStr := r.FormValue("max")

	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var items []Item
	for _, item := range s.Items {
		if item.Price >= min && item.Price <= max {
			items = append(items, item)
		}
	}

	json.NewEncoder(w).Encode(items)
}

func (s *System) FilterByRatingHandler(w http.ResponseWriter,
	r *http.Request) {
	ratingStr := r.FormValue("rating")

	rating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var items []Item
	for _, item := range s.Items {
		if item.Rating >= rating {
			items = append(items, item)
		}
	}

	json.NewEncoder(w).Encode(items)
}

// Register func
func (s *System) Register(username, password, email string) error {
	for _, u := range s.Users {
		if u.Username == username {
			return fmt.Errorf("username already exists")
		}
	}

	s.Users = append(s.Users, User{Username: username, Password: password, Email: email})
	return nil
}

// Authorise func
func (s *System) Authorise(username, password string) bool {
	for _, u := range s.Users {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

// SearchByName func
func (s *System) SearchByName(name string) []Item {
	var items []Item
	for _, item := range s.Items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(name)) {
			items = append(items, item)
		}
	}
	return items
}

// FilterByPrice func
func (s *System) FilterByPrice(min, max float64) []Item {
	var items []Item
	for _, item := range s.Items {
		if item.Price >= min && item.Price <= max {
			items = append(items, item)
		}
	}
	return items
}

// FilterByRating func
func (s *System) FilterByRating(rating float64) []Item {
	var items []Item
	for _, item := range s.Items {
		if item.Rating >= rating {
			items = append(items, item)
		}
	}
	return items
}

// RateItem func
func (s *System) RateItem(username, name string, rating float64) error {
	for i, item := range s.Items {
		if item.Name == name {
			s.Items[i].Rating = rating
			return nil
		}
	}
	return fmt.Errorf("item not found")
}
func main() {
	s := &System{
		Users: []User{
			{Username: "user1", Password: "pass1", Email: "user1@example.com"},
			{Username: "user2", Password: "pass2", Email: "user2@example.com"},
		},
		Items: []Item{
			{Name: "item1", Price: 10.0, Rating: 4.5},
			{Name: "item2", Price: 20.0, Rating: 3.5},
			{Name: "item3", Price: 30.0, Rating: 5.0},
		},
	} // Register a new user
	err := s.Register("user3", "pass3", "user3@example.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user registered successfully")
	}

	// Authorize a user
	if s.Authorise("user1", "pass1") {
		fmt.Println("user1 authorized")
	} else {
		fmt.Println("authorization failed")
	}

	// Search items by name
	items := s.SearchByName("item1")
	fmt.Println("Search results:")
	for _, item := range items {
		fmt.Println("-", item.Name)
	}

	// Filter items by price
	items = s.FilterByPrice(15.0, 25.0)
	fmt.Println("Filter by price results:")
	for _, item := range items {
		fmt.Println("-", item.Name)
	}

	// Filter items by rating
	items = s.FilterByRating(4.0)
	fmt.Println("Filter by rating results:")
	for _, item := range items {
		fmt.Println("-", item.Name)
	}

	// Rate an item
	err = s.RateItem("user1", "item1", 4.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("item rated successfully")
	}

	// add http endpoints
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		if username == "" || password == "" || email == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}
		err := s.Register(username, password, email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("User registered successfully"))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		json.NewEncoder(w).Encode(s.Users)
	})

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		json.NewEncoder(w).Encode(s.Items)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "" || password == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}
		if s.Authorise(username, password) {
			w.Write([]byte("User authorized"))
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		name := r.FormValue("name")
		if name ==
			"" {
			http.Error(w, "Missing search query", http.StatusBadRequest)
			return
		}
		items := s.SearchByName(name)
		var resp []string
		for _, item := range items {
			resp = append(resp, item.Name)
		}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/filter", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var minPrice, maxPrice, minRating float64
		var err error
		minPriceStr := r.FormValue("minPrice")
		if minPriceStr != "" {
			minPrice, err = strconv.ParseFloat(minPriceStr, 64)
			if err != nil {
				http.Error(w, "Invalid minPrice value", http.StatusBadRequest)
				return
			}
		}
		maxPriceStr := r.FormValue("maxPrice")
		if maxPriceStr != "" {
			maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
			if err != nil {
				http.Error(w, "Invalid maxPrice value", http.StatusBadRequest)
				return
			}
		}
		minRatingStr := r.FormValue("minRating")
		if minRatingStr != "" {
			minRating, err = strconv.ParseFloat(minRatingStr, 64)
			if err != nil {
				http.Error(w, "Invalid minRating value", http.StatusBadRequest)
				return
			}
		}
		items := s.FilterByPrice(minPrice, maxPrice)
		items = s.FilterByRating(minRating)
		var resp []string
		for _, item := range items {
			resp = append(resp, item.Name)
		}
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/rate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		itemName := r.FormValue("itemName")
		ratingStr := r.FormValue("rating")
		if username == "" || password == "" || itemName == "" || ratingStr == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}
		rating, err := strconv.ParseFloat(ratingStr, 64)
		if err != nil {
			http.Error(w, "Invalid rating value", http.StatusBadRequest)
			return
		}
		err = s.RateItem(username, itemName, rating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("Item rated successfully"))
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
