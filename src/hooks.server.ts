//import { serializeNonPOJOs } from '$lib/utils';
//import { redirect } from '@sveltejs/kit';
import { newInstance } from '$lib/pocketbase';
import {IDENTITY, PASSWORD} from '$env/static/private';

export const handle = async ({ event, resolve }) => {
    const adminPb = newInstance();
    const userPb = newInstance();
    
    //sign in
    await adminPb.admins.authWithPassword(IDENTITY, PASSWORD);
    event.locals.adminPb = adminPb;
    event.locals.userPb = userPb;
    event.locals.tables = await adminPb.collection('v2tables').getFullList();
    event.locals.games = await adminPb.collection('v2game').getFullList();
    // Load the authStore from the cookie
	event.locals.userPb.authStore.loadFromCookie(event.request.headers.get('cookie') || '');
    try{
        //refresh the auth if it is valid
        if(!event.locals.userPb.authStore.isValid) await event.locals.userPb.authStore.authRefresh();
        //spread the model to locals.user to be available in all pages
        event.locals.user = {...event.locals.userPb.authStore.model};
    } catch (err) {
        console.log('error in hooks')
		event.locals.userPb.authStore.clear();
	}
    const response = await resolve(event);
    // Set the cookie
	response.headers.append('set-cookie', event.locals.userPb.authStore.exportToCookie({httpOnly: false}));
    return response;
};
