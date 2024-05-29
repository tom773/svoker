import { writable } from 'svelte/store';
import { hand, flop, turn, river } from './game';

export const websocket = writable<WebSocket | null>(null);

export function initWS(url: string) {
    const ws = new WebSocket(url);
    ws.onopen = () => {
        console.log('connected');    
    };
    ws.onclose = () => {
        console.log('disconnected');
        websocket.set(null);
    };
    ws.onerror = () => {
        console.log('error');
        websocket.set(null);
    };
    ws.onmessage = (e) => {
        let parsed = JSON.parse(e.data);
        let key = Object.keys(parsed)[0] as string;
        switch(key){
            case "hand":
                hand.set(parsed["hand"]);
                break;
            case "deck":
                flop.set(parsed["deck"].slice(0, 3));
                turn.set(parsed["deck"][4]);
                river.set(parsed["deck"][5]);
                break;
            case "reset":
                hand.set([]);
                flop.set([]);
                turn.set([]);
                river.set([]);
                break;
            default:
                break;
        }
    }
    
    websocket.set(ws);
}
