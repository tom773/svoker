<script lang="ts">
    import Card from './Card.svelte';
    import { fade } from 'svelte/transition';
    import { cardset } from '$lib/cardset';
    import Stack from './Stack.svelte';
    import Flop from './Flop.svelte';
    import Turn from './Turn.svelte';
    import River from './River.svelte';
    import { PlayerType } from './players';
    import Player from './Player.svelte';
    import YourChips from './YourChips.svelte';

    let drawn: string[] = [];
    let flop: string[] = [];
    let turn: string[] = [];
    let river: string[] = [];

    let cards = [...cardset];
    function dealCards() {
        drawn = [];
        currentPhase = 0;
        for (let i = 0; i < 2; i++) {
            let card = draw(cards);
            drawn.push(card);
            cards.splice(cards.indexOf(card), 1);
        }
    }
    function setFlop(){
        flop = [];
        for (let i = 0; i < 3; i++) {
            let card = draw(cards);
            flop.push(card);
            cards.splice(cards.indexOf(card), 1);
        }
    }
    function setTurn(){
        turn = [];
        for (let i = 0; i < 1; i++) {
            let card = draw(cards);
            turn.push(card);
            cards.splice(cards.indexOf(card), 1);
        }
    }
    function setRiver(){
        river = [];
        for (let i = 0; i < 1; i++) {
            let card = draw(cards);
            river.push(card);
            cards.splice(cards.indexOf(card), 1);
        }
    }
    function draw(cards: string[]) {
        return cards[Math.floor(Math.random() * cards.length)];
    }
    let currentPhase = -1;
    function nextPhase() {
            
        currentPhase++;
        if (currentPhase === 1){
            setFlop();
        }
        if (currentPhase === 2){
            setTurn();
        }
        if (currentPhase === 3){
            setRiver();
        }

    }
    let players = [];
    let you = new PlayerType("Tommy", 12960, "avatar.webp");
    let playerone = new PlayerType("Jack", 32645, "avatar.webp");
    players.push(playerone);

    function reset() {
        drawn = [];
        flop = [];
        turn = [];
        river = [];
        cards = [...cardset];
        currentPhase = -1;
    }

</script>
<div class="flex w-full h-full items-center justify-center flex-col">
    <div class="flex flex-col items-center justify-center">

        <div class="table flex flex-col items-center justify-center">
            <div class="topplayers absolute">
               {#each players as player}
                    <Player name={player.name} chips={player.chips} avatar={player.avatar}/>
                {/each} 
            </div>
            <div class="flex flex-row">
                <Stack />
                {#key flop}
                    <div in:fade={{delay: 200, duration: 1000}} id="flop" style="">
                        <Flop flop={flop} />
                    </div>
                {/key}
                {#key turn}
                    <div in:fade={{delay: 200, duration: 1000}} id="turn" style="">
                        <Turn turn={turn} />
                    </div>
                {/key}
                {#key river}
                    <div in:fade={{delay: 200, duration: 1000}} id="river" style="">
                        <River river={river} />
                    </div>
                {/key}
            </div>
        </div>

        <div class="deal">
            {#if currentPhase === -1}
                <button on:click={dealCards} type="button" class="btn my-5 variant-filled">Deal</button>
            {/if}
            {#if currentPhase >= 0}
                <button on:click={nextPhase} type="button" class="btn my-5 variant-filled">Check</button>
                <button on:click={reset} type="button" class="btn my-5 variant-filled">Fold</button>
            {/if}
        </div>
    </div>

    <div class="cardset w-1/4 items-center justify-start flex flex-row">
        {#each drawn as card, index}
            <div transition:fade={{delay: index*1000, duration: 400}}>
                {#if index === 0}
                    <Card drawn={card} --rot="-10deg"/>
                {:else} 
                    <Card drawn={card} --rot="10deg"/>
                {/if}
            </div>
        {/each}

    </div>
    <YourChips you={you} />
</div>
<style>

.cardset{
    position: absolute;
    bottom: 0;
    left: 10px;
    filter: brightness(0.8);
    padding: 10px;
}

.table {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 70vw;
    height: 80vh;
    background: url('./velv.jpg');
    border-radius: 500px / 400px;
    border: 10px solid #996515;
    box-shadow: 0 0 90px rgba(0, 0, 0, 0.8) inset, 0 0 45px rgba(0, 0, 0, 0.8); 
    
}

.topplayers {
    display: flex;
    top: 5px;
}


</style>
