// Command mkl.darkbulb turns mkl.lightbulb off. It's the only command you
// should need, really :troll:
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/eclipse/paho.mqtt.golang"
)

func mainError() (err error) {
	// If you're new to Go, know that this MQTT package is a great example
	// of how not to design a library in Go.

	options := mqtt.NewClientOptions().AddBroker("tcp://75.69.78.171:1883")
	client := mqtt.NewClient(options)

	token := client.Connect()
	// No idea what Wait returning false means because the documentation
	// is bad.
	qq := token.Wait()
	if !qq {
		return errors.New(`wait returned false after connect ¯\_(ツ)_/¯`)
	}
	err = token.Error()
	if err != nil {
		return err
	}

	// No idea what qos (the 0) and retained (the bool) do, but this does
	// work.
	token = client.Publish("mkl.lytebulb", 0, false, "{ GPIO-1: 0, GPIO-2: 0, GPIO-3: 0, GPIO-5:0, GPIO-6:0, GPIO-7:0 }")
	qq = token.Wait()
	if !qq {
		return errors.New(`wait returned false after publish ¯\_(ツ)_/¯`)
	}
	err = token.Error()
	if err != nil {
		return err
	}

	return nil
}

func mainCode() int {
	err := mainError()
	if err == nil {
		return 0
	}
	fmt.Fprintf(os.Stderr, "%v: Error: %v\n", filepath.Base(os.Args[0]), err)
	return 1
}

func main() {
	os.Exit(mainCode())
}
