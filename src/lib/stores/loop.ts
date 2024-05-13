import { writable } from 'svelte/store';
import PocketBase from 'pocketbase';

const pb = new PocketBase('http://localhost:8090');

export const loopStore = async (tableid: string) => {
    const { subscribe, set } = writable(); 
    
    let socket: any; 
    

    async function connect(){
        const record = await pb.collection('game').getList(1, 50, {
            "filter": `table = "${tableid}"`
        });
        let gameId = record.items[0].id;
        socket = await pb.collection('game').subscribe(gameId, function (e) {
            set(e.record.action);
        });

        await pb.collection('game').getOne(gameId).then(table => {
            set(table.action);
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

export const tableStore = (tableid: string) => loopStore(tableid);
