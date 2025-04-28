package ffi

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"encoding/json"
	"unsafe"

	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	ffistatus "github.com/flutterando/anaki/anaki_drivers_adapters/shared/ffi"
)

func SetupDatabaseConnection(config contracts.Config) int {
	configJson, err := json.Marshal(config)
	if err != nil {
		return ffistatus.SQL_ERROR
	}

	cConfigJson := C.CString(string(configJson))
	defer C.free(unsafe.Pointer(cConfigJson))

	result := Connect(cConfigJson)
	return int(result)
}

func SetupDatabaseClose() int {
	return int(Close())
}
