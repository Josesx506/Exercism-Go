package main

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Language struct {
	langName  string
	greetName string
}

func (l Language) LanguageName() string {
	return l.langName
}

func (l Language) Greet(visitor string) string {
	return fmt.Sprintf("%s %s!", l.greetName, visitor)
}

type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
}

func SayHello(visitor string, g Greeter) string {
	lang := g.LanguageName()
	greeting := g.Greet(visitor)
	return fmt.Sprintf("I can speak %s: %s", lang, greeting)
}

var brazilianLang Language = Language{
	langName:  "Brazilian",
	greetName: "Holá",
}

func main() {

	italianLang := Language{
		langName:  "Italian",
		greetName: "Ciao",
	}

	portugueseLang := Language{
		langName:  "Portuguese",
		greetName: "Olá",
	}

	fmt.Println(SayHello("Joses", italianLang))
	fmt.Println(SayHello("Joses", portugueseLang))
	fmt.Println(SayHello("Joses", brazilianLang))
}

// Reusing method names for different Structs
type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(visitor string) string {
	return fmt.Sprintf("Ciao %s!", visitor)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("Olá %s!", visitor)
}
