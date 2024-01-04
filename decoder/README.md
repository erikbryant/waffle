# Decoder

Each day a new daily waffle is served up. This waffle is available in its raw form as [base64 encoded JSON](https://wafflegame.net/daily1.txt).

The work process when a new waffle comes out is to add it to the regression test suite. Normally this is done by manual transcription of what is seen on the waffle web page. Instead of this manual work, download the JSON, decode it, and generate a serialized representation that can be pasted into the regression suite.

## Using It

```zsh
go run decoder.go
```

Paste the output into ../regress/regress.go.
