package main

import (
	"os"

	"github.com/pchchv/golog"
	"github.com/pchchv/torrent-client/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		golog.Fatal(err.Error())
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		golog.Fatal(err.Error())
	}
}
