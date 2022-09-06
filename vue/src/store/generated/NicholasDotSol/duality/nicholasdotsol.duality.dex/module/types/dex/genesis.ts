/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { TickMap } from "../dex/tick_map";
import { PairMap } from "../dex/pair_map";
import { Tokens } from "../dex/tokens";
import { TokenMap } from "../dex/token_map";

export const protobufPackage = "nicholasdotsol.duality.dex";

/** GenesisState defines the dex module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  tickMapList: TickMap[];
  pairMapList: PairMap[];
  tokensList: Tokens[];
  tokensCount: number;
  /** this line is used by starport scaffolding # genesis/proto/state */
  tokenMapList: TokenMap[];
}

const baseGenesisState: object = { tokensCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.tickMapList) {
      TickMap.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.pairMapList) {
      PairMap.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.tokensList) {
      Tokens.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.tokensCount !== 0) {
      writer.uint32(40).uint64(message.tokensCount);
    }
    for (const v of message.tokenMapList) {
      TokenMap.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.tickMapList = [];
    message.pairMapList = [];
    message.tokensList = [];
    message.tokenMapList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.tickMapList.push(TickMap.decode(reader, reader.uint32()));
          break;
        case 3:
          message.pairMapList.push(PairMap.decode(reader, reader.uint32()));
          break;
        case 4:
          message.tokensList.push(Tokens.decode(reader, reader.uint32()));
          break;
        case 5:
          message.tokensCount = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.tokenMapList.push(TokenMap.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.tickMapList = [];
    message.pairMapList = [];
    message.tokensList = [];
    message.tokenMapList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.tickMapList !== undefined && object.tickMapList !== null) {
      for (const e of object.tickMapList) {
        message.tickMapList.push(TickMap.fromJSON(e));
      }
    }
    if (object.pairMapList !== undefined && object.pairMapList !== null) {
      for (const e of object.pairMapList) {
        message.pairMapList.push(PairMap.fromJSON(e));
      }
    }
    if (object.tokensList !== undefined && object.tokensList !== null) {
      for (const e of object.tokensList) {
        message.tokensList.push(Tokens.fromJSON(e));
      }
    }
    if (object.tokensCount !== undefined && object.tokensCount !== null) {
      message.tokensCount = Number(object.tokensCount);
    } else {
      message.tokensCount = 0;
    }
    if (object.tokenMapList !== undefined && object.tokenMapList !== null) {
      for (const e of object.tokenMapList) {
        message.tokenMapList.push(TokenMap.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.tickMapList) {
      obj.tickMapList = message.tickMapList.map((e) =>
        e ? TickMap.toJSON(e) : undefined
      );
    } else {
      obj.tickMapList = [];
    }
    if (message.pairMapList) {
      obj.pairMapList = message.pairMapList.map((e) =>
        e ? PairMap.toJSON(e) : undefined
      );
    } else {
      obj.pairMapList = [];
    }
    if (message.tokensList) {
      obj.tokensList = message.tokensList.map((e) =>
        e ? Tokens.toJSON(e) : undefined
      );
    } else {
      obj.tokensList = [];
    }
    message.tokensCount !== undefined &&
      (obj.tokensCount = message.tokensCount);
    if (message.tokenMapList) {
      obj.tokenMapList = message.tokenMapList.map((e) =>
        e ? TokenMap.toJSON(e) : undefined
      );
    } else {
      obj.tokenMapList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.tickMapList = [];
    message.pairMapList = [];
    message.tokensList = [];
    message.tokenMapList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.tickMapList !== undefined && object.tickMapList !== null) {
      for (const e of object.tickMapList) {
        message.tickMapList.push(TickMap.fromPartial(e));
      }
    }
    if (object.pairMapList !== undefined && object.pairMapList !== null) {
      for (const e of object.pairMapList) {
        message.pairMapList.push(PairMap.fromPartial(e));
      }
    }
    if (object.tokensList !== undefined && object.tokensList !== null) {
      for (const e of object.tokensList) {
        message.tokensList.push(Tokens.fromPartial(e));
      }
    }
    if (object.tokensCount !== undefined && object.tokensCount !== null) {
      message.tokensCount = object.tokensCount;
    } else {
      message.tokensCount = 0;
    }
    if (object.tokenMapList !== undefined && object.tokenMapList !== null) {
      for (const e of object.tokenMapList) {
        message.tokenMapList.push(TokenMap.fromPartial(e));
      }
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
