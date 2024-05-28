import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
    let tableid = params.tableid;
    let response = await fetch("http://localhost:8090/api/table/avatar", {
        method: 'POST',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({tableid: tableid}),
    })
    return { tablePlayers: await response.json() };
};
