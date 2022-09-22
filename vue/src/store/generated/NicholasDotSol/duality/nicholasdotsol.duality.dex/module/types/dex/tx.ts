/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "nicholasdotsol.duality.dex";

export interface MsgDeposit {
  creator: string;
  tokenA: string;
  tokenB: string;
  amountA: string;
  amountB: string;
  priceIndex: number;
  feeIndex: number;
}

export interface MsgDepositResponse {}

export interface MsgWithdrawal {
  creator: string;
  tokenA: string;
  tokenB: string;
  sharesToRemove: string;
  priceIndex: number;
  feeIndex: number;
  receiver: string;
}

export interface MsgWithdrawalResponse {}

export interface MsgSwap {
  creator: string;
  tokenA: string;
  tokenB: string;
  amountIn: string;
  tokenIn: string;
  minOut: string;
}

export interface MsgSwapResponse {}

const baseMsgDeposit: object = {
  creator: "",
  tokenA: "",
  tokenB: "",
  amountA: "",
  amountB: "",
  priceIndex: 0,
  feeIndex: 0,
};

export const MsgDeposit = {
  encode(message: MsgDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenA !== "") {
      writer.uint32(18).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(26).string(message.tokenB);
    }
    if (message.amountA !== "") {
      writer.uint32(34).string(message.amountA);
    }
    if (message.amountB !== "") {
      writer.uint32(42).string(message.amountB);
    }
    if (message.priceIndex !== 0) {
      writer.uint32(48).int64(message.priceIndex);
    }
    if (message.feeIndex !== 0) {
      writer.uint32(56).uint64(message.feeIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeposit } as MsgDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenA = reader.string();
          break;
        case 3:
          message.tokenB = reader.string();
          break;
        case 4:
          message.amountA = reader.string();
          break;
        case 5:
          message.amountB = reader.string();
          break;
        case 6:
          message.priceIndex = longToNumber(reader.int64() as Long);
          break;
        case 7:
          message.feeIndex = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeposit {
    const message = { ...baseMsgDeposit } as MsgDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = String(object.tokenA);
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = String(object.tokenB);
    } else {
      message.tokenB = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = String(object.amountA);
    } else {
      message.amountA = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = String(object.amountB);
    } else {
      message.amountB = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = Number(object.priceIndex);
    } else {
      message.priceIndex = 0;
    }
    if (object.feeIndex !== undefined && object.feeIndex !== null) {
      message.feeIndex = Number(object.feeIndex);
    } else {
      message.feeIndex = 0;
    }
    return message;
  },

  toJSON(message: MsgDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.amountA !== undefined && (obj.amountA = message.amountA);
    message.amountB !== undefined && (obj.amountB = message.amountB);
    message.priceIndex !== undefined && (obj.priceIndex = message.priceIndex);
    message.feeIndex !== undefined && (obj.feeIndex = message.feeIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeposit>): MsgDeposit {
    const message = { ...baseMsgDeposit } as MsgDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = object.tokenA;
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = object.tokenB;
    } else {
      message.tokenB = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = object.amountA;
    } else {
      message.amountA = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = object.amountB;
    } else {
      message.amountB = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = object.priceIndex;
    } else {
      message.priceIndex = 0;
    }
    if (object.feeIndex !== undefined && object.feeIndex !== null) {
      message.feeIndex = object.feeIndex;
    } else {
      message.feeIndex = 0;
    }
    return message;
  },
};

const baseMsgDepositResponse: object = {};

export const MsgDepositResponse = {
  encode(_: MsgDepositResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    return message;
  },

  toJSON(_: MsgDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDepositResponse>): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    return message;
  },
};

const baseMsgWithdrawal: object = {
  creator: "",
  tokenA: "",
  tokenB: "",
  sharesToRemove: "",
  priceIndex: 0,
  feeIndex: 0,
  receiver: "",
};

export const MsgWithdrawal = {
  encode(message: MsgWithdrawal, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenA !== "") {
      writer.uint32(18).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(26).string(message.tokenB);
    }
    if (message.sharesToRemove !== "") {
      writer.uint32(34).string(message.sharesToRemove);
    }
    if (message.priceIndex !== 0) {
      writer.uint32(40).int64(message.priceIndex);
    }
    if (message.feeIndex !== 0) {
      writer.uint32(48).uint64(message.feeIndex);
    }
    if (message.receiver !== "") {
      writer.uint32(58).string(message.receiver);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawal } as MsgWithdrawal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenA = reader.string();
          break;
        case 3:
          message.tokenB = reader.string();
          break;
        case 4:
          message.sharesToRemove = reader.string();
          break;
        case 5:
          message.priceIndex = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.feeIndex = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.receiver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawal {
    const message = { ...baseMsgWithdrawal } as MsgWithdrawal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = String(object.tokenA);
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = String(object.tokenB);
    } else {
      message.tokenB = "";
    }
    if (object.sharesToRemove !== undefined && object.sharesToRemove !== null) {
      message.sharesToRemove = String(object.sharesToRemove);
    } else {
      message.sharesToRemove = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = Number(object.priceIndex);
    } else {
      message.priceIndex = 0;
    }
    if (object.feeIndex !== undefined && object.feeIndex !== null) {
      message.feeIndex = Number(object.feeIndex);
    } else {
      message.feeIndex = 0;
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver);
    } else {
      message.receiver = "";
    }
    return message;
  },

  toJSON(message: MsgWithdrawal): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.sharesToRemove !== undefined &&
      (obj.sharesToRemove = message.sharesToRemove);
    message.priceIndex !== undefined && (obj.priceIndex = message.priceIndex);
    message.feeIndex !== undefined && (obj.feeIndex = message.feeIndex);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWithdrawal>): MsgWithdrawal {
    const message = { ...baseMsgWithdrawal } as MsgWithdrawal;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = object.tokenA;
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = object.tokenB;
    } else {
      message.tokenB = "";
    }
    if (object.sharesToRemove !== undefined && object.sharesToRemove !== null) {
      message.sharesToRemove = object.sharesToRemove;
    } else {
      message.sharesToRemove = "";
    }
    if (object.priceIndex !== undefined && object.priceIndex !== null) {
      message.priceIndex = object.priceIndex;
    } else {
      message.priceIndex = 0;
    }
    if (object.feeIndex !== undefined && object.feeIndex !== null) {
      message.feeIndex = object.feeIndex;
    } else {
      message.feeIndex = 0;
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver;
    } else {
      message.receiver = "";
    }
    return message;
  },
};

const baseMsgWithdrawalResponse: object = {};

export const MsgWithdrawalResponse = {
  encode(_: MsgWithdrawalResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawalResponse } as MsgWithdrawalResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgWithdrawalResponse {
    const message = { ...baseMsgWithdrawalResponse } as MsgWithdrawalResponse;
    return message;
  },

  toJSON(_: MsgWithdrawalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgWithdrawalResponse>): MsgWithdrawalResponse {
    const message = { ...baseMsgWithdrawalResponse } as MsgWithdrawalResponse;
    return message;
  },
};

const baseMsgSwap: object = {
  creator: "",
  tokenA: "",
  tokenB: "",
  amountIn: "",
  tokenIn: "",
  minOut: "",
};

export const MsgSwap = {
  encode(message: MsgSwap, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenA !== "") {
      writer.uint32(18).string(message.tokenA);
    }
    if (message.tokenB !== "") {
      writer.uint32(26).string(message.tokenB);
    }
    if (message.amountIn !== "") {
      writer.uint32(34).string(message.amountIn);
    }
    if (message.tokenIn !== "") {
      writer.uint32(42).string(message.tokenIn);
    }
    if (message.minOut !== "") {
      writer.uint32(50).string(message.minOut);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwap {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwap } as MsgSwap;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenA = reader.string();
          break;
        case 3:
          message.tokenB = reader.string();
          break;
        case 4:
          message.amountIn = reader.string();
          break;
        case 5:
          message.tokenIn = reader.string();
          break;
        case 6:
          message.minOut = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = String(object.tokenA);
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = String(object.tokenB);
    } else {
      message.tokenB = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = String(object.amountIn);
    } else {
      message.amountIn = "";
    }
    if (object.tokenIn !== undefined && object.tokenIn !== null) {
      message.tokenIn = String(object.tokenIn);
    } else {
      message.tokenIn = "";
    }
    if (object.minOut !== undefined && object.minOut !== null) {
      message.minOut = String(object.minOut);
    } else {
      message.minOut = "";
    }
    return message;
  },

  toJSON(message: MsgSwap): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB);
    message.amountIn !== undefined && (obj.amountIn = message.amountIn);
    message.tokenIn !== undefined && (obj.tokenIn = message.tokenIn);
    message.minOut !== undefined && (obj.minOut = message.minOut);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwap>): MsgSwap {
    const message = { ...baseMsgSwap } as MsgSwap;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = object.tokenA;
    } else {
      message.tokenA = "";
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = object.tokenB;
    } else {
      message.tokenB = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = object.amountIn;
    } else {
      message.amountIn = "";
    }
    if (object.tokenIn !== undefined && object.tokenIn !== null) {
      message.tokenIn = object.tokenIn;
    } else {
      message.tokenIn = "";
    }
    if (object.minOut !== undefined && object.minOut !== null) {
      message.minOut = object.minOut;
    } else {
      message.minOut = "";
    }
    return message;
  },
};

const baseMsgSwapResponse: object = {};

export const MsgSwapResponse = {
  encode(_: MsgSwapResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwapResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    return message;
  },

  toJSON(_: MsgSwapResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSwapResponse>): MsgSwapResponse {
    const message = { ...baseMsgSwapResponse } as MsgSwapResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse>;
  Withdrawal(request: MsgWithdrawal): Promise<MsgWithdrawalResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Swap(request: MsgSwap): Promise<MsgSwapResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse> {
    const data = MsgDeposit.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "Deposit",
      data
    );
    return promise.then((data) => MsgDepositResponse.decode(new Reader(data)));
  }

  Withdrawal(request: MsgWithdrawal): Promise<MsgWithdrawalResponse> {
    const data = MsgWithdrawal.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "Withdrawal",
      data
    );
    return promise.then((data) =>
      MsgWithdrawalResponse.decode(new Reader(data))
    );
  }

  Swap(request: MsgSwap): Promise<MsgSwapResponse> {
    const data = MsgSwap.encode(request).finish();
    const promise = this.rpc.request(
      "nicholasdotsol.duality.dex.Msg",
      "Swap",
      data
    );
    return promise.then((data) => MsgSwapResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
