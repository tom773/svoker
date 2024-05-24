// Blueprint for the new API

function gameLoop(s: WebSocket) {
    let sampleData = {
        "type": "deal"
    }
    s.send(JSON.stringify(sampleData));
}

function dealToTable(msg: any) {
    let gameCards = [];
    let parsed = JSON.parse(msg);

    for (let i = 0; i < 52; i++) {
        let card: string = parsed[i].Rank + parsed[i].Suit;
        gameCards.push(card);
    }

    return gameCards;
}

export { gameLoop, dealToTable };
