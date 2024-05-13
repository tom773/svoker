import { Cardset } from "./cardset";
import { bet_ } from "$lib/store";
import { currentPhase_ } from "$lib/store";
import { is_pair } from "$lib/pkg/pokerutil";
import { ranks_ } from "$lib/store";
import { handtype_ } from "$lib/store";
import { suits_ } from "$lib/store";
import PocketBase from 'pocketbase';
import { json } from '@sveltejs/kit';   

const pb = new PocketBase('http://localhost:8090');

let cards = [...Cardset];
let drawnSoFar = [];

async function updateDrawn(id: string){
    
    const alreadydrawn = await pb.collection('game').getList(1, 50, {
        "filter": `table = "${id}"`
    });
    await pb.collection('game').update(alreadydrawn.items[0].id, {
        "drawn": drawnSoFar,
    });
}

async function dealCards(userid: string, tableid: string){
    let dealt = [];
    for (let i = 0; i < 2; i++) {
        let card = draw(cards);
        dealt.push(card);
        cards.splice(cards.indexOf(card), 1);
    }
    // Fetch record that contains user & table info
    const record = await pb.collection('gametable').getList(1, 50, {
        "filter": `table = "${tableid}" && user = "${userid}"`
    });
    // Fetch table data
    const gamedata = await pb.collection('game').getList(1, 50, {
        "filter": `table = "${tableid}"`
    });
    // Draw two cards and sent to local users record in gametables
    await pb.collection('gametable').update(record.items[0].id, {
        "cards": {"c1": dealt[0], "c2": dealt[1]},
    });
    // Update game data with drawn cards
    await updateDrawn(tableid);
    await nextaction(tableid); 

}
async function setFlop(tableid: string){
    
    let flop = [];
    for (let i = 0; i < 3; i++) {
        let card = draw(cards);
        flop.push(card);
        cards.splice(cards.indexOf(card), 1);
    }
    const gamedata = await pb.collection('game').getList(1, 50, {
        "filter": `table = "${tableid}"`
    });
    await pb.collection('game').update(gamedata.items[0].id, {
        "phases": flop
    });
    await updateDrawn(tableid);

}
async function setTurn(tableid: string){
    let turn = [];
    let card = draw(cards);
    turn.push(card);
    cards.splice(cards.indexOf(card), 1);
    const gamedata = await pb.collection('game').getList(1, 50, {
        "filter": `table = "${tableid}"`
    });
    turn = gamedata.items[0].phases.concat(turn);
    await pb.collection('game').update(gamedata.items[0].id, {
        "phases": turn
    });
    await updateDrawn(tableid);
}
async function setRiver(tableid: string){
    let river = [];
    let card = draw(cards);
    river.push(card);
    cards.splice(cards.indexOf(card), 1);
    const gamedata = await pb.collection('game').getList(1, 50, {
        "filter": `table = "${tableid}"`
    });
    river = gamedata.items[0].phases.concat(river);
    await pb.collection('game').update(gamedata.items[0].id, {
        "phases": river,
    });
    await updateDrawn(tableid);

    let result = JSON.parse(is_pair(drawnSoFar));
    ranks_.update(()=>result.ranks);
    handtype_.update(()=>result.hand);
    suits_.update(()=>result.suits);
}
function draw(cards: string[]) {
    let selected = cards[Math.floor(Math.random() * cards.length)];
    drawnSoFar.push(selected);
    return selected;
}

async function resetCards(tableid: string, userid: string){
    const record = await pb.collection('gametable').getList(1, 50, {
        "filter": `table = "${tableid}"`
    });
    for (let i = 0; i < record.items.length; i++){
        await pb.collection('gametable').update(record.items[i].id, {
            "cards": {"c1": "", "c2": ""},
        });
    }
    const gamedata = await pb.collection('game').getList(1, 50, {
        "filter": `table = "${tableid}"`
    });
    await pb.collection('game').update(gamedata.items[0].id, {
        "phases": [],
        "action": -1,
        "drawn": []
    });
}

function reset(tableid: string, userid: string) {1
    drawnSoFar = [];
    resetCards(tableid, userid);
    cards = [...Cardset];
    ranks_.update(()=>[]);
    handtype_.update(()=>"")
    suits_.update(()=>[])
    
}
async function nextaction(id: string){
    const alreadydrawn = await pb.collection('game').getList(1, 50, {
        "filter": `table = "${id}"`
    });
    await pb.collection('game').update(alreadydrawn.items[0].id, {
        "action+": 1,
    });
}

function betfunc() {
    bet_.subscribe((value) => {
        console.log(value); 
    });
}

export {nextaction, dealCards, setFlop, setTurn, setRiver, reset, betfunc };
