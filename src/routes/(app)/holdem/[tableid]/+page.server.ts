export const load = ({ locals }) => {
	if (!locals.userPb.authStore.isValid) {
	    console.log('Not authenticated');	
	}
};
