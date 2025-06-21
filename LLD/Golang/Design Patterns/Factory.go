// Implementing Factory Design Patterns
package Design

import (
	"fmt"
	"strings"
)

type Weapons interface {
	getName() string
	setName(name string)
	getPower() int
	setPower(power int)
}

// Thông thường interface và các struct implement các method của nó sẽ tách ra đứng riêng từng file một nhưng vì demo nên mình sẽ làm cho nó gọn
// On a daily basis, an interface and the classes (structs) implementing that interface would be seperated by each one but because I'm currently demo so I'd do it in summary
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

func (g *Gun) setPower(power int) {
	g.power = power
}

// ///////////////////////////////////////////////////////////////////////////////////
//
//	In Golang, we don't have acronym called "Inheritance" but "Composition" which means we can avoid boilerplate code - reusing existed code to be cleaner and
//	more efficient by embedding the class (struct) that has method we want to reuse into another class (struct)
type Glock18 struct {
	Gun
}

func NewGlock18() Weapons {
	return &Glock18{
		Gun: Gun{
			name:  "Glock18",
			power: 3,
		},
	}
}

// ///////////////////////////////////////////////////////////////////////////////////
type AWP struct {
	Gun
}

func NewAwp() Weapons {
	return &AWP{
		Gun: Gun{
			name:  "Awp",
			power: 8,
		},
	}
}

// //////////////////////////////////////////////////////////////////////////////////
//
//	main.go
func NewGun(gunType string) (Weapons, error) {
	if strings.EqualFold(gunType, "glock18") {
		return NewGlock18(), nil
	} else if strings.EqualFold(gunType, "awp") {
		return NewAwp(), nil
	} else {
		return nil, fmt.Errorf("Please enter the supported gun types!")
	}
}
