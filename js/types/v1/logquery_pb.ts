// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/logquery.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Val } from "./types_pb";

/**
 * @generated from message types.v1.LogQuery
 */
export class LogQuery extends Message<LogQuery> {
  /**
   * @generated from field: google.protobuf.Timestamp from = 1;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 2;
   */
  to?: Timestamp;

  /**
   * @generated from field: types.v1.Context context = 100;
   */
  context?: Context;

  /**
   * @generated from field: types.v1.Expr query = 101;
   */
  query?: Expr;

  constructor(data?: PartialMessage<LogQuery>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.LogQuery";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "from", kind: "message", T: Timestamp },
    { no: 2, name: "to", kind: "message", T: Timestamp },
    { no: 100, name: "context", kind: "message", T: Context },
    { no: 101, name: "query", kind: "message", T: Expr },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogQuery {
    return new LogQuery().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogQuery {
    return new LogQuery().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogQuery {
    return new LogQuery().fromJsonString(jsonString, options);
  }

  static equals(a: LogQuery | PlainMessage<LogQuery> | undefined, b: LogQuery | PlainMessage<LogQuery> | undefined): boolean {
    return proto3.util.equals(LogQuery, a, b);
  }
}

/**
 * @generated from message types.v1.Context
 */
export class Context extends Message<Context> {
  /**
   * @generated from field: optional types.v1.Expr machine_id = 101;
   */
  machineId?: Expr;

  /**
   * @generated from field: optional types.v1.Expr session_id = 102;
   */
  sessionId?: Expr;

  constructor(data?: PartialMessage<Context>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Context";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "machine_id", kind: "message", T: Expr, opt: true },
    { no: 102, name: "session_id", kind: "message", T: Expr, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Context {
    return new Context().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Context {
    return new Context().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Context {
    return new Context().fromJsonString(jsonString, options);
  }

  static equals(a: Context | PlainMessage<Context> | undefined, b: Context | PlainMessage<Context> | undefined): boolean {
    return proto3.util.equals(Context, a, b);
  }
}

/**
 * @generated from message types.v1.Expr
 */
export class Expr extends Message<Expr> {
  /**
   * @generated from oneof types.v1.Expr.expr
   */
  expr: {
    /**
     * @generated from field: types.v1.Val literal = 101;
     */
    value: Val;
    case: "literal";
  } | {
    /**
     * @generated from field: types.v1.UnaryOp unary = 102;
     */
    value: UnaryOp;
    case: "unary";
  } | {
    /**
     * @generated from field: types.v1.BinaryOp binary = 103;
     */
    value: BinaryOp;
    case: "binary";
  } | {
    /**
     * @generated from field: types.v1.FuncCall func_call = 104;
     */
    value: FuncCall;
    case: "funcCall";
  } | {
    /**
     * @generated from field: types.v1.Identifier identifier = 105;
     */
    value: Identifier;
    case: "identifier";
  } | {
    /**
     * msg.hello.world
     *
     * @generated from field: types.v1.Selector selector = 106;
     */
    value: Selector;
    case: "selector";
  } | {
    /**
     * @generated from field: types.v1.Indexor indexor = 107;
     */
    value: Indexor;
    case: "indexor";
  } | {
    /**
     * @generated from field: types.v1.Pipe pipe = 108;
     */
    value: Pipe;
    case: "pipe";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Expr>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Expr";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "literal", kind: "message", T: Val, oneof: "expr" },
    { no: 102, name: "unary", kind: "message", T: UnaryOp, oneof: "expr" },
    { no: 103, name: "binary", kind: "message", T: BinaryOp, oneof: "expr" },
    { no: 104, name: "func_call", kind: "message", T: FuncCall, oneof: "expr" },
    { no: 105, name: "identifier", kind: "message", T: Identifier, oneof: "expr" },
    { no: 106, name: "selector", kind: "message", T: Selector, oneof: "expr" },
    { no: 107, name: "indexor", kind: "message", T: Indexor, oneof: "expr" },
    { no: 108, name: "pipe", kind: "message", T: Pipe, oneof: "expr" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Expr {
    return new Expr().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Expr {
    return new Expr().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Expr {
    return new Expr().fromJsonString(jsonString, options);
  }

  static equals(a: Expr | PlainMessage<Expr> | undefined, b: Expr | PlainMessage<Expr> | undefined): boolean {
    return proto3.util.equals(Expr, a, b);
  }
}

/**
 * @generated from message types.v1.UnaryOp
 */
export class UnaryOp extends Message<UnaryOp> {
  /**
   * @generated from field: types.v1.UnaryOp.Operator op = 1;
   */
  op = UnaryOp_Operator.INVALID;

  /**
   * @generated from field: types.v1.Expr arg = 2;
   */
  arg?: Expr;

  constructor(data?: PartialMessage<UnaryOp>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.UnaryOp";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "op", kind: "enum", T: proto3.getEnumType(UnaryOp_Operator) },
    { no: 2, name: "arg", kind: "message", T: Expr },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UnaryOp {
    return new UnaryOp().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UnaryOp {
    return new UnaryOp().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UnaryOp {
    return new UnaryOp().fromJsonString(jsonString, options);
  }

  static equals(a: UnaryOp | PlainMessage<UnaryOp> | undefined, b: UnaryOp | PlainMessage<UnaryOp> | undefined): boolean {
    return proto3.util.equals(UnaryOp, a, b);
  }
}

/**
 * @generated from enum types.v1.UnaryOp.Operator
 */
export enum UnaryOp_Operator {
  /**
   * @generated from enum value: INVALID = 0;
   */
  INVALID = 0,

  /**
   * @generated from enum value: NOT = 1;
   */
  NOT = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(UnaryOp_Operator)
proto3.util.setEnumType(UnaryOp_Operator, "types.v1.UnaryOp.Operator", [
  { no: 0, name: "INVALID" },
  { no: 1, name: "NOT" },
]);

/**
 * @generated from message types.v1.BinaryOp
 */
export class BinaryOp extends Message<BinaryOp> {
  /**
   * @generated from field: types.v1.Expr lhs = 1;
   */
  lhs?: Expr;

  /**
   * @generated from field: types.v1.BinaryOp.Operator op = 2;
   */
  op = BinaryOp_Operator.INVALID;

  /**
   * @generated from field: types.v1.Expr rhs = 3;
   */
  rhs?: Expr;

  constructor(data?: PartialMessage<BinaryOp>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.BinaryOp";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "lhs", kind: "message", T: Expr },
    { no: 2, name: "op", kind: "enum", T: proto3.getEnumType(BinaryOp_Operator) },
    { no: 3, name: "rhs", kind: "message", T: Expr },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BinaryOp {
    return new BinaryOp().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BinaryOp {
    return new BinaryOp().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BinaryOp {
    return new BinaryOp().fromJsonString(jsonString, options);
  }

  static equals(a: BinaryOp | PlainMessage<BinaryOp> | undefined, b: BinaryOp | PlainMessage<BinaryOp> | undefined): boolean {
    return proto3.util.equals(BinaryOp, a, b);
  }
}

/**
 * @generated from enum types.v1.BinaryOp.Operator
 */
export enum BinaryOp_Operator {
  /**
   * @generated from enum value: INVALID = 0;
   */
  INVALID = 0,

  /**
   * @generated from enum value: LOG_AND = 101;
   */
  LOG_AND = 101,

  /**
   * @generated from enum value: LOG_OR = 102;
   */
  LOG_OR = 102,

  /**
   * @generated from enum value: NUM_ADD = 201;
   */
  NUM_ADD = 201,

  /**
   * @generated from enum value: NUM_SUB = 202;
   */
  NUM_SUB = 202,

  /**
   * @generated from enum value: NUM_DIV = 203;
   */
  NUM_DIV = 203,

  /**
   * @generated from enum value: NUM_MUL = 204;
   */
  NUM_MUL = 204,

  /**
   * @generated from enum value: CMP_EQ = 301;
   */
  CMP_EQ = 301,

  /**
   * @generated from enum value: CMP_NOTEQ = 302;
   */
  CMP_NOTEQ = 302,

  /**
   * @generated from enum value: CMP_GT = 303;
   */
  CMP_GT = 303,

  /**
   * @generated from enum value: CMP_GTE = 304;
   */
  CMP_GTE = 304,

  /**
   * @generated from enum value: CMP_LT = 305;
   */
  CMP_LT = 305,

  /**
   * @generated from enum value: CMP_LTE = 306;
   */
  CMP_LTE = 306,

  /**
   * @generated from enum value: SET_IN = 401;
   */
  SET_IN = 401,

  /**
   * @generated from enum value: SET_NOTIN = 402;
   */
  SET_NOTIN = 402,
}
// Retrieve enum metadata with: proto3.getEnumType(BinaryOp_Operator)
proto3.util.setEnumType(BinaryOp_Operator, "types.v1.BinaryOp.Operator", [
  { no: 0, name: "INVALID" },
  { no: 101, name: "LOG_AND" },
  { no: 102, name: "LOG_OR" },
  { no: 201, name: "NUM_ADD" },
  { no: 202, name: "NUM_SUB" },
  { no: 203, name: "NUM_DIV" },
  { no: 204, name: "NUM_MUL" },
  { no: 301, name: "CMP_EQ" },
  { no: 302, name: "CMP_NOTEQ" },
  { no: 303, name: "CMP_GT" },
  { no: 304, name: "CMP_GTE" },
  { no: 305, name: "CMP_LT" },
  { no: 306, name: "CMP_LTE" },
  { no: 401, name: "SET_IN" },
  { no: 402, name: "SET_NOTIN" },
]);

/**
 * @generated from message types.v1.FuncCall
 */
export class FuncCall extends Message<FuncCall> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: repeated types.v1.Expr args = 2;
   */
  args: Expr[] = [];

  constructor(data?: PartialMessage<FuncCall>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.FuncCall";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "args", kind: "message", T: Expr, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FuncCall {
    return new FuncCall().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FuncCall {
    return new FuncCall().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FuncCall {
    return new FuncCall().fromJsonString(jsonString, options);
  }

  static equals(a: FuncCall | PlainMessage<FuncCall> | undefined, b: FuncCall | PlainMessage<FuncCall> | undefined): boolean {
    return proto3.util.equals(FuncCall, a, b);
  }
}

/**
 * @generated from message types.v1.Identifier
 */
export class Identifier extends Message<Identifier> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<Identifier>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Identifier";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Identifier {
    return new Identifier().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Identifier {
    return new Identifier().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Identifier {
    return new Identifier().fromJsonString(jsonString, options);
  }

  static equals(a: Identifier | PlainMessage<Identifier> | undefined, b: Identifier | PlainMessage<Identifier> | undefined): boolean {
    return proto3.util.equals(Identifier, a, b);
  }
}

/**
 * @generated from message types.v1.Selector
 */
export class Selector extends Message<Selector> {
  /**
   * @generated from field: types.v1.Expr x = 1;
   */
  x?: Expr;

  /**
   * @generated from field: types.v1.Identifier identifier = 2;
   */
  identifier?: Identifier;

  constructor(data?: PartialMessage<Selector>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Selector";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "x", kind: "message", T: Expr },
    { no: 2, name: "identifier", kind: "message", T: Identifier },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Selector {
    return new Selector().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Selector {
    return new Selector().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Selector {
    return new Selector().fromJsonString(jsonString, options);
  }

  static equals(a: Selector | PlainMessage<Selector> | undefined, b: Selector | PlainMessage<Selector> | undefined): boolean {
    return proto3.util.equals(Selector, a, b);
  }
}

/**
 * @generated from message types.v1.Indexor
 */
export class Indexor extends Message<Indexor> {
  /**
   * @generated from field: types.v1.Expr x = 1;
   */
  x?: Expr;

  /**
   * @generated from field: types.v1.Expr index = 2;
   */
  index?: Expr;

  constructor(data?: PartialMessage<Indexor>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Indexor";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "x", kind: "message", T: Expr },
    { no: 2, name: "index", kind: "message", T: Expr },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Indexor {
    return new Indexor().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Indexor {
    return new Indexor().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Indexor {
    return new Indexor().fromJsonString(jsonString, options);
  }

  static equals(a: Indexor | PlainMessage<Indexor> | undefined, b: Indexor | PlainMessage<Indexor> | undefined): boolean {
    return proto3.util.equals(Indexor, a, b);
  }
}

/**
 * @generated from message types.v1.Pipe
 */
export class Pipe extends Message<Pipe> {
  /**
   * @generated from field: types.v1.Expr head = 1;
   */
  head?: Expr;

  /**
   * @generated from field: types.v1.Expr tail = 2;
   */
  tail?: Expr;

  constructor(data?: PartialMessage<Pipe>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Pipe";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "head", kind: "message", T: Expr },
    { no: 2, name: "tail", kind: "message", T: Expr },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pipe {
    return new Pipe().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pipe {
    return new Pipe().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pipe {
    return new Pipe().fromJsonString(jsonString, options);
  }

  static equals(a: Pipe | PlainMessage<Pipe> | undefined, b: Pipe | PlainMessage<Pipe> | undefined): boolean {
    return proto3.util.equals(Pipe, a, b);
  }
}

