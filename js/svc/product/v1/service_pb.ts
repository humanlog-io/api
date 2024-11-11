// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file svc/product/v1/service.proto (package svc.product.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Cursor } from "../../../types/v1/cursor_pb";
import { Product } from "../../../types/v1/product_pb";

/**
 * @generated from message svc.product.v1.ListProductRequest
 */
export class ListProductRequest extends Message<ListProductRequest> {
  /**
   * @generated from field: types.v1.Cursor cursor = 1;
   */
  cursor?: Cursor;

  /**
   * @generated from field: int32 limit = 2;
   */
  limit = 0;

  /**
   * @generated from field: string category = 3;
   */
  category = "";

  constructor(data?: PartialMessage<ListProductRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.product.v1.ListProductRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "cursor", kind: "message", T: Cursor },
    { no: 2, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "category", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListProductRequest {
    return new ListProductRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListProductRequest {
    return new ListProductRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListProductRequest {
    return new ListProductRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListProductRequest | PlainMessage<ListProductRequest> | undefined, b: ListProductRequest | PlainMessage<ListProductRequest> | undefined): boolean {
    return proto3.util.equals(ListProductRequest, a, b);
  }
}

/**
 * @generated from message svc.product.v1.ListProductResponse
 */
export class ListProductResponse extends Message<ListProductResponse> {
  /**
   * @generated from field: types.v1.Cursor next = 1;
   */
  next?: Cursor;

  /**
   * @generated from field: repeated svc.product.v1.ListProductResponse.ListItem items = 2;
   */
  items: ListProductResponse_ListItem[] = [];

  constructor(data?: PartialMessage<ListProductResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.product.v1.ListProductResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next", kind: "message", T: Cursor },
    { no: 2, name: "items", kind: "message", T: ListProductResponse_ListItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListProductResponse {
    return new ListProductResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListProductResponse {
    return new ListProductResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListProductResponse {
    return new ListProductResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListProductResponse | PlainMessage<ListProductResponse> | undefined, b: ListProductResponse | PlainMessage<ListProductResponse> | undefined): boolean {
    return proto3.util.equals(ListProductResponse, a, b);
  }
}

/**
 * @generated from message svc.product.v1.ListProductResponse.ListItem
 */
export class ListProductResponse_ListItem extends Message<ListProductResponse_ListItem> {
  /**
   * @generated from field: types.v1.Product product = 1;
   */
  product?: Product;

  constructor(data?: PartialMessage<ListProductResponse_ListItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.product.v1.ListProductResponse.ListItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "product", kind: "message", T: Product },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListProductResponse_ListItem {
    return new ListProductResponse_ListItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListProductResponse_ListItem {
    return new ListProductResponse_ListItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListProductResponse_ListItem {
    return new ListProductResponse_ListItem().fromJsonString(jsonString, options);
  }

  static equals(a: ListProductResponse_ListItem | PlainMessage<ListProductResponse_ListItem> | undefined, b: ListProductResponse_ListItem | PlainMessage<ListProductResponse_ListItem> | undefined): boolean {
    return proto3.util.equals(ListProductResponse_ListItem, a, b);
  }
}

