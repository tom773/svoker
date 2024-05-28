import { redirect } from '@sveltejs/kit';

export const load = ({ locals }) => {
	if (!locals.userPb.authStore.isValid) {
		throw redirect(303, '/signin');
	}
};
