package main

import (
	"errors"
	"fmt"
	"github.com/qamarian-mop/rx-lib"
	"math/big"
	"runtime"
	"time"
)

func init () {
	/*
	create data
	get init-order
	store order

	recursion
	spinner
	*/
}

func spinUp1_KNL (threadTypeId string, waitTime time.Duration) (string, error) {}

func spinUp2_KNL (threadTypeId string, waitTime time.Duration) (string, error) {

	if _, ok := mTHREAD_TYPES [threadTypeId]; ok == false {
		return "", errors.New ("Thread type is not registered in macro mTHREAD_TYPES")
	}

	newInstanceId := "0"

	if count, ok1 := instanceCount [threadTypeId]; ok1 == true {
		_, ok2 := big.NewInt (0).SetString (count)

		if ok2 == false {
			return "", errors.New ("Fatal error: Instance count of thread type does " +
				"not seem to be a valid number.")
		}

		newInstanceId = count
	}

	instanceFullId := threadTypeId . "." . newInstanceId
	commChan [instanceFullId] = struct {
		writChan chan *rxlib.Message
		readChan chan *rxlib.Message
	} {
		make (chan *rxlib.Message, mKNL_DEL_BUFFER_SIZE),
		make (chan *rxlib.Message, mDEL_THR_BUFFER_SIZE)
	}

	startupSignal = 0
	usShutdownSignal [instanceFullId] = 0
	dsShutdownSignal [instanceFullId] = 0

	go del (threadTypeId, newInstanceId, &(commChan.writChan), &(commChan.readChan),
		&startupSignal, &shutdownSignal [instanceFullId])

	waitLimit := time.Now ().Add (waitTime)

	for {
		if startupSignal == 1 {
			break
		}

		if startupSignal >= 2 {
			delete (commChan, instanceFullId)
			delete (shutdownSignal, instanceFullId)
			return "", errors.New ("Problem with thread: Instance of thread could not " +
				"start up successfully.")
		}

		if waitLimit.After (time.Now ()) != true {
			return "", errors.New ("Thread startup timed out.")
		}

		runtime.Gosched ()
	}
}

func main () {
	fmt.Println ("Rexa does only the most fundamental tasks.")
}

var (
	instanceCount map[string]string = map[string]string {}

	commChan map[string]struct {
		writChan chan *rxlib.Message
		readChan chan *rxlib.Message
	} = map[string]struct {
		writChan chan *rxlib.Message
		readChan chan *rxlib.Message
	} {}

	// US (up stream): from del > knl
	usShutdownSignal map[string]byte = map[string]byte {}
	// DS (down stream): from knl > del
	dsShutdownSignal map[string]byte = map[string]byte {}
)
