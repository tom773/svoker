import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals, params }) => {
	if (!locals.userPb.authStore.isValid) {
	    console.log('Not authenticated');	
	}
   let tableid = params.tableid;
    let response = await fetch("http://localhost:8090/api/table/avatar", {
        method: 'POST',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({tableid: tableid}),
    })
    return { tablePlayers: await response.json() }; 
};
