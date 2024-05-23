// Blueprint for the new API

function deal(s: WebSocket) {
    let sampleData = {
        "type": "deal"
    }
    s.send(JSON.stringify(sampleData));
}

function doGame(msg: any) {
    let gameCards = [];
    let parsed = JSON.parse(msg);
    for (let i = 0; i < 7; i++) {
        let card: string = parsed[i].Rank + parsed[i].Suit;
        gameCards.push(card);
    }

    return gameCards;
}

export { deal, doGame };
