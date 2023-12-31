// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "common/events/v1/service.proto" (package "common.events.v1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Event } from "./event";
/**
 * Get single event by id.
 *
 * @generated from protobuf message common.events.v1.EventRequest
 */
export interface EventRequest {
    /**
     * @generated from protobuf field: string event_id = 1;
     */
    eventId: string;
}
/**
 * @generated from protobuf message common.events.v1.EventResponse
 */
export interface EventResponse {
    /**
     * @generated from protobuf field: common.events.v1.Event data = 1;
     */
    data?: Event;
}
/**
 * Get New Events since a given date.
 *
 * @generated from protobuf message common.events.v1.NewEventsRequest
 */
export interface NewEventsRequest {
    /**
     * @generated from protobuf field: uint64 since = 1;
     */
    since: bigint;
}
/**
 * @generated from protobuf message common.events.v1.NewEventsResponse
 */
export interface NewEventsResponse {
    /**
     * @generated from protobuf field: common.events.v1.Event data = 1;
     */
    data?: Event;
}
/**
 * Record User Events stream.
 *
 * @generated from protobuf message common.events.v1.RecordUserEventsRequest
 */
export interface RecordUserEventsRequest {
    /**
     * @generated from protobuf field: repeated common.events.v1.Event data = 1;
     */
    data: Event[];
}
/**
 * @generated from protobuf message common.events.v1.RecordUserEventsResponse
 */
export interface RecordUserEventsResponse {
    /**
     * @generated from protobuf field: bool ok = 1;
     */
    ok: boolean;
}
/**
 * Events stream.
 *
 * @generated from protobuf message common.events.v1.EventsRequest
 */
export interface EventsRequest {
    /**
     * @generated from protobuf field: common.events.v1.Event data = 1;
     */
    data?: Event;
}
/**
 * @generated from protobuf message common.events.v1.EventsResponse
 */
export interface EventsResponse {
    /**
     * @generated from protobuf field: common.events.v1.Event data = 1;
     */
    data?: Event;
}
// @generated message type with reflection information, may provide speed optimized methods
class EventRequest$Type extends MessageType<EventRequest> {
    constructor() {
        super("common.events.v1.EventRequest", [
            { no: 1, name: "event_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<EventRequest>): EventRequest {
        const message = { eventId: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<EventRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EventRequest): EventRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string event_id */ 1:
                    message.eventId = reader.string();
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
    internalBinaryWrite(message: EventRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string event_id = 1; */
        if (message.eventId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.eventId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message common.events.v1.EventRequest
 */
export const EventRequest = new EventRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EventResponse$Type extends MessageType<EventResponse> {
    constructor() {
        super("common.events.v1.EventResponse", [
            { no: 1, name: "data", kind: "message", T: () => Event }
        ]);
    }
    create(value?: PartialMessage<EventResponse>): EventResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<EventResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EventResponse): EventResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* common.events.v1.Event data */ 1:
                    message.data = Event.internalBinaryRead(reader, reader.uint32(), options, message.data);
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
    internalBinaryWrite(message: EventResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* common.events.v1.Event data = 1; */
        if (message.data)
            Event.internalBinaryWrite(message.data, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message common.events.v1.EventResponse
 */
export const EventResponse = new EventResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NewEventsRequest$Type extends MessageType<NewEventsRequest> {
    constructor() {
        super("common.events.v1.NewEventsRequest", [
            { no: 1, name: "since", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<NewEventsRequest>): NewEventsRequest {
        const message = { since: 0n };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<NewEventsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: NewEventsRequest): NewEventsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* uint64 since */ 1:
                    message.since = reader.uint64().toBigInt();
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
    internalBinaryWrite(message: NewEventsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* uint64 since = 1; */
        if (message.since !== 0n)
            writer.tag(1, WireType.Varint).uint64(message.since);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message common.events.v1.NewEventsRequest
 */
export const NewEventsRequest = new NewEventsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NewEventsResponse$Type extends MessageType<NewEventsResponse> {
    constructor() {
        super("common.events.v1.NewEventsResponse", [
            { no: 1, name: "data", kind: "message", T: () => Event }
        ]);
    }
    create(value?: PartialMessage<NewEventsResponse>): NewEventsResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<NewEventsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: NewEventsResponse): NewEventsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* common.events.v1.Event data */ 1:
                    message.data = Event.internalBinaryRead(reader, reader.uint32(), options, message.data);
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
    internalBinaryWrite(message: NewEventsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* common.events.v1.Event data = 1; */
        if (message.data)
            Event.internalBinaryWrite(message.data, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message common.events.v1.NewEventsResponse
 */
export const NewEventsResponse = new NewEventsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RecordUserEventsRequest$Type extends MessageType<RecordUserEventsRequest> {
    constructor() {
        super("common.events.v1.RecordUserEventsRequest", [
            { no: 1, name: "data", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Event }
        ]);
    }
    create(value?: PartialMessage<RecordUserEventsRequest>): RecordUserEventsRequest {
        const message = { data: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<RecordUserEventsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RecordUserEventsRequest): RecordUserEventsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated common.events.v1.Event data */ 1:
                    message.data.push(Event.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: RecordUserEventsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated common.events.v1.Event data = 1; */
        for (let i = 0; i < message.data.length; i++)
            Event.internalBinaryWrite(message.data[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message common.events.v1.RecordUserEventsRequest
 */
export const RecordUserEventsRequest = new RecordUserEventsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RecordUserEventsResponse$Type extends MessageType<RecordUserEventsResponse> {
    constructor() {
        super("common.events.v1.RecordUserEventsResponse", [
            { no: 1, name: "ok", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<RecordUserEventsResponse>): RecordUserEventsResponse {
        const message = { ok: false };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<RecordUserEventsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RecordUserEventsResponse): RecordUserEventsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool ok */ 1:
                    message.ok = reader.bool();
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
    internalBinaryWrite(message: RecordUserEventsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool ok = 1; */
        if (message.ok !== false)
            writer.tag(1, WireType.Varint).bool(message.ok);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message common.events.v1.RecordUserEventsResponse
 */
export const RecordUserEventsResponse = new RecordUserEventsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EventsRequest$Type extends MessageType<EventsRequest> {
    constructor() {
        super("common.events.v1.EventsRequest", [
            { no: 1, name: "data", kind: "message", T: () => Event }
        ]);
    }
    create(value?: PartialMessage<EventsRequest>): EventsRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<EventsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EventsRequest): EventsRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* common.events.v1.Event data */ 1:
                    message.data = Event.internalBinaryRead(reader, reader.uint32(), options, message.data);
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
    internalBinaryWrite(message: EventsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* common.events.v1.Event data = 1; */
        if (message.data)
            Event.internalBinaryWrite(message.data, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message common.events.v1.EventsRequest
 */
export const EventsRequest = new EventsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EventsResponse$Type extends MessageType<EventsResponse> {
    constructor() {
        super("common.events.v1.EventsResponse", [
            { no: 1, name: "data", kind: "message", T: () => Event }
        ]);
    }
    create(value?: PartialMessage<EventsResponse>): EventsResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<EventsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EventsResponse): EventsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* common.events.v1.Event data */ 1:
                    message.data = Event.internalBinaryRead(reader, reader.uint32(), options, message.data);
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
    internalBinaryWrite(message: EventsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* common.events.v1.Event data = 1; */
        if (message.data)
            Event.internalBinaryWrite(message.data, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message common.events.v1.EventsResponse
 */
export const EventsResponse = new EventsResponse$Type();
/**
 * @generated ServiceType for protobuf service common.events.v1.EventService
 */
export const EventService = new ServiceType("common.events.v1.EventService", [
    { name: "Event", options: {}, I: EventRequest, O: EventResponse },
    { name: "NewEvents", serverStreaming: true, options: {}, I: NewEventsRequest, O: NewEventsResponse },
    { name: "RecordUserEvents", clientStreaming: true, options: {}, I: RecordUserEventsRequest, O: RecordUserEventsResponse },
    { name: "Events", serverStreaming: true, clientStreaming: true, options: {}, I: EventsRequest, O: EventsResponse }
]);
