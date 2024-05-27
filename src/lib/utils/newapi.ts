// This will replace game.ts probably

function dealToTable(s: WebSocket, user: string, table: string) {
    let gameCards = [];
    
    let payload = {
        "type": "deal",
        "user": user,
        "gameID": table,
    }

    s.send(JSON.stringify(payload));

    return gameCards;
}

function resetTable(s: WebSocket, user: string, table: string) {
    let payload = {
        "type": "reset",
        "user": user,
        "gameID": table,
    }

    s.send(JSON.stringify(payload));
}

function showdown(s: WebSocket) {
    let payload = {
        "type": "showdown",
    }

    s.send(JSON.stringify(payload));
}

export { dealToTable, resetTable, showdown };

// Poker Game:
// 1. Deal - Implemented
// 2. Preflop Betting
// 3. Flop
// 4. Flop Betting
// 5. Turn
// 6. Turn Betting
// 7. River
// 8. River Betting
// 9. Showdown/Winning
// 10. Reset - Implemented
