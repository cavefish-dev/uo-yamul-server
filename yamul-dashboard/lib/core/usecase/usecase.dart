import 'package:dartz/dartz.dart';

typedef UseCaseResponse<T> = Either<String, T>;

abstract class UseCase<T, P> {
  final Duration timeoutDuration;

  UseCase({this.timeoutDuration = const Duration(seconds: 1)});

  Future<UseCaseResponse<T>> call(P param);
}
