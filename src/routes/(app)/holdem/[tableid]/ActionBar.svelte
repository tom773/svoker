<script lang="ts">
    // Components
    import { Slider } from '$lib/components/ui/slider';
    import { Button } from '$lib/components/ui/button/index';
    import Card from './Card.svelte';
    import Result from './\(results)/Result.svelte';
    import YourChips from './YourChips.svelte';
    // Functions
    import { fade } from 'svelte/transition';
    import { dealToTable, resetTable, showdown } from '$lib/utils/newapi';
    // Stores
    import { websocket } from '$lib/stores/websocket';
    import { hand } from '$lib/stores/game';
    import { get } from 'svelte/store';
    // Exports
    export let currentPhase_: any;
    export let data: any;
    export let tableid: any;
    
    // Websocket stuff from here
    let gameid: string;
     
    for (let game of data.games){
        if (tableid == game['table']){
            gameid = game['id'];
        }
    }

    function showDown() {
        const ws: WebSocket = get(websocket);
        showdown(ws, data.user.id, gameid);
    }
    function dealDeckToTable() {
        const ws: WebSocket = get(websocket);
        dealToTable(ws, data.user.id, gameid);
    }
    function reset() {
        const ws: WebSocket = get(websocket);
        resetTable(ws, data.user.id, gameid);
        hand.set([]);
    }

</script>
<div class="bar"> 
{#if data.user}
    <div class="actions">
        <div class="cardset_ m-auto flex flex-row justify-evenly">
            <div class="flex items-center actionprof">
                <YourChips data={data}/>
            </div>
            <div style="font-size: 18px;" class="deal m-auto justify-start mx-5 items-center flex flex-row">
                    <div class="flex flex-col items-center m-auto justify-center">
                        <form class="flex flex-col items-center justify-center mx-5 m-auto">
                            <Slider class="w-32 my-2" max={data.user.balance} />
                            <div class="flex flex-row items-center justify-center">
                                <Button style="font-size: 18px;" variant="ghost" type="button" on:click={()=>console.log("Bet Event")} class="btn my-2 m-auto variant-filled">Bet</Button>
                                <p style="font-size: 18px;">$0</p>
                            </div>
                        </form>
                    </div>
                    <Button style="font-size: 18px;" variant="ghost" on:click={()=>dealDeckToTable()} type="button" class="btn mx-2 my-5 variant-filled">Check (Play Entire Game)</Button>
                    <Button style="font-size: 18px;" variant="ghost" on:click={()=>reset()} type="button" class="btn mx-2 my-5 variant-filled">Fold</Button>
                    <Button style="font-size: 18px;" variant="ghost" on:click={()=>showDown()} type="button" class="btn mx-2 my-5 variant-filled">Showdown</Button>
                
            </div>
            <div class="cardset justify-center items-center flex flex-row">
                <div class="cs items-center justify-center flex flex-row" in:fade={{delay: 1000, duration: 200}}>
                    {#if $hand != null}
                        {#if $hand[0]}
                            <div class="cs items-center justify-center flex flex-row" in:fade={{delay: 1000, duration: 200}}>
                                <Card drawn={$hand[0]} --rot="-10deg"/>
                                <Card drawn={$hand[1]} --rot="10deg"/>
                            </div>
                        {/if}
                    {/if}
                </div>
                <div class="mx-20">
                    <Result />
                </div>
            </div>
            <div class="flex items-center nexthand">
                {#if currentPhase_ == -1}
                    <Button variant="ghost" on:click={()=>reset()} type="button" style="display: inline-flex; align-items: center;" class="btn my-5 w-40 variant-filled">
                        <img width="32" src="../next.png" alt="next"/>&nbsp;Deal</Button>
                {:else}
                    <Button variant="ghost" on:click={()=>reset()} type="button" style="display: inline-flex; align-items: center;" class="btn my-5 w-40 variant-filled">
                        <img width="32" src="../next.png" alt="next"/>&nbsp;Skip Hand</Button>
                {/if}
            </div>
        </div>
    </div>
{:else}
    <div class="actions">
        <div class="cardset_ m-auto flex flex-row justify-evenly">
            <Button href="/signin">Sign In</Button>
        </div>
    </div>
{/if}
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

