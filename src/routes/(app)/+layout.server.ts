export const load = ({ locals }) => {
	let userToken = locals.userPb.authStore.token;
    
    if (locals.user) {
		return {
			user: locals.user
		};
	}

	return {
		user: undefined
	};
};
