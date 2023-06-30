package engine

type PValue = any
type PError = any

type PString = string
type PInteger = int64
type PFloat = float64
type PObject = map[string]PValue
type PArray = []PValue
type PFunction = func(stack Stack, args []PValue) (PValue, PValue)

// type PNil = nil
type PBoolean = bool
