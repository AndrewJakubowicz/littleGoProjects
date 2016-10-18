// Package let's make music!
// Thank you @jeffowler who wrote http://blog.jfo.click/how-rust-do/
// Without that post I would have never tried this!
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

const sampleRateNumber = 44100

// concatByteSlice recursively appends any number of byte slices.
// This means byte slices must all be in big or little endian before concat.
func concatByteSlice(a ...[]byte) []byte {
	if len(a) == 1 {
		return a[0]
	}
	return append(a[0], (concatByteSlice(a[1:]...))...)
}

func main() {
	// This is the buffer that the header will be written to.
	buf := new(bytes.Buffer)

	// WAV spec
	chunkId := []byte("RIFF")
	chunkSize := [4]byte{0, 0, 172, 104} // 36 + subchunk size 2
	format := []byte("WAVE")
	subChunk1ID := []byte("fmt ") // Extra space so it takes up 4 bytes
	subChunk1Size := [4]byte{16, 0, 0, 0}
	audioFormat := [2]byte{1, 0}
	numChannels := [2]byte{1, 0}
	sampleRate := [4]byte{68, 172, 0, 0} // 44.1kHz

	byteRate := [4]byte{68, 172, 0, 0} // samplerate * number of channels * (bits per sample / 8)
	blockAlign := [2]byte{1, 0}
	bitsPerSample := [2]byte{8, 0}

	// These were manually inputted as little Endian.
	subChunk1Data := concatByteSlice(subChunk1Size[:], audioFormat[:], numChannels[:], sampleRate[:], byteRate[:], blockAlign[:], bitsPerSample[:])

	subChunk2ID := []byte("data")
	// subChunk2Size == numSamples * numChannels * bitsPerSample / 8
	subChunk2Size := [4]byte{0, 0, 172, 68} // for 1 second

	if err := binary.Write(buf, binary.LittleEndian, chunkId); err != nil {
		fmt.Println("binary.Write of chunkId failed:", err)
	}

	if err := binary.Write(buf, binary.LittleEndian, chunkSize); err != nil {
		fmt.Println("binary.Write of chunkSize failed:", err)
	}

	if err := binary.Write(buf, binary.BigEndian, append(format, subChunk1ID...)); err != nil {
		fmt.Println("binary.Write of chunkSize failed:", err)
	}

	if err := binary.Write(buf, binary.BigEndian, subChunk1Data); err != nil {
		fmt.Println("binary.Write of subChunk1Data failed:", err)
	}

	if err := binary.Write(buf, binary.BigEndian, subChunk2ID); err != nil {
		fmt.Println("binary.Write of chunkSize failed:", err)
	}

	if err := binary.Write(buf, binary.LittleEndian, subChunk2Size); err != nil {
		fmt.Println("binary.Write of subChunk2Size failed:", err)
	}

	// Out header
	os.Stdout.Write(buf.Bytes())

	// Out data (melody)
	databuf := new(bytes.Buffer)
	for i := 0; i < sampleRateNumber; i++ {
		binary.Write(databuf, binary.LittleEndian, sawTooth(i))
	}
	os.Stdout.Write(databuf.Bytes())
}

func sawTooth(x int) uint8 {
	return uint8(x + 1%255)
}
