import { error, redirect } from '@sveltejs/kit';
import { _players } from '$lib/stores/table';

export const actions = {
    addToTable: async ({request, locals}) => {
        const body = await request.formData();
        try {
            const tables = await locals.pb.collection('tables').update(body.get('table'), {
                "players+": locals.user.id,
                "currentplayers+": 1,
            });
            locals.tables = tables;
            _players.update((players) => {
                players.push(locals.user.id);
                console.log(_players); 
                return players;
            });
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        throw redirect(303, 'holdem/'+body.get('tnum'));
    },
    removeFromTable: async ({request, locals}) => {
        const body = await request.formData();
        try {
            const tabup = await locals.pb.collection('tables').update(body.get('table_'), {
                "players-": locals.user.id,
                "currentplayers-": 1,
            });
            locals.tables = tabup;
            _players.update((players) => {
                players.splice(players.indexOf(locals.user.id), 1);
                return players;
            });
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        return redirect(303, 'holdem');
    }

};
