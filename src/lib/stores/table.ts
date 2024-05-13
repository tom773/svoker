import PocketBase from 'pocketbase';
import { writable } from 'svelte/store';
const pb = new PocketBase('http://localhost:8090');

async function createPlayerStore(tableid: string){
    const { subscribe, set } = writable([]); 
    let socket: any; 
    
    async function connect(){
        socket = await pb.collection('tables').subscribe(tableid, function (e) {
            if (e.record && e.record.players) {
                set(e.record.players);
            }
        });

        await pb.collection('tables').getOne(tableid).then((table) => {
            if (table && table.players) {
                set(table.players);
            }
        });
    }

    async function disconnect(){
        if(socket){
            pb.collection("tables").unsubscribe(socket);
        }
    }
    connect();
    return {
        subscribe,
        reconnect: () => {
            disconnect();
            connect();
        }, 
        disconnect,
    }
}
export const playerStore = (tableid: string) => createPlayerStore(tableid);

