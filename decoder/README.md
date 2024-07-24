# Decoder

Each day a new daily waffle is served up. This waffle is available in its raw form as base64 encoded JSON [daily1.txt](https://wafflegame.net/daily1.txt). There is also a second waffle at [daily2.txt](https://wafflegame.net/daily2.txt).

Each Sunday a new deluxe waffle is served up. This waffle is available in its raw form as base64 encoded JSON [deluxe1.txt](https://wafflegame.net/deluxe1.txt). There is also a second waffle at [deluxe2.txt](https://wafflegame.net/deluxe2.txt).

It is not clear what the difference is between the `1` and `2` versions of the waffles. Maybe country-specific variants?

The work process when a new waffle comes out is to add it to the regression test suite. Normally this is done by manual transcription of what is seen on the waffle web page. Instead of this manual work, download the JSON, decode it, and generate a serialized representation that can be pasted into the regression suite.

## Using It

```zsh
go run decoder.go
```

Paste the output into ../regress/regress.go.

Alternately, generate the signatures from the archive. It is a similar format to the daily file, but there is one extra layer of encoding; botli.

Download the latest archive file from the link (or whatever name it has). Rename it to `archive-daily.json` or `archive-deluxe.json`.

* https://assets.waffle.game/archive-27fddeb9.json (daily)
* https://assets.waffle.game/archive-63f7eed4.json (deluxe)

```zsh
go run decoder.go -file archive-daily.json
```
