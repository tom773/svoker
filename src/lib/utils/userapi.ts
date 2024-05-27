import PocketBase from 'pocketbase';
const pb = new PocketBase("http://localhost:8090");

function userData(s: WebSocket, t: string) {
    let payload = {
        "type": "getplayers",
        "table": t,
    }
    s.send(JSON.stringify(payload));
}
async function getUserObj(id: string) {
    return await pb.collection('users').getOne(id);
}
export { userData, getUserObj };
