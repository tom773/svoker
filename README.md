<h1>Svoker</h1>

Learning Svelte, Networking, Websockets by building a poker game. Aim is to have this be a full fledged poker game with multiplayer support.

<h3>To Do</h3>

<hr><h4>To Do</h4>

- [ ] Send drawn hands to DB. 
- [ ] Implement table specific state for WebSockets
- [ ] WebSockets for game state updates
- [ ] Chat - WebSockets? 
- [ ] Betting 
- [ ] Advanced Hand Evaluation 

<h4>Done</h4>

- [x] Basic card dealing
- [x] Basic poker hand evaluation
- [x] Flop, River and Turn Simulations 
- [x] Cleaner UI
- [x] SSE for seat management. Some cleanup to do but API is working. 
- [x] WebSockets for dealing

<h4><strong>Note to self:</strong> You're onto something with unmarshalling DB tables into Go structs. We could use this persist game state and still use websockets to trigger events and
hydration of client tables.</h4>
