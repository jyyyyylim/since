# since
quickly count back from present time, usage intended to be within the scope of 24h<br>
intended usecase: compiled executable in PATH for terminal access

## usage
program accepts time in formats: `HH:mm` and `HH mm`<br>

usage examples: 
```
> since 8 02 (current time: 1200)
< 04h 00m
```

```
> since 13:25 (current time: 1100)
< 21h 35m
```

for minutes, only zero-padded singledigit values are tolerated<br>
failsafes will be added when deemed necessary

## compilation
`go build` in project root, assuming go/bin is already in PATH. may vary with env and associated aliases

