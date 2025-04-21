package ffi

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
)

func SetupDatabaseConnection(config contracts.Config) string {
	configJson, err := json.Marshal(config)
	if err != nil {
		return fmt.Sprintf("Error marshaling config: %v", err)
	}

	cConfigJson := C.CString(string(configJson))
	defer C.free(unsafe.Pointer(cConfigJson))

	result := Connect(cConfigJson)
	return C.GoString(result)
}

func SetupDatabaseClose() string {
	return C.GoString(Close())
}
