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
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	ffistatus "github.com/flutterando/anaki/anaki_drivers_adapters/shared/ffi"
)

var postgresDriver *driver.PostgresDriver

//export Connect
func Connect(configJson *C.char) C.int {
	if postgresDriver != nil {
		return ffistatus.SQL_SUCCESS
	}

	configStr := C.GoString(configJson)
	var cfg contracts.Config
	err := json.Unmarshal([]byte(configStr), &cfg)
	if err != nil {
		return ffistatus.SQL_ERROR
	}

	postgresDriver = &driver.PostgresDriver{}
	ctx := context.Background()

	err = postgresDriver.Connect(ctx, cfg)
	if err != nil {
		return ffistatus.SQL_ERROR
	}

	return ffistatus.SQL_SUCCESS
}

//export Execute
func Execute(query *C.char, paramsJson *C.char) (*C.char, C.int) {
	if postgresDriver == nil {
		return C.CString("Driver not connected"), ffistatus.SQL_ERROR
	}

	queryStr := C.GoString(query)
	params := make(map[string]interface{})

	if len(C.GoString(paramsJson)) > 0 {
		if err := json.Unmarshal([]byte(C.GoString(paramsJson)), &params); err != nil {
			return C.CString(fmt.Sprintf("Error unmarshaling params: %v", err)), ffistatus.SQL_ERROR
		}
	}

	ctx := context.Background()
	result, err := postgresDriver.Execute(ctx, queryStr, params)
	if err != nil {
		return C.CString(fmt.Sprintf("Query execution failed: %v", err)), ffistatus.SQL_ERROR
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return C.CString(fmt.Sprintf("Error marshaling result: %v", err)), ffistatus.SQL_ERROR
	}

	return C.CString(string(resultJson)), ffistatus.SQL_SUCCESS
}

//export Close
func Close() C.int {
	if postgresDriver == nil {
		return ffistatus.SQL_ERROR
	}

	err := postgresDriver.Disconnect()
	if err != nil {
		return ffistatus.SQL_ERROR
	}

	postgresDriver = nil
	return ffistatus.SQL_SUCCESS
}
