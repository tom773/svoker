// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
        type PocketBase = import('pocketbase').default;
        interface Locals {
            pb?: PocketBase;
            user?: Record<string, T>;
		}
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
