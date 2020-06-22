package flags

import (
	"flag"
	"fmt"
)

// Config será el titular de nuestras banderas
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// Setup inicializa una configuración a partir de banderas que
// se pasan
func (c *Config) Setup() {
	// puedes establecer una bandera directamente así:
	// var someVar = flag.String ("flag_name", "default_val", "description")
	// pero en la práctica ponerlo en una estructura es generalmente mejor

	// longhand
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	// shorthand
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")

	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")

	// tipo de variable personalizada
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// GetMessage usa todos los internos
// config vars y devuelve una oración
func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is NOT awesome"
	}

	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())
	return msg
}
