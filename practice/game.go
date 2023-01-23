package main

import (
	"fmt"
)

type Unit struct {
	name        string
	gender      string
	class       string
	health      int
	attack      int
	damageTaken int
}

func (u *Unit) TakeDamage(damage int) {
	u.health -= damage
	u.damageTaken += damage
}

type Warrior struct {
	Unit
}

func (w *Warrior) Attack() {
	fmt.Println(w.name, "attacks with a sword for", w.attack, "damage!")
}

type Archer struct {
	Unit
}

func (a *Archer) RangedAttack() {
	fmt.Println(a.name, "shoots an arrow for", a.attack, "damage!")
}

type Tank struct {
	Unit
}

func (t *Tank) AbsorbDamage() {
	fmt.Println(t.name, "absorbs the damage with shield for", t.damageTaken, "damage!")
}

type Support struct {
	Unit
}

func (s *Support) Heal() {
	fmt.Println(s.name, "heals for", s.attack, "health!")
}

func main() {
	fmt.Println("Welcome to the game! Please choose a class, gender, and name for your hero.")
	var class, gender, name string
	fmt.Print("Class (Warrior, Archer, Tank, Support): ")
	fmt.Scan(&class)
	fmt.Print("Gender (Male, Female): ")
	fmt.Scan(&gender)
	fmt.Print("Name: ")
	fmt.Scan(&name)

	var hero Unit
	hero.name = name
	hero.gender = gender
	hero.class = class
	hero.health = 100
	hero.attack = 10
	hero.damageTaken = 0

	switch class {
	case "Warrior":
		hero = Warrior{hero}
	case "Archer":
		hero = Archer{hero}
	case "Tank":
		hero = Tank{hero}
	case "Support":
		hero = Support{hero}
	default:
		fmt.Println("Invalid class choice.")
	}

	fmt.Println("You have chosen", hero.name, "the", hero.gender, hero.class)

	// game loop
	for {
		// get player input
		var action string
		fmt.Print("What would you like to do? (Attack, Ranged Attack, Absorb Damage, Heal): ")
		fmt.Scan(&action)

		// perform action
		switch action {
		case "Attack":
			if hero.class == "Warrior" {
				hero.(*Warrior).Attack()
			} else {
				fmt.Println("Invalid action for your class.")
			}
		case "Ranged Attack":
			if hero.class == "Archer" {
				hero.(*Archer).RangedAttack()
			} else {
				fmt.Println("Invalid action for your class.")
			}
		case "Absorb Damage":
			if hero.class == "Tank" {
				hero.(*Tank).AbsorbDamage()
			} else {
				fmt.Println("Invalid action for your class.")
			}
		case "Heal":
			if hero.class == "Support" {
				hero.(*Support).Heal()
			} else {
				fmt.Println("Invalid action for your class.")
			}
		default:
			fmt.Println("Invalid action.")
		} // check if the game is over
		if hero.health <= 0 {
			fmt.Println("Game over.")
			break
		}
	}
}
