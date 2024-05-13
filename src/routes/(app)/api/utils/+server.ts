import PocketBase from 'pocketbase';
import { json } from '@sveltejs/kit'
import { serializeNonPOJOs } from '$lib/utils';
const pb = new PocketBase("http://localhost:8090");

export async function POST({request}) {
    const body = await request.json();
    const playersobjs = [];
    for await (let player of body.players) {
        const record = await pb.collection('users').getOne(player);
        if (!record) {
            return json({error: 'User not found'}, {status: 404});
        } else {
            playersobjs.push(record);
        }
    }
    const tablesobj = serializeNonPOJOs(playersobjs);
    return json(tablesobj); 
}
