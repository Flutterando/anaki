import 'package:anaki_driver_manager/src/anaki_driver_manager.dart';
import 'package:anaki_driver_manager/src/status.dart';
import 'package:anaki_driver_manager/src/types.dart';

void main() async {
  final postgresUrl = 'postgresql://user:password@localhost:5432/mydatabase';

  final driver = AnakiDriverManager();

  try {
    final connectResult = await driver.connect(Config(url: postgresUrl));

    if (connectResult == Status.SQL_SUCCESS) {
      final result = await driver.execute('SELECT version()', {});
      print('PostgreSQL version: ${result.rows.first['version']}');

      final userResult = await driver.execute(
          'SELECT * FROM users WHERE id = :id',
          {'id': 'b628511a-1d38-4d1b-a40c-3c1d4bd57368'});
      print(userResult.toString());
    }
  } catch (e) {
    print('Error: $e');
  } finally {
    await driver.close();
  }
}
