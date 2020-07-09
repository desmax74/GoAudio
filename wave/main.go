package main

import (
	"fmt"
	"github.com/DylanMeeus/Audio/wave/internal"
	"os"
)

func main() {
	fmt.Println("Parsing wave file..")
	args := os.Args
	if len(args) < 1 {
		panic("please provide a file..")
	}
	filename := args[1]
	wave, err := internal.ReadWaveFile(filename)
	if err != nil {
		panic("Could not parse wave file")
	}
	fmt.Printf("Read %v samples\n", len(wave.Samples))

	// now try to write this file
	if err := internal.WriteSamples(wave.Samples, wave.WaveFmt, ""); err != nil {
		panic(err)
	}

	fmt.Println("done")
}