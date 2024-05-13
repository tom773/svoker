import { writable } from 'svelte/store';
import PocketBase from 'pocketbase';

const pb = new PocketBase('http://localhost:8090');

export const handStore = async (userid: string, tableid: string) => {
    const { subscribe, set } = writable([]); 
    let socket: any; 
    
    async function connect(){
        const record = await pb.collection('gametable').getList(1, 50, {
            "filter": `user = "${userid}" && table = "${tableid}"`,
        });

        socket = await pb.collection('gametable').subscribe(record.items[0].id, function (e) {
            if (e.record && e.record.cards) {
                set(e.record.cards);
            }
        });

        await pb.collection('gametable').getOne(record.items[0].id).then((cards) => {
            if (cards && cards.cards) {
                set(cards.cards);
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

export const hand_store = (userid: string, tableid: string) => handStore(userid, tableid);

