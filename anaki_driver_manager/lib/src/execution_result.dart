class ExecutionResult {
  final List<Map<String, dynamic>> rows;
  final int rowsAffected;

  ExecutionResult({
    required this.rows,
    required this.rowsAffected,
  });

  factory ExecutionResult.fromJson(Map<String, dynamic> json) {
    final rows =
        (json['Rows'] as List?)?.map((e) => e as Map<String, dynamic>) ?? [];
    return ExecutionResult(
      rows: rows.toList(),
      rowsAffected: json['RowsAffected'] as int? ?? 0,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'Rows': rows,
      'RowsAffected': rowsAffected,
    };
  }

  @override
  String toString() {
    return 'ExecutionResult(Rows: ${rows.length} rows, RowsAffected: $rowsAffected)';
  }
}
