import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:get_it/get_it.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_state.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_state.dart';
import 'package:uo_yamul_dashboard/data/models/auth_singin_params.dart';
import 'package:uo_yamul_dashboard/domain/entities/login_info.dart';
import 'package:uo_yamul_dashboard/domain/repository/auth.dart';
import 'package:uo_yamul_dashboard/main.dart';

class _FakeAuthRepository implements AuthRepository {
  @override
  Future<Either<String, void>> login(AuthLoginParams params) async =>
      Right<String, void>(null);

  @override
  Future<void> logout() async {}

  @override
  Future<LoginInfo?> getLoginInfo() async => null;
}

void main() {
  final sl = GetIt.instance;

  setUp(() {
    sl.registerSingleton<AuthRepository>(_FakeAuthRepository());
    sl.registerSingleton(AuthCubit(AuthStateUnknown()));
    sl.registerSingleton(LoadingMapsCubit(LoadingMapsStateLoading()));
  });

  tearDown(() async {
    await sl.reset();
  });

  testWidgets('Dashboard smoke test - login page is shown', (tester) async {
    await tester.pumpWidget(MyApp());
    await tester.pump(); // settle initial frame

    // Login page should be rendered (no valid session → guard redirects to /login)
    expect(find.text('Submit'), findsOneWidget);
  });
}
