<script lang="ts">
    import Card from './Card.svelte';
    import { fade } from 'svelte/transition';
    import Stack from './Stack.svelte';
    import Flop from './Flop.svelte';
    import Turn from './Turn.svelte';
    import River from './River.svelte';
    import Player from './Player.svelte';
    import YourChips from './YourChips.svelte';
    import {dealCards, nextPhase, players, you, reset, betfunc} from '$lib/utils/game';
    import { drawn_, currentPhase_, flop_, turn_, river_, bet_ } from '$lib/store';
    

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
                {#key $flop_}
                    <div in:fade={{delay: 200, duration: 200}} id="flop" style="">
                        <Flop flop={$flop_} />
                    </div>
                {/key}
                {#key $turn_}
                    <div in:fade={{delay: 200, duration: 200}} id="turn" style="">
                        <Turn turn={$turn_} />
                    </div>
                {/key}
                {#key $river_}
                    <div in:fade={{delay: 200, duration: 200}} id="river" style="">
                        <River river={$river_} />
                    </div>
                {/key}
            </div>
        </div>

        <div class="deal justify-center flex flex-row">
            {#if $currentPhase_ === -1}
                <button on:click={dealCards} type="button" class="btn my-5 variant-filled">Deal</button>
            {/if}
            {#if $currentPhase_ >= 0}
                <div class="flex flex-col items-center m-auto justify-center">
                    <form class="flex flex-col justify-center m-auto">
                        <input type="range" bind:value={$bet_} max={you.chips} />
                        <button type="button" on:click={betfunc} class="btn m-auto my-5 variant-filled">Bet</button>
                    </form>
                </div>
                <button on:click={nextPhase} type="button" class="btn mx-2 my-5 variant-filled">Check</button>
                <button on:click={reset} type="button" class="btn mx-2 my-5 variant-filled">Fold</button>
            {/if}
        </div>
    </div>
    <div class="cardset_ flex flex-col justify-start">
        <div class="cardset w-1/4 items-center justify-start flex flex-row">
            {#each $drawn_ as card, index}
                <div in:fade={{delay: index*1000, duration: 200}}>
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
</div>
<style>
.cardset_{
    position: absolute;
    bottom: 0;
    left: 10px;
    padding: 10px;
}

.cardset{
    position: absolute;
    bottom: 0;
    left: 10px;
    filter: brightness(0.7);
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
