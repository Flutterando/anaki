class Status {
  static const int SQL_SUCCESS = 0;
  static const int SQL_ERROR = -1;
  static const int SQL_INVALID_HANDLE = -2;

  static const Map<int, String> statusMessages = {
    SQL_SUCCESS: 'Operation successful',
    SQL_ERROR: 'Operation failed',
    SQL_INVALID_HANDLE: 'Invalid handle'
  };

  static String getStatusMessage(int status) {
    return statusMessages[status] ?? 'Unknown status';
  }
}
