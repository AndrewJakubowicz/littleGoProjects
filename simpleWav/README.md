First ever attempt to use a language I'm learning to make something completely out of my comfort zone.
Don't take any of this code as idiomatic go code.
This is a very rough first attempt.

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


