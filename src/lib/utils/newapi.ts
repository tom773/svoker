// Blueprint for the new API

function dealToTable(s: WebSocket, user: string) {
    let gameCards = [];
    
    let payload = {
        "type": "deal",
        "user": user,
        "gameID": "zfu2qstfnvqggr7",
    }

    s.send(JSON.stringify(payload));

    return gameCards;
}

export { dealToTable };
