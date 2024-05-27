// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
        type PocketBase = import('pocketbase').default;
        interface Locals {
            adminPb?: PocketBase;
            userPb?: PocketBase;
            user?: Record<string, T>;
            tables?: Record<string, T>;
            games?: Record<string, T>;
		}
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
