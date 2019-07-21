package main

/* thread type id, instance id, init function of thread, dnit function of thread, read channel for
delegate (to read message from kernel), write channel for delegate (to write message to kernel),
start-up signal, shutdown signal */

/* Possible values of start-up and shutdown signal:

    start-up signal: 0 - yet to start-up; 1 - has started up; 2 - fatal error; x (where x >= 3)
would be considered invalid data, and would be treated as value "2". Meaning of "x" might change
in the future.

    shutdown signal: 0 - yet to shutdown; 1 - has shutdown; x (where x >= 2) would be treated
considered invalid data, and would be treated as value "1". Meaning of "x" might change in the
future. */

func del (threadTypeID, instanceID string, readFromChan <-chan *Message, sendToChan chan<-
	*Message, startupSignal, shutdowSignal byte) {
	//
}
