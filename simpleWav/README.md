## To get a wav file use:

```go run main.go > out.wav```

## To see the music score:

```go run main.go | xxd -b```


Thank you @jeffowler who wrote [How Rust do](http://blog.jfo.click/how-rust-do/).
Without that post I would have never tried this!


### Improvements

Need to work out why that hardcoded information doesn't create a 1 second song.
Currently getting a 2.5ish second "song".

Need to work out how to abstract away the config with the actual header creation.