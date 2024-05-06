import * as wasm from "./pokerutil_bg.wasm";
import { __wbg_set_wasm } from "./pokerutil_bg.js";
__wbg_set_wasm(wasm);
export * from "./pokerutil_bg.js";
