package puppy

import (
	"github.com/RicHardicode/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func Bigbark() string {
	return dog.WhenGrownUp(Bark())
}

func Bigbarks() string {
	return dog.WhenGrownUp(Barks())
}
