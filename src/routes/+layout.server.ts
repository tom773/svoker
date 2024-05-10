export const load = ({ locals }) => {

	if (locals.user) {
		return {
            tables: locals.tables,
			user: locals.user
		};
	}

	return {
        tables: locals.tables,
		user: undefined
	};
};

