import { writable } from 'svelte/store';
import PocketBase from 'pocketbase';

const pb = new PocketBase('http://localhost:8090');

export const gameStore = async (tableid: string) => {
    const { subscribe, set } = writable([]); 
    let socket: any; 
    
    const record = await pb.collection('game').getList(1, 50, {
        "filter": `table = "${tableid}"`
    });
    let gameId = record.items[0].id;

    async function connect(){
        socket = await pb.collection('game').subscribe(gameId, function (e) {
            if (e.record && e.record.phases) {
                set(e.record.phases);
            }
        });

        await pb.collection('game').getOne(gameId).then((table) => {
            if (table && table.phases) {
                set(table.phases);
            }
        });
    }

    async function disconnect(){
        if(socket){
            pb.collection("game").unsubscribe(socket);
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

export const tableStore = (tableid: string) => gameStore(tableid);
