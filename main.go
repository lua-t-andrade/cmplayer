package main

import (
	"fmt"
	"os"
	"time"
	"log"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/mp3"
)

func usage() {
	fmt.Println("Usage: mplayer [filename] -- play mp3 file")
	fmt.Println("       mplayer dir        -- shows current directory")
	fmt.Println("       mplayer help       -- shows this text")
}


func showFiles() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e)
	}
}

func setArgs() string {
	// I set all the args logic manually, but this could be done
	// with a library.
	args := os.Args[1:]
	var argc int = 0
	for _, _ = range args {
		argc++
	}

	if argc == 0 {
		fmt.Println("Err: What file?")
		usage()
		os.Exit(0)
	}
	
	if argc > 1 {
		fmt.Println("Err: Only one file name at a time")
		usage()
		os.Exit(0)
	}

	if args[0] == "dir" {
		showFiles()
		os.Exit(0)
	} else if args[0] == "help" {
		usage()
		os.Exit(0)
	}
	
	return args[0]
}

func playFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done


}


func main() {
	var fileName string
	fileName = setArgs()
	playFile(fileName)
	
}
