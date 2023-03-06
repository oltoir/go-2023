package main

import (
	"fmt"
	"strings"
)

type Item struct {
	name   string
	price  float64
	rating float64
}

type User struct {
	username string
	password string
	email    string
}

type System struct {
	users []User
	items []Item
}

func (s *System) Register(username, password, email string) error {
	for _, user := range s.users {
		if user.username == username {
			return fmt.Errorf("username already exists")
		}
	}

	s.users = append(s.users, User{username, password, email})
	return nil
}

func (s *System) Authorise(username, password string) bool {
	for _, user := range s.users {
		if user.username == username && user.password == password {
			return true
		}
	}

	return false
}

func (s *System) SearchByName(name string) []Item {
	var items []Item
	for _, item := range s.items {
		if strings.Contains(strings.ToLower(item.name), strings.ToLower(name)) {
			items = append(items, item)
		}
	}

	return items
}

func (s *System) FilterByPrice(min, max float64) []Item {
	var items []Item
	for _, item := range s.items {
		if item.price >= min && item.price <= max {
			items = append(items, item)
		}
	}

	return items
}

func (s *System) FilterByRating(min float64) []Item {
	var items []Item
	for _, item := range s.items {
		if item.rating >= min {
			items = append(items, item)
		}
	}

	return items
}

func (s *System) RateItem(username string, itemName string, rating float64) error {
	var item *Item
	for i, it := range s.items {
		if it.name == itemName {
			item = &s.items[i]
			break
		}
	}

	if item == nil {
		return fmt.Errorf("item not found")
	}

	var ratedByUser bool
	for _, user := range s.users {
		if user.username == username {
			ratedByUser = true
			break
		}
	}

	if !ratedByUser {
		return fmt.Errorf("user not found or not authorised")
	}

	item.rating = (item.rating + rating) / 2
	return nil
}

func main() {
	s := &System{
		users: []User{
			{username: "user1", password: "pass1", email: "user1@example.com"},
			{username: "user2", password: "pass2", email: "user2@example.com"},
		},
		items: []Item{
			{name: "item1", price: 10.0, rating: 4.5},
			{name: "item2", price: 20.0, rating: 3.5},
			{name: "item3", price: 30.0, rating: 5.0},
		},
	}

	// Register a new user
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
		fmt.Println("-", item.name)
	}

	// Filter items by price
	items = s.FilterByPrice(15.0, 25.0)
	fmt.Println("Filter by price results:")
	for _, item := range items {
		fmt.Println("-", item.name)
	}

	// Filter items by rating
	items = s.FilterByRating(4.0)
	fmt.Println("Filter by rating results:")
	for _, item := range items {
		fmt.Println("-", item.name)
	}

	// Rate an item
	err = s.RateItem("user1", "item1", 4.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("item rated successfully")
	}
}
