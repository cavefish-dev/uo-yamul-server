//
//  Generated code. Do not modify.
//  source: yamul-dashboard-maps.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class GetMapsResponse extends $pb.GeneratedMessage {
  factory GetMapsResponse({
    $core.Iterable<GetMapsResponseItem>? items,
  }) {
    final $result = create();
    if (items != null) {
      $result.items.addAll(items);
    }
    return $result;
  }
  GetMapsResponse._() : super();
  factory GetMapsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMapsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetMapsResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'dashboard.maps'), createEmptyInstance: create)
    ..pc<GetMapsResponseItem>(1, _omitFieldNames ? '' : 'items', $pb.PbFieldType.PM, subBuilder: GetMapsResponseItem.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMapsResponse clone() => GetMapsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMapsResponse copyWith(void Function(GetMapsResponse) updates) => super.copyWith((message) => updates(message as GetMapsResponse)) as GetMapsResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetMapsResponse create() => GetMapsResponse._();
  GetMapsResponse createEmptyInstance() => create();
  static $pb.PbList<GetMapsResponse> createRepeated() => $pb.PbList<GetMapsResponse>();
  @$core.pragma('dart2js:noInline')
  static GetMapsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMapsResponse>(create);
  static GetMapsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<GetMapsResponseItem> get items => $_getList(0);
}

class GetMapsResponseItem extends $pb.GeneratedMessage {
  factory GetMapsResponseItem({
    $core.int? index,
    $core.String? name,
  }) {
    final $result = create();
    if (index != null) {
      $result.index = index;
    }
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  GetMapsResponseItem._() : super();
  factory GetMapsResponseItem.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMapsResponseItem.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetMapsResponseItem', package: const $pb.PackageName(_omitMessageNames ? '' : 'dashboard.maps'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'index', $pb.PbFieldType.OU3)
    ..aOS(2, _omitFieldNames ? '' : 'name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMapsResponseItem clone() => GetMapsResponseItem()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMapsResponseItem copyWith(void Function(GetMapsResponseItem) updates) => super.copyWith((message) => updates(message as GetMapsResponseItem)) as GetMapsResponseItem;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetMapsResponseItem create() => GetMapsResponseItem._();
  GetMapsResponseItem createEmptyInstance() => create();
  static $pb.PbList<GetMapsResponseItem> createRepeated() => $pb.PbList<GetMapsResponseItem>();
  @$core.pragma('dart2js:noInline')
  static GetMapsResponseItem getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMapsResponseItem>(create);
  static GetMapsResponseItem? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get index => $_getIZ(0);
  @$pb.TagNumber(1)
  set index($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIndex() => $_has(0);
  @$pb.TagNumber(1)
  void clearIndex() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
