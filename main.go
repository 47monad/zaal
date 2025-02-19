package main

import (
	sercon "github.com/47monad/sercon/lib"
)

func main() {
	// var (
	// 	pklPath    string
	// 	outputPath string
	// )
	//
	// flag.StringVar(&pklPath, "pkl", "", "pkl file name")
	// flag.StringVar(&outputPath, "output", "", "output file name")
	//
	// flag.Parse()
	//
	// if outputPath == "" {
	// 	outputPath = ".sercon/app.json"
	// }
	//
	// if pklPath == "" {
	// 	pklPath = "./config/app.pkl"
	// }
	//
	// err := sercon.Build(pklPath, outputPath)
	// if err != nil {
	// 	panic(err)
	// }

	sercon.BuildCue()

	// fmt.Printf("Config was built inside %s", outputPath)
}
