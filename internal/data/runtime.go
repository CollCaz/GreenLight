package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJson() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quoatedJsonValue := strconv.Quote(jsonValue)

	return []byte(quoatedJsonValue), nil
}
