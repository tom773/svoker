import { redirect } from '@sveltejs/kit';

export const POST = ({ locals }) => {
	locals.userPb.authStore.clear();
	locals.user = undefined;

	throw redirect(303, '/');
};
