// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "auth/user/v1/service.proto" (package "auth.user.v1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { User } from "./user";
/**
 * @generated from protobuf message auth.user.v1.UserServiceCreateRequest
 */
export interface UserServiceCreateRequest {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: string email = 2;
     */
    email: string;
    /**
     * @generated from protobuf field: repeated uint64 roles = 3;
     */
    roles: bigint[];
}
/**
 * @generated from protobuf message auth.user.v1.UserServiceCreateResponse
 */
export interface UserServiceCreateResponse {
}
/**
 * @generated from protobuf message auth.user.v1.UserServiceUpdateRequest
 */
export interface UserServiceUpdateRequest {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: string name = 2;
     */
    name: string;
    /**
     * @generated from protobuf field: string email = 3;
     */
    email: string;
    /**
     * @generated from protobuf field: repeated uint64 roles = 4;
     */
    roles: bigint[];
}
/**
 * @generated from protobuf message auth.user.v1.UserServiceUpdateResponse
 */
export interface UserServiceUpdateResponse {
}
/**
 * @generated from protobuf message auth.user.v1.UserServiceByIDRequest
 */
export interface UserServiceByIDRequest {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
}
/**
 * @generated from protobuf message auth.user.v1.UserServiceByIDResponse
 */
export interface UserServiceByIDResponse {
    /**
     * @generated from protobuf field: auth.user.v1.User user = 1;
     */
    user?: User;
}
/**
 * @generated from protobuf message auth.user.v1.UserServiceAllRequest
 */
export interface UserServiceAllRequest {
}
/**
 * @generated from protobuf message auth.user.v1.UserServiceAllResponse
 */
export interface UserServiceAllResponse {
    /**
     * @generated from protobuf field: repeated auth.user.v1.User user = 1;
     */
    user: User[];
}
// @generated message type with reflection information, may provide speed optimized methods
class UserServiceCreateRequest$Type extends MessageType<UserServiceCreateRequest> {
    constructor() {
        super("auth.user.v1.UserServiceCreateRequest", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "email", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "roles", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<UserServiceCreateRequest>): UserServiceCreateRequest {
        const message = { name: "", email: "", roles: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UserServiceCreateRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserServiceCreateRequest): UserServiceCreateRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* string email */ 2:
                    message.email = reader.string();
                    break;
                case /* repeated uint64 roles */ 3:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.roles.push(reader.uint64().toBigInt());
                    else
                        message.roles.push(reader.uint64().toBigInt());
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UserServiceCreateRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* string email = 2; */
        if (message.email !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.email);
        /* repeated uint64 roles = 3; */
        if (message.roles.length) {
            writer.tag(3, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.roles.length; i++)
                writer.uint64(message.roles[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.user.v1.UserServiceCreateRequest
 */
export const UserServiceCreateRequest = new UserServiceCreateRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserServiceCreateResponse$Type extends MessageType<UserServiceCreateResponse> {
    constructor() {
        super("auth.user.v1.UserServiceCreateResponse", []);
    }
    create(value?: PartialMessage<UserServiceCreateResponse>): UserServiceCreateResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UserServiceCreateResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserServiceCreateResponse): UserServiceCreateResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: UserServiceCreateResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.user.v1.UserServiceCreateResponse
 */
export const UserServiceCreateResponse = new UserServiceCreateResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserServiceUpdateRequest$Type extends MessageType<UserServiceUpdateRequest> {
    constructor() {
        super("auth.user.v1.UserServiceUpdateRequest", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "email", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "roles", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<UserServiceUpdateRequest>): UserServiceUpdateRequest {
        const message = { id: "", name: "", email: "", roles: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UserServiceUpdateRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserServiceUpdateRequest): UserServiceUpdateRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* string name */ 2:
                    message.name = reader.string();
                    break;
                case /* string email */ 3:
                    message.email = reader.string();
                    break;
                case /* repeated uint64 roles */ 4:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.roles.push(reader.uint64().toBigInt());
                    else
                        message.roles.push(reader.uint64().toBigInt());
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UserServiceUpdateRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* string name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        /* string email = 3; */
        if (message.email !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.email);
        /* repeated uint64 roles = 4; */
        if (message.roles.length) {
            writer.tag(4, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.roles.length; i++)
                writer.uint64(message.roles[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.user.v1.UserServiceUpdateRequest
 */
export const UserServiceUpdateRequest = new UserServiceUpdateRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserServiceUpdateResponse$Type extends MessageType<UserServiceUpdateResponse> {
    constructor() {
        super("auth.user.v1.UserServiceUpdateResponse", []);
    }
    create(value?: PartialMessage<UserServiceUpdateResponse>): UserServiceUpdateResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UserServiceUpdateResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserServiceUpdateResponse): UserServiceUpdateResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: UserServiceUpdateResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.user.v1.UserServiceUpdateResponse
 */
export const UserServiceUpdateResponse = new UserServiceUpdateResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserServiceByIDRequest$Type extends MessageType<UserServiceByIDRequest> {
    constructor() {
        super("auth.user.v1.UserServiceByIDRequest", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<UserServiceByIDRequest>): UserServiceByIDRequest {
        const message = { id: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UserServiceByIDRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserServiceByIDRequest): UserServiceByIDRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UserServiceByIDRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.user.v1.UserServiceByIDRequest
 */
export const UserServiceByIDRequest = new UserServiceByIDRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserServiceByIDResponse$Type extends MessageType<UserServiceByIDResponse> {
    constructor() {
        super("auth.user.v1.UserServiceByIDResponse", [
            { no: 1, name: "user", kind: "message", T: () => User }
        ]);
    }
    create(value?: PartialMessage<UserServiceByIDResponse>): UserServiceByIDResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UserServiceByIDResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserServiceByIDResponse): UserServiceByIDResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* auth.user.v1.User user */ 1:
                    message.user = User.internalBinaryRead(reader, reader.uint32(), options, message.user);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UserServiceByIDResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* auth.user.v1.User user = 1; */
        if (message.user)
            User.internalBinaryWrite(message.user, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.user.v1.UserServiceByIDResponse
 */
export const UserServiceByIDResponse = new UserServiceByIDResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserServiceAllRequest$Type extends MessageType<UserServiceAllRequest> {
    constructor() {
        super("auth.user.v1.UserServiceAllRequest", []);
    }
    create(value?: PartialMessage<UserServiceAllRequest>): UserServiceAllRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UserServiceAllRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserServiceAllRequest): UserServiceAllRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: UserServiceAllRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.user.v1.UserServiceAllRequest
 */
export const UserServiceAllRequest = new UserServiceAllRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UserServiceAllResponse$Type extends MessageType<UserServiceAllResponse> {
    constructor() {
        super("auth.user.v1.UserServiceAllResponse", [
            { no: 1, name: "user", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => User }
        ]);
    }
    create(value?: PartialMessage<UserServiceAllResponse>): UserServiceAllResponse {
        const message = { user: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UserServiceAllResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UserServiceAllResponse): UserServiceAllResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated auth.user.v1.User user */ 1:
                    message.user.push(User.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UserServiceAllResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated auth.user.v1.User user = 1; */
        for (let i = 0; i < message.user.length; i++)
            User.internalBinaryWrite(message.user[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message auth.user.v1.UserServiceAllResponse
 */
export const UserServiceAllResponse = new UserServiceAllResponse$Type();
/**
 * @generated ServiceType for protobuf service auth.user.v1.UserService
 */
export const UserService = new ServiceType("auth.user.v1.UserService", [
    { name: "All", options: {}, I: UserServiceAllRequest, O: UserServiceAllResponse },
    { name: "ByID", options: {}, I: UserServiceByIDRequest, O: UserServiceByIDResponse },
    { name: "Create", options: {}, I: UserServiceCreateRequest, O: UserServiceCreateResponse },
    { name: "Update", options: {}, I: UserServiceUpdateRequest, O: UserServiceUpdateResponse }
]);
