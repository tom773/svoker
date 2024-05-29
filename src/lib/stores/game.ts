import { writable } from 'svelte/store';

export const hand = writable();
export const flop = writable([]);
export const turn = writable([]);
export const river = writable([]);
