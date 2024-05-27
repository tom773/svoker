import PocketBase from 'pocketbase';
import { serializeNonPOJOs } from '$lib/utils';
import { redirect } from '@sveltejs/kit';

export const handle = async ({ event, resolve }) => {
    // New PB instance
	event.locals.pb = new PocketBase('http://127.0.0.1:8090');
    // Get all tables
    event.locals.tables = await event.locals.pb.collection('tables').getFullList({
        sort: 'tnum',
    });
    // Get authStore from cookie
	event.locals.pb.authStore.loadFromCookie(event.request.headers.get('cookie') || '');

    try{
	    if (event.locals.pb.authStore.isValid) {
		    await event.locals.pb.collection('users').authRefresh()
            event.locals.user = serializeNonPOJOs(event.locals.pb.authStore.model);

        }
	} catch (err){
		event.locals.pb.authStore.clear();
	}

    if (event.url.pathname === '/holdem') {
        if (!event.locals.user) {
            throw redirect(303, '/signin');
        }
    }

    const isProd = process.env.NODE_ENV === 'production' ? true : false;
	
    const response = await resolve(event);
	
    response.headers.set('set-cookie', event.locals.pb.authStore.exportToCookie({ secure: isProd, sameSite: 'Lax', httpOnly: false }));
 
	return response;
};
