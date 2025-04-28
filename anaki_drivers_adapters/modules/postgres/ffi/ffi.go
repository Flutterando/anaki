package ffi

/*
#include <stdlib.h>
*/
import "C"
import (
	"context"
	"encoding/json"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/driver"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	ffistatus "github.com/flutterando/anaki/anaki_drivers_adapters/shared/ffi"
)

//export Connect
func Connect(configJson *C.char) C.int {
	configStr := C.GoString(configJson)

	var cfg contracts.Config
	err := json.Unmarshal([]byte(configStr), &cfg)
	if err != nil {
		return ffistatus.SQL_ERROR
	}

	driver := &driver.PostgresDriver{}
	ctx := context.Background()

	err = driver.Connect(ctx, cfg)
	if err != nil {
		return ffistatus.SQL_ERROR
	}

	return ffistatus.SQL_SUCCESS
}

//export Close
func Close() C.int {
	driver := &driver.PostgresDriver{}

	if driver == nil {
		return ffistatus.SQL_INVALID_HANDLE
	}

	err := driver.Disconnect()
	if err != nil {
		return ffistatus.SQL_ERROR
	}

	driver = nil
	return ffistatus.SQL_SUCCESS
}
