package main

func Start(channel chan Logger) {
	for {
		data := <-channel
		data.Log()
	}
}
