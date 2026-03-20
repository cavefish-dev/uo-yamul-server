import 'dart:developer';

import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:uo_yamul_dashboard/domain/repository/auth.dart';
import 'package:uo_yamul_dashboard/presentation/home/home_page.dart';
import 'package:uo_yamul_dashboard/presentation/loading_page.dart';
import 'package:uo_yamul_dashboard/presentation/login/login_page.dart';
import 'package:uo_yamul_dashboard/presentation/login/login_page_route.dart';
import 'package:uo_yamul_dashboard/presentation/maps/maps_page.dart';

import 'service_locator.dart';

@AutoRouterConfig()
class AppRouter extends RootStackRouter {
  @override
  RouteType get defaultRouteType =>
      RouteType.material(); //.cupertino, .adaptive ..etc

  @override
  List<AutoRoute> get routes => [
        AutoRoute(page: LoginPageRoute.pageInfo),
        _createRoute(
            HomePage.routeName, initialRoute: true, (data) => const HomePage()),
        _createRoute(MapsPage.routeName, (data) => const MapsPage()),
        _createRoute(LoadingPage.routeName, (data) => const LoadingPage()),
      ];

  @override
  late final List<AutoRouteGuard> guards = [
    _LoggedInGuard(this),
    // add more guards here
  ];

  AutoRoute _createRoute(String path, Widget Function(RouteData) builder,
      {bool initialRoute = false}) {
    final cleanPath = path.startsWith('/') ? path.substring(1) : path;
    return AutoRoute(
        page: PageInfo(cleanPath, builder: builder), initial: initialRoute);
  }
}

class _LoggedInGuard extends AutoRouteGuard {
  final AppRouter _router;
  _LoggedInGuard(this._router);

  @override
  void onNavigation(NavigationResolver resolver, StackRouter router) async {
    final loginRouteName = LoginPage.routeName.replaceFirst('/', '');
    if (resolver.routeName == loginRouteName) {
      // Avoids bug of login routing to itself
      if (router.current.name == loginRouteName) return;
      return resolver.next();
    }

    log('[${resolver.routeName}] Checkin permissions');
    var loginInfo = await sl<AuthRepository>().getLoginInfo();
    // TODO add some proper access control
    if (loginInfo != null) return resolver.next();
    log('[${resolver.routeName}] Cannot access. Redirecting to ${LoginPage.routeName}');
    _router.popAndPush(LoginPageRoute(redirectTo: resolver.routeName));
  }
}
