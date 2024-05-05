import { cardset } from './cardset';

console.log(cardset);

let cards = ['aceh', 'acec', 'threec', 'fourc', 'sevenh', 'fiveh', 'queend']
let suits = new Map<String, String>();

suits.set('h', 'hearts');
suits.set('c', 'clubs');
suits.set('d', 'diamonds');
suits.set('s', 'spades');

class Card {

    suit: string;
    rank: number;

    constructor(suit: string, rank: number){

        this.suit = suit;
        this.rank = rank;

    }

}

for (let i = 0; i < cards.length; i++){

    var lastChar = cards[i].slice(-1);
    var allChar = cards[i].slice(0, -1);
    console.log(allChar)

}
