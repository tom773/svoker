import type {LayoutServerLoad} from './$types';

export type OutputType = { user: object; isLoggedIn: boolean};

export const load: LayoutServerLoad = async ({ locals }) => {
    
    const user = locals.user;
    console.log('user: ', user);
	if (user) {
		return {
			user: { user, isLoggedIn: true },
		};
	}
	return {
		user: undefined,
        islLoggedIn: false,
	};
};

