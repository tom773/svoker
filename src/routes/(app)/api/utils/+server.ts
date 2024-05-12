import PocketBase from 'pocketbase';
import { json } from '@sveltejs/kit'
import { serializeNonPOJOs } from '$lib/utils';
const pb = new PocketBase("http://localhost:8090");

export async function GET() {
    const tables = await pb.collection('tables').getFullList();
    const tablesobj = serializeNonPOJOs(tables);
    return json(tablesobj); 
}
