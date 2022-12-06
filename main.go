package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/nomnemonic/nomnemonic"
)

const (
	// nomnemonic library version
	_versionLibrary = nomnemonic.Version

	// nomnemonic algorithm version
	_versionAlgorithm = nomnemonic.VersionAlgorithm
)

//go:embed data/english.txt
var _words []byte

var _mnemonicer nomnemonic.Mnemonicer

func init() {
	var err error
	_mnemonicer, err = nomnemonic.New(strings.Split(string(_words), "\n"))
	if err != nil {
		panic("couldn't load words")
	}
}

func main() {
	fmt.Println("nomnemonic: deterministic mnemonic generator")
	js.Global().Set("generate", js.FuncOf(generate))
	js.Global().Set("libraryVersion", js.FuncOf(libraryVersion))
	js.Global().Set("algorithmVersion", js.FuncOf(algorithmVersion))
	select {}
}

func generate(this js.Value, inputs []js.Value) interface{} {
	identifier := inputs[0].String()
	password := inputs[1].String()
	passcode := inputs[2].String()
	size, _ := strconv.Atoi(inputs[3].String())

	words, err := _mnemonicer.Generate(identifier, password, passcode, size)
	if err != nil {
		return "Error: " + err.Error()
	}

	return strings.Join(words, " ")
}

func libraryVersion(this js.Value, inputs []js.Value) interface{} {
	return _versionLibrary
}

func algorithmVersion(this js.Value, inputs []js.Value) interface{} {
	return _versionAlgorithm
}
