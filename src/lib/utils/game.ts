import { PlayerType } from "$lib/utils/players";
import { Cardset } from "./cardset";
import { bet_ } from "$lib/store";
import { turn_ } from "$lib/store";
import { flop_ } from "$lib/store";
import { river_ } from "$lib/store";
import { drawn_ } from "$lib/store";
import { currentPhase_ } from "$lib/store";
import { is_pair } from "$lib/pkg/pokerutil";

let cards = [...Cardset];
let drawnSoFar = [];
function dealCards() {
    drawn_.update(() => []);
    currentPhase_.update(() => 0);
    for (let i = 0; i < 2; i++) {
        let card = draw(cards);
        drawn_.update((drawn_) => [...drawn_, card]);
        cards.splice(cards.indexOf(card), 1);
    }
}
function setFlop(){
    flop_.update(()=> []);
    for (let i = 0; i < 3; i++) {
        let card = draw(cards);
        flop_.update((flop_) => [...flop_, card]);
        cards.splice(cards.indexOf(card), 1);
    }
}
function setTurn(){
    turn_.update(()=> []);
    for (let i = 0; i < 1; i++) {
        let card = draw(cards);
        turn_.update((turn_) => [...turn_, card]);
        cards.splice(cards.indexOf(card), 1);
    }
}
function setRiver(){
    river_.update(() => []);
    for (let i = 0; i < 1; i++) {
        let card = draw(cards);
        river_.update((river_) => [...river_, card]);
        cards.splice(cards.indexOf(card), 1);
    }
    console.log(drawnSoFar);
    console.log(is_pair(drawnSoFar));
}
function draw(cards: string[]) {
    let selected = cards[Math.floor(Math.random() * cards.length)];
    drawnSoFar.push(selected);
    return selected;
}
function nextPhase() {
    let phase = 0; 
    currentPhase_.update((currentPhase_) => currentPhase_ += 1);
    currentPhase_.subscribe((value) => {
        phase = value;
    });
    if (phase === 1){
        setFlop();
    }
    if (phase === 2){
        setTurn();
    }
    if (phase === 3){
        setRiver();
    }

}
let players = [];
let you = new PlayerType("Tommy", 12960, "avatar.webp");
let playerone = new PlayerType("Jack", 32645, "avatar.webp");
players.push(playerone);

function reset() {
    drawn_.update(() => []);
    drawnSoFar = [];
    flop_.update(() => []);
    turn_.update(() => []);
    river_.update(() => []);
    cards = [...Cardset];
    currentPhase_.update(() => -1);
    
}
function betfunc() {
    bet_.subscribe((value) => {
        console.log(value)
    });
}

export { dealCards, currentPhase_, nextPhase, players, you, reset, betfunc };
