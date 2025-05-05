import 'dart:io';

import 'package:anaki_driver_manager/src/anaki_driver_manager.dart';
import 'package:anaki_driver_manager/src/types.dart';
import 'package:anaki_driver_manager/src/status.dart';
import 'package:anaki_driver_manager/src/execution_result.dart';
import 'package:test/test.dart';

void main() {
  final driver = AnakiDriverManager();

  test('should create table and insert data', () async {
    final connStr = Platform.environment['POSTGRES_TEST_DATABASE_URL'];

    final result = await driver.connect(Config(url: connStr!));
    expect(result, equals(Status.SQL_SUCCESS));

    final createResult = await driver.execute('''
      CREATE TABLE test_table (
        id BIGINT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
      )
    ''', {});
    expect(createResult, isA<ExecutionResult>());

    final insertResult = await driver.execute('''
      INSERT INTO test_table (id, name) VALUES 
        (177602, 'Marcos'),
        (177603, 'Jacob'),
        (177604, 'Max')
    ''', {});
    expect(insertResult, isA<ExecutionResult>());

    final selectResult = await driver
        .execute('SELECT * FROM test_table WHERE id = :id', {'id': '177603'});
    expect(selectResult, isA<ExecutionResult>());
    expect(selectResult.rows, isNotEmpty);
    expect(selectResult.rows.first['name'], equals('Jacob'));

    await driver.execute('DROP TABLE IF EXISTS test_table', {});

    await driver.close();
  });
}
