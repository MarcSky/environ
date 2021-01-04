# environ #

Library for converting environment variable to certain type

## Smooth conversation #
Smooth conversation (with default value if env not exist) string to
- int
- int64
- uin64
- string (with default value)
- float64

Example
`ServerPort := environ.StrGetEnv("SERVER_PORT", 80)`

##Strict conversation #
Strict conversation (with **panic** if var not exist) string to
- int
- int64
- uin64
- string (with default value)
- float64

Example
`ServerPort := environ.MustGetString("SERVER_PORT")`

