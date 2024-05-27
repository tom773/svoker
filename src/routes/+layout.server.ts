import type {LayoutServerLoad} from './$types';

export type OutputType = { user: object; isLoggedIn: boolean};

export const load: LayoutServerLoad = async ({ locals }) => {
    
    const user = locals.user;
	if (user) {
		return {
			user: { user, isLoggedIn: true },
		    tables: locals.tables,
            games: locals.games,
        };
	}
	return {
		user: undefined,
        islLoggedIn: false,
	};
};

