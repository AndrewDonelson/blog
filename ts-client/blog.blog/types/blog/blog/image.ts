/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "blog.blog";

export interface Image {
  creator: string;
  id: number;
  ipfsurl: string;
  tags: string;
  views: string;
  likes: string;
  dislikes: string;
}

function createBaseImage(): Image {
  return { creator: "", id: 0, ipfsurl: "", tags: "", views: "", likes: "", dislikes: "" };
}

export const Image = {
  encode(message: Image, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.ipfsurl !== "") {
      writer.uint32(26).string(message.ipfsurl);
    }
    if (message.tags !== "") {
      writer.uint32(34).string(message.tags);
    }
    if (message.views !== "") {
      writer.uint32(42).string(message.views);
    }
    if (message.likes !== "") {
      writer.uint32(50).string(message.likes);
    }
    if (message.dislikes !== "") {
      writer.uint32(58).string(message.dislikes);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Image {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseImage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.ipfsurl = reader.string();
          break;
        case 4:
          message.tags = reader.string();
          break;
        case 5:
          message.views = reader.string();
          break;
        case 6:
          message.likes = reader.string();
          break;
        case 7:
          message.dislikes = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Image {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      ipfsurl: isSet(object.ipfsurl) ? String(object.ipfsurl) : "",
      tags: isSet(object.tags) ? String(object.tags) : "",
      views: isSet(object.views) ? String(object.views) : "",
      likes: isSet(object.likes) ? String(object.likes) : "",
      dislikes: isSet(object.dislikes) ? String(object.dislikes) : "",
    };
  },

  toJSON(message: Image): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.ipfsurl !== undefined && (obj.ipfsurl = message.ipfsurl);
    message.tags !== undefined && (obj.tags = message.tags);
    message.views !== undefined && (obj.views = message.views);
    message.likes !== undefined && (obj.likes = message.likes);
    message.dislikes !== undefined && (obj.dislikes = message.dislikes);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Image>, I>>(object: I): Image {
    const message = createBaseImage();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.ipfsurl = object.ipfsurl ?? "";
    message.tags = object.tags ?? "";
    message.views = object.views ?? "";
    message.likes = object.likes ?? "";
    message.dislikes = object.dislikes ?? "";
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}