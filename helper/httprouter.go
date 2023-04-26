package helper

import (
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetParamInt(params *httprouter.Params, key string) int {
	value, err := strconv.Atoi(params.ByName(key))
	PanicIfError(err)
	return value
}