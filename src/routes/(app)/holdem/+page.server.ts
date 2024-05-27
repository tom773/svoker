import { error, redirect } from '@sveltejs/kit';
import { json } from '@sveltejs/kit';
import { serializeNonPOJOs } from '$lib/utils';
import type { PageServerLoad } from './$types';

export const actions = {
    addToTable: async ({request, locals}) => {
        const body = await request.formData();
        try {
            const tables = await locals.userPb.collection('v2tables').update(body.get('table'), {
                "players+": locals.user.id,
                "currentplayers+": 1,
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
            const tabup = await locals.pb.collection('v2tables').update(body.get('table_'), {
                "players-": locals.user.id,
                "currentplayers-": 1,
            });
            } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        return redirect(303, 'holdem');
    }
};
