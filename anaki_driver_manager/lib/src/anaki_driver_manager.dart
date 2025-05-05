import 'dart:convert';
import 'dart:ffi';
import 'dart:io';
import 'package:anaki_driver_manager/src/status.dart';
import 'package:anaki_driver_manager/src/types.dart';
import 'package:anaki_driver_manager/src/execution_result.dart';
import 'package:ffi/ffi.dart';

class AnakiDriverManager {
  static final DynamicLibrary _lib = DynamicLibrary.open(
      'lib/drivers/${Platform.operatingSystem.toLowerCase()}/amd64/postgresql.so');

  static final _connect = _lib.lookupFunction<Int32 Function(Pointer<Utf8>),
      int Function(Pointer<Utf8>)>('Connect');

  static final _execute = _lib.lookupFunction<
      Pointer<Utf8> Function(Pointer<Utf8>, Pointer<Utf8>),
      Pointer<Utf8> Function(Pointer<Utf8>, Pointer<Utf8>)>('Execute');

  static final _close =
      _lib.lookupFunction<Int32 Function(), int Function()>('Close');

  Future<int> connect(Config config) async {
    final configJson = jsonEncode(config.toJson()).toNativeUtf8();
    try {
      final status = _connect(configJson);
      if (status != Status.SQL_SUCCESS) {
        throw Exception(
            'Failed to connect to database. Status: ${Status.getStatusMessage(status)}');
      }
      return status;
    } finally {
      malloc.free(configJson);
    }
  }

  Future<ExecutionResult> execute(
      String query, Map<String, dynamic> params) async {
    final queryPtr = query.toNativeUtf8();
    final paramsJson = jsonEncode(params).toNativeUtf8();

    try {
      final resultPtr = _execute(queryPtr, paramsJson);
      final result = resultPtr.toDartString();
      return ExecutionResult.fromJson(jsonDecode(result));
    } finally {
      malloc.free(queryPtr);
      malloc.free(paramsJson);
    }
  }

  Future<int> close() async {
    return _close();
  }
}
