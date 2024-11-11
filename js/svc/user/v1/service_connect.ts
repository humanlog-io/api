// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/user/v1/service.proto (package svc.user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateOrganizationRequest, CreateOrganizationResponse, GetLogoutURLRequest, GetLogoutURLResponse, ListOrganizationRequest, ListOrganizationResponse, WhoamiRequest, WhoamiResponse } from "./service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service svc.user.v1.UserService
 */
export const UserService = {
  typeName: "svc.user.v1.UserService",
  methods: {
    /**
     * @generated from rpc svc.user.v1.UserService.Whoami
     */
    whoami: {
      name: "Whoami",
      I: WhoamiRequest,
      O: WhoamiResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.user.v1.UserService.GetLogoutURL
     */
    getLogoutURL: {
      name: "GetLogoutURL",
      I: GetLogoutURLRequest,
      O: GetLogoutURLResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.user.v1.UserService.CreateOrganization
     */
    createOrganization: {
      name: "CreateOrganization",
      I: CreateOrganizationRequest,
      O: CreateOrganizationResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.user.v1.UserService.ListOrganization
     */
    listOrganization: {
      name: "ListOrganization",
      I: ListOrganizationRequest,
      O: ListOrganizationResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

