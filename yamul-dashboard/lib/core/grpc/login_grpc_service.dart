import 'package:grpc/grpc_web.dart';
import 'package:uo_yamul_dashboard/generated/grpc/yamul-dashboard-login.pbgrpc.dart';

class LoginGrpcService {
  static Future<DashboardLoginServiceClient> createClient() async {
    final grpcUri = Uri(
      scheme: Uri.base.scheme,
      host: Uri.base.host,
      port: Uri.base.port,
      path: '/grpc/',
    );
    final channel = GrpcWebClientChannel.xhr(grpcUri);
    return DashboardLoginServiceClient(channel);
  }
}
