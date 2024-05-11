import { error, redirect } from '@sveltejs/kit';

export const actions = {
    addToTable: async ({request, locals}) => {
        const body = await request.formData();
        try {
            const { table } = await locals.pb.collection('tables').update(body.get('table'), {
                players: locals.user.id,
            });
            if (table){
                locals.tables.push(table); 
            }
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        throw redirect(303, 'holdem/'+body.get('tnum'));
        
    },

};
