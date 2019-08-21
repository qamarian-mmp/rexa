package main

import (
	"fmt"
	"gopkg.in/qamarian-dtp/rnet.v1"
	"gopkg.in/qamarian-dtp/system.v1"
	"gopkg.in/qamarian-etc/slices.v1"
	"gopkg,in/qamarian-mmp/rxlib.v0"
	"runtime"
	"sync"
)

func main () {
	if osLog == nil {
		fmt.Println ("A log was not provided.")
		return
	}
	osLog.Record ("Rexa-based Software (RbS) starting up...", rxlib.LrtStandard)
	if mains == nil {
		osLog.Record ("Value of the main registry is nil.", rxlib.LrtError)
		return
	}
	if len (mains) == 0 {
		osLog.Record ("No main is registered with this OS... nothing to run... " +
			"RbS shutting down...", rxlib.LrtWarning)
		return
	}
	defer func () {
		panicReason = recover ()
		if panicReason != nil {
			osLog.Record ("Software is shutting down due to a panic.",
				rxlib.LrtError)
			fmt.Println (panicReason)
		}
	} ()
	validMains := map[string]*Register {}
	rbsSystem := system.New ()
	for _, reg := range mains {
		if reg.ID () == 0 {
			osLog.Record ("A main is using an empty string as ID.",
				rxlib.LrtWarning)
			return
		}
		dependencies := reg.Dep ()
		for _, dep := range dependencies {
			if dep == "" {
				errMssg := fmt.Sprintf ("Main '%s' is using an empty " +
					"string as the ID of one of its dependencies.",
					reg.ID ())
				osLog.Record (errMssg, rxlib.LrtWarning)
				return
			}
		}
		if reg.StartupFunc () == nil {
			errMssg := fmt.Sprintf ("Main '%s' is using nil as its startup " +
					"function.", reg.ID ())
			osLog.Record (errMssg, rxlib.LrtWarning)
			return
		}
		validMains[reg.ID ()] = reg
		errX := rbsSystem.AddElement (reg.ID (), reg.Dep ())
		if errX != nil {
			errMssg := fmt.Sprintf ("Unable to add main '%s', as element, " +
				"to the data needed to determine this RbS's startup "
				"order. [%s]", reg.ID (), errX.Error ())
			osLog.Record (errMssg, rxlib.LrtError)
			return
		}
	}
	startupOrder, errY, more := rbsSystem.InitOrder ()
	if errY != nil {
		errMssg := fmt.Sprintf ("Unable to create a startup order. [%s [%s]]",
			errY.Error (), more)
		osLog.Record (errMssg, rxlib.LrtError)
		return
	}
	net := rnet.New ()
	shutdownChanLocker := &sync.Mutex {}
	shutdownChan := sync.NewCond (shutdownChanLocker)
	shutdownKeys := map[string]MasterKey
	defer shutdown (slices.RevStringSlice (startupOrder), shutdownKeys)
	for _, someMain := range startupOrder {
		outX := fmt.Sprintf ("Starting up main '%s'...", someMain)
		osLog.Record (outX, rxlib.LrtStandard)
		ppo, errZ := net.NewPPO (somMain)
		if errZ != nil {
			errMssg := fmt.Sprintf ("A communication channel (PPO) could " +
				"not be created for main '%s'. [%s]", someMain,
				errZ.Error ())
			osLog.Record (errMssg, rxlib.LrtError)
			return
		}
		rxkey := rxlib.NewRxKey (ppo, shutdownChan, net)
		var (
			masterKey rxlib.MasterKey = rxkey
			key rxlib.Key = rxkey
		)
		go validMains[someMain].StartupFunc (key)
		for result, note := masterKey.StartupResult ();
			result != rxlib.SrStartedUp {
			if result == rxlib.SrStartupFailed {
				errMssg := fmt.Sprintf ("Unable to startup main '%s'. " +
					"[%s]", someMain, note)
				osLog.Record (errMssg, rxlib.LrtError)
				return
			}
			runtime.Gosched ()
		}
		outY := fmt.Sprintf ("Main '%s' has started up.", someMain)
		osLog.Record (outY, rxlib.LrtStandard)
		runtime.Gosched ()
	}
	shutdownChanLocker.Lock ()
	shutdownChan.Wait ()
	shutdownChanLocker.Unlock ()
}

func shutdown (shutdownOrder []string, shutdownKeys MasterKey) {
	osLog.Record ("Graceful shutdown has started.", rxlib.LrtStandard)
	for _, someMain := range shutdownOrder {
		masterKey := shutdownKeys[someMain]
		masterKey.ShutdownMain ()
		if masterKey.ShutdownState () == SsNotApplicable {
			runtime.Gosched ()
			continue
		}
		for masterKey.ShutdownState () != SsHasShutdown {
			runtime.Gosched ()
		}
		outX := fmt.Sprintf ("Main '%s' has been shutdown.", someMain)
		osLog.Record (outX, rxlib.LrtStandard)
		runtime.Gosched ()
	}
}
