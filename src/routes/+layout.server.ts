export const load = ({ locals }) => {

	if (locals.user) {
		return {
            tables: locals.tables,
			user: locals.user,
            usertables: locals.usertables
		};
	}

	return {
        tables: locals.tables,
		user: undefined,
        usertable: undefined
	};
};

