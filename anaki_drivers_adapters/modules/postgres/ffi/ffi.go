package ffi

/*
#include <stdlib.h>
*/
import "C"
import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/driver"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/types"
)

//export Connect
func Connect(configJson *C.char) *C.char {
	configStr := C.GoString(configJson)

	var cfg types.Config
	err := json.Unmarshal([]byte(configStr), &cfg)
	if err != nil {
		return C.CString(fmt.Sprintf("Invalid config JSON: %v", err))
	}

	driver := &driver.PostgresDriver{}
	ctx := context.Background()

	err = driver.Connect(ctx, cfg)
	if err != nil {
		return C.CString(fmt.Sprintf("Connection failed: %v", err))
	}

	return C.CString("Connection successful")
}

//export Close
func Close() *C.char {
	driver := &driver.PostgresDriver{}

	err := driver.Disconnect()
	if err != nil {
		return C.CString(fmt.Sprintf("Error disconnecting database: %v", err))
	}

	driver = nil

	return C.CString("Database disconnected successfully")
}
