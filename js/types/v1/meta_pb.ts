// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/meta.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message types.v1.ReqMeta
 */
export class ReqMeta extends Message<ReqMeta> {
  /**
   * @generated from field: int64 machine_id = 2;
   */
  machineId = protoInt64.zero;

  constructor(data?: PartialMessage<ReqMeta>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.ReqMeta";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 2, name: "machine_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReqMeta {
    return new ReqMeta().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReqMeta {
    return new ReqMeta().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReqMeta {
    return new ReqMeta().fromJsonString(jsonString, options);
  }

  static equals(a: ReqMeta | PlainMessage<ReqMeta> | undefined, b: ReqMeta | PlainMessage<ReqMeta> | undefined): boolean {
    return proto3.util.equals(ReqMeta, a, b);
  }
}

/**
 * @generated from message types.v1.ResMeta
 */
export class ResMeta extends Message<ResMeta> {
  /**
   * @generated from field: int64 machine_id = 2;
   */
  machineId = protoInt64.zero;

  constructor(data?: PartialMessage<ResMeta>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.ResMeta";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 2, name: "machine_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResMeta {
    return new ResMeta().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResMeta {
    return new ResMeta().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResMeta {
    return new ResMeta().fromJsonString(jsonString, options);
  }

  static equals(a: ResMeta | PlainMessage<ResMeta> | undefined, b: ResMeta | PlainMessage<ResMeta> | undefined): boolean {
    return proto3.util.equals(ResMeta, a, b);
  }
}

