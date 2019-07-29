// Command mkl.darkbulb deals with the Big Room Studios office.
//
// * With no arguments, it turns mkl.lyghtbulb off.
// * Sometimes the dark is the problem so `-l` turns mkl.lyghtbulb on.
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

	var command string
	var topic string
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-l":
			topic = "mkl.lytebulb"
			command = `{L1:1,L2:1,L3:1,L4:1,L5:1,L6:1}`
		default:
			return fmt.Errorf("unknown switch: %s", os.Args[1])
		}
	} else {
		topic = "mkl.lytebulb"
		command = `{L1:0,L2:0,L3:0,L4:0,L5:0,L6:0}`
	}

	addr := "tcp://mqtt.bigroomstudios.com:1883"

	options := mqtt.NewClientOptions().AddBroker(addr)
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

	// QoS 1 is supposed to ensure the instruction is received at least once.
	// No idea what retained (the bool) does and the docs don't help.
	token = client.Publish(topic, 1, false, command)
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
