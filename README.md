# since
quickly count back from present time, usage intended to be within the scope of 24h<br>
intended usecase: compiled executable in PATH for terminal access

## usage
program accepts time in formats: `HH:mm` and `HH mm`<br>

usage examples: 
```
> since 08:00 (current time: 1200)
< 04h 00m
```

```
> since 13 25 (current time: 1100)
< 21h 35m
```

for hours, values not zero-padded are tolerated<br>
failsafes will be added when deemed necessary

## installation
windows: `go build` in project root

