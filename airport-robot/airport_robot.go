package airportrobot

import "fmt"

// Greeter represents the appropriate greeting in a particular language.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

// SayHello() returns the appropriate greeting for the given name and given
// language.
func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

// Italian represents the appropriate greeting in Italian.
type Italian struct{}

// LanguageName returns the name of the language.
func (l Italian) LanguageName() string {
	return "Italian"
}

// Greet returns the greeting in Italian.
func (l Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// Portuguese represents the appropriate greeting in Portuguese.
type Portuguese struct{}

// LanguageName returns the name of the language.
func (l Portuguese) LanguageName() string {
	return "Portuguese"
}

// Greet returns the greeting in Portuguese.
func (l Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
