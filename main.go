package main

import (
	"fmt"
	"os"
	"tdd/mock"
	"time"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {

	fmt.Println(Hello("world", ""))

	// depinj.Greet(os.Stdout, "Elodie")
	// http.ListenAndServe(":5000", http.HandlerFunc(depinj.MyGreeterHandler))

	sleeper := &mock.ConfigurableSleeper{DurationConfig: 1 * time.Second, SleepConfig: time.Sleep}
	mock.Countdown(os.Stdout, sleeper)

}
