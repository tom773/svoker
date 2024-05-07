<h1>Svoker</h1>

Learning Svelte by building a poker game.

Now featuring WebAssembly! Ever wondered how to Shoehorn more Rust into a simple frontend project? Look no further than WASM!

<h2>How to run locally</h2>

1. Clone the repository
2. Run `npm install`
3. Spin up a postgres DB & Create a `.env` file with a `DATABASE_URL` variable pointing to it 
4. Run `npx primsa migrate dev --init`
5. Run `npx prisma generate`
6. Run `npm run dev`
7. Play poker!

The WebAssembly is already compiled, but if you make changes to lib.rs, simple run `make` to recompile it.

<h3>Features</h3>
<ul>
    <li>[*] Basic card dealing</li>
    <li>[*] Basic poker hand evaluation</li>
    <li>[*] Flop, River and Turn Simulations </li>
    <hr>Planned
    <li>[ ] Cleaner UI</li>
    <li>[ ] Hand Comparison</li>
    <li>[ ] AI</li>
    <li>[ ] Betting</li>
    <li>[ ] Seats</li>
    <li>[ ] Chat</li>
    <li>[ ] User Authentication</li>
    <li>[ ] Multiplayer</li>
</ul>
