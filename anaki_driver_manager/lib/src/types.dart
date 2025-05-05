class Config {
  final String url;

  Config({required this.url});

  Map<String, dynamic> toJson() => {
        'url': url,
      };

  @override
  String toString() => 'Config(url: $url)';
}
