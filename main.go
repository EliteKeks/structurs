package main

import (
	"fmt"
)

type html struct{ HTML string }     // создаю структуру html и помечаю ее как тип string с названием HTML
type web struct{ WEB string }       // создаю структуру web и помечаю ее как тип string с названием WEB
type pr struct{ PR string }         // создаю структуру pr и помечаю ее как тип string с названием PR
type seo struct{ SEO string }       // создаю структуру seo и помечаю ее как тип string с названием SEO
type lawyer struct{ LAWYER string } //создаю структуру lawyer и помечаю ее как тип string с названием LAWYER

type People struct {
	name, freelance, stationary string
	age                         int
	html
	pr
	seo
	web
	lawyer
}

type Worker struct {
	freelance  string
	stationary string
	People
}

func main() {

	var ivan = People{name: "Ivan", age: 18} //создаю объект, присваию его к структуре люди, и присваиваю переменные имя и возраст
	var petya = People{name: "Petya", age: 33}
	var slava = People{name: "Slava", age: 21}
	var grisha = People{name: "Grisha", age: 22}
	var kolya = People{name: "Kolya", age: 20}

	ivan.freelance = "Фриланс"
	petya.stationary = "Стационарный работник"
	slava.freelance = "Фриланс"
	grisha.freelance = "Фриланс"
	kolya.stationary = "Стационарный работник"

	ivan.HTML = "Верстальщик,"
	petya.WEB = "Дизайнер,"
	slava.LAWYER = "Юрист,"
	grisha.PR = "Лицо компании,"
	kolya.SEO = "Специалист по развитию,"

	fmt.Println(ivan.name, ivan.age, ivan.HTML, ivan.freelance)
	fmt.Println(petya.name, petya.age, petya.WEB, petya.stationary)
	fmt.Println(slava.name, slava.age, slava.LAWYER, slava.freelance)
	fmt.Println(grisha.name, grisha.age, grisha.PR, grisha.freelance)
	fmt.Println(kolya.name, kolya.age, kolya.SEO, kolya.stationary)
}
