import { writable } from 'svelte/store';

export const bet_  = writable(0);
export const drawn_ = writable([]);
export const flop_ = writable([]);
export const turn_ = writable([]);
export const river_ = writable([]);
export const currentPhase_ = writable(-1);
export const handtype_ = writable("");
export const ranks_ = writable([]);
