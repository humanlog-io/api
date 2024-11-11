// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/ingest/v1/service.proto (package svc.ingest.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetHeartbeatRequest, GetHeartbeatResponse, IngestBidiStreamRequest, IngestBidiStreamResponse, IngestRequest, IngestResponse, IngestStreamRequest, IngestStreamResponse } from "./service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service svc.ingest.v1.IngestService
 */
export const IngestService = {
  typeName: "svc.ingest.v1.IngestService",
  methods: {
    /**
     * @generated from rpc svc.ingest.v1.IngestService.GetHeartbeat
     */
    getHeartbeat: {
      name: "GetHeartbeat",
      I: GetHeartbeatRequest,
      O: GetHeartbeatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.ingest.v1.IngestService.Ingest
     */
    ingest: {
      name: "Ingest",
      I: IngestRequest,
      O: IngestResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.ingest.v1.IngestService.IngestStream
     */
    ingestStream: {
      name: "IngestStream",
      I: IngestStreamRequest,
      O: IngestStreamResponse,
      kind: MethodKind.ClientStreaming,
    },
    /**
     * @generated from rpc svc.ingest.v1.IngestService.IngestBidiStream
     */
    ingestBidiStream: {
      name: "IngestBidiStream",
      I: IngestBidiStreamRequest,
      O: IngestBidiStreamResponse,
      kind: MethodKind.BiDiStreaming,
    },
  }
} as const;

