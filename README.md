# go-julianday

A Go library to compute [Julian day](https://en.wikipedia.org/wiki/Julian_day) given current a date or `time.Time`.

# Install

`go get github.com/starryalley/go-julianday`

# Usage

Import it:

```
import "github.com/starryalley/go-julianday"
```

Get Julian day number (JDN) (type `int64`) of a given date at noon (12:00PM):


```
j := julianday.Number(2020, time.January, 17)
```


Get Julian date (type `float64`) of a given time:

```
j := julianday.Date(time.Now())
```

