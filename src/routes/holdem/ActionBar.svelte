<script>
    import { Slider } from '$lib/components/ui/slider';
    import { Button } from '$lib/components/ui/button/index';
    import Card from './Card.svelte';
    import { fade } from 'svelte/transition';
    import {dealCards, nextPhase, you, reset, betfunc} from '$lib/utils/game';
    import { drawn_, currentPhase_, bet_ } from '$lib/store';
    import YourChips from './YourChips.svelte';
    import Result from './\(results)/Result.svelte';
</script>

<div class="bar"> 
    <div class="actions">
        <div class="cardset_ m-auto flex flex-row justify-evenly">
            <div class="flex items-center actionprof">
                <YourChips player={you}/>
            </div>
            <div style="font-size: 18px;" class="deal m-auto justify-start mx-5 items-center flex flex-row">
                {#if $currentPhase_ >= 0}
                    <div class="flex flex-col items-center m-auto justify-center">
                        <form class="flex flex-col items-center justify-center mx-5 m-auto">
                            <Slider class="w-32 my-2" bind:value={$bet_} max={you.chips} />
                            <div class="flex flex-row items-center justify-center">
                                <Button style="font-size: 18px;" variant="ghost" type="button" on:click={betfunc} class="btn my-2 m-auto variant-filled">Bet</Button>
                                <p style="font-size: 18px;">${$bet_}</p>
                            </div>
                        </form>
                    </div>
                    <Button style="font-size: 18px;" variant="ghost" on:click={nextPhase} type="button" class="btn mx-2 my-5 variant-filled">Check</Button>
                    <Button style="font-size: 18px;" variant="ghost" on:click={reset} type="button" class="btn mx-2 my-5 variant-filled">Fold</Button>
                {/if}
            </div>
            <div class="cardset justify-center flex flex-row">
                {#each $drawn_ as card, index}
                    <div class="cs" in:fade={{delay: index*1000, duration: 200}}>
                        {#if index === 0}
                            <Card drawn={card} --rot="-10deg"/>
                        {:else} 
                            <Card drawn={card} --rot="10deg"/>
                        {/if}
                    </div>
                {/each}
                <div class="mx-20">
                    <Result />
                </div>
            </div>
            <div class="flex items-center nexthand">
                {#if $currentPhase_ === -1}
                    <Button variant="ghost" on:click={dealCards} type="button" style="display: inline-flex; align-items: center;" class="btn my-5 w-40 variant-filled">
                        <img width="32" src="./next.png" alt="next"/>&nbsp;Deal</Button>
                {:else}
                    <Button variant="ghost" on:click={reset} type="button" style="display: inline-flex; align-items: center;" class="btn my-5 w-40 variant-filled">
                        <img width="32" src="./next.png" alt="next"/>&nbsp;Skip Hand</Button>
                {/if}
            </div>
        </div>
    </div>
</div>

<style>
.cs {
   bottom: 10px; 
}
.cardset {
    object-fit: contain;
    display: flex;
    justify-content: start;
    align-items: center;
    width: 100%;
    height: 100%;
}
.cardset_ {
    height: 100%;
    width: 100%;
}
.bar {
    width: 100%;
    height: 10%;
    position: fixed;
    bottom: 0;
    padding: 0 1rem;
    background-color: #0F0F0F;
}
.actions {
    display: flex;
    width: 100%;
    height: 100%;
}
</style>

