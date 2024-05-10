import PocketBase from 'pocketbase';
import { serializeNonPOJOs } from '$lib/utils';
import { redirect } from '@sveltejs/kit';

export const handle = async ({ event, resolve }) => {
	event.locals.pb = new PocketBase('http://localhost:8090');
	event.locals.pb.authStore.loadFromCookie(event.request.headers.get('cookie') || '');

	if (event.locals.pb.authStore.isValid) {
		event.locals.user = serializeNonPOJOs(event.locals.pb.authStore.model);
	} else {
		event.locals.user = undefined;
	}
    
    if (event.url.pathname === '/holdem') {
        if (!event.locals.user) {
            throw redirect(303, '/signin');
        }
    }
    event.locals.tables = await event.locals.pb.collection('tables').getFullList({
        sort: 'tnum',
    });
	const response = await resolve(event);

	response.headers.set('set-cookie', event.locals.pb.authStore.exportToCookie({ secure: false }));
 
	return response;
};
