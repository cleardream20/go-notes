package main

import (
	"fmt"
	"io"
	"os"
	// "time"
	"flag"
	"strings"
)

type Speaker interface {
	Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
	return "Haji wang!"
}

type Cat struct {}

func (c Cat) Speak() string {
	return "Haji mi!"
}

func aniSpeak() {
	animals := []Speaker{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

func main() {
	aniSpeak()
	theSameInterface()
	cmdParseEx()
}

func theSameInterface() {
	var w io.Writer
	w = os.Stdout // √ os.Stdout所属类 *os.File也实现了Write方法
	// w = time.Second × // time.Duration does not implement io.Writer (missing method Write)
	w.Write([]byte("hello"))
	print("\n")
}

type Temperature float64

// 1.实现String() & Set()方法
func (t *Temperature) String() string {
    return fmt.Sprintf("%.2f°C", *t)
}

func (t *Temperature) Set(s string) error {
    var value float64
    var unit string
    _, err := fmt.Sscanf(s, "%f%s", &value, &unit)
    if err != nil {
        return err
    }
    switch strings.ToUpper(unit) {
    case "C":
        *t = Temperature(value)
    case "F":
        *t = Temperature((value - 32) * 5 / 9)
    default:
        return fmt.Errorf("invalid unit %q", unit)
    }
    return nil
}

func cmdParseEx() {
    var temp Temperature
	// 1.flag.Var()
	flag.Var(&temp, "temp", "Temperature in Celsius (e.g., 20C) or Fahrenheit (e.g., 68F)")
	// 2.flag.Parse()
    flag.Parse()
    fmt.Printf("Current temperature: %v\n", temp)
}