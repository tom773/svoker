import { error, redirect } from '@sveltejs/kit';
import { json } from '@sveltejs/kit';
import { serializeNonPOJOs } from '$lib/utils';

export const actions = {
    addToTable: async ({request, locals}) => {
        const body = await request.formData();
        try {
            const tables = await locals.pb.collection('tables').update(body.get('table'), {
                "players+": locals.user.id,
                "currentplayers+": 1,
            });
            const gametable = await locals.pb.collection('gametable').create({
                "user": locals.user.id,
                "table": body.get('table'),
                "cards": json([]),
            });
            locals.tables = tables;
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        throw redirect(303, 'holdem/'+body.get('table'));
    },
    removeFromTable: async ({request, locals}) => {
        const body = await request.formData();
        try {
            const tabup = await locals.pb.collection('tables').update(body.get('table_'), {
                "players-": locals.user.id,
                "currentplayers-": 1,
            });
            const gametables = await locals.pb.collection('gametable').getList(1, 50,{
                "filter": `user = "${locals.user.id}" && table = "${body.get('table_')}"`,
            });
            let gtobj = serializeNonPOJOs(gametables);
            console.log(gtobj.items[0].id);
            if (gtobj.items.length != 0) {
                const updated_gametables = await locals.pb.collection('gametable').delete(gtobj.items[0].id);
            }
            locals.tables = tabup;
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        return redirect(303, 'holdem');
    }

};
