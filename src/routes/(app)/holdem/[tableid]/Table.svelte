<script lang="ts">
    //import { fade } from 'svelte/transition';
    import Stack from './Stack.svelte';
    import Player from './Player.svelte';
    import ActionBar from './ActionBar.svelte';
    import { onMount } from 'svelte';
    // Stores
    import { flop, turn, river } from '$lib/stores/game';
    export let data: any;
    export let tableid: any;

    $: currentPhase_ = 0;

    onMount(() => {});
</script>
<div class="flex w-full h-full items-center justify-center flex-col">
    
    <Player data={data}/>
    <div class="flex flex-col items-center justify-center">

        <div class="table flex flex-col items-center justify-center">
            <div class="flex tabcards flex-row">
                <Stack />
                {#if $flop.length != 0}
                    {#each $flop as card}
                        <img class="w-24 p-2 h-36" src="../{card['Rank']+card['Suit']}.png" alt="River Card" /> 
                    {/each}
                {:else}
                    <img class="w-24 p-2 h-36" src="../undefined.png" alt="River Card" />
                    <img class="w-24 p-2 h-36" src="../undefined.png" alt="River Card" />
                    <img class="w-24 p-2 h-36" src="../undefined.png" alt="River Card" />
                {/if}
                {#if $turn.length != 0}
                    <img class="w-24 p-2 h-36" src="../{$turn['Rank']+$turn['Suit']}.png" alt="River Card" />
                {:else}
                    <img class="w-24 p-2 h-36" src="../undefined.png" alt="River Card" />
                {/if}
                {#if $river.length != 0}
                    <img class="w-24 p-2 h-36" src="../{$river['Rank']+$river['Suit']}.png" alt="River Card" />
                {:else}
                    <img class="w-24 p-2 h-36" src="../undefined.png" alt="River Card" />
                {/if}
            </div>
        </div>

    </div>
    <ActionBar currentPhase_={currentPhase_} tableid={tableid} data={data}/> 
</div>
<style>
.tabcards{
    padding: 40px;
    border-radius: 500px / 400px;
    border: 4px solid yellow;
}

.table {
    margin-bottom: 10%;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 70vw;
    height: 70vh;
    background: url('../velv.jpg');
    border-radius: 500px / 400px;
    border: 10px solid #481E14;
    box-shadow: 0 0 90px rgba(0, 0, 0, 0.9) inset, 0 0 90px rgba(0, 0, 0, 1); 
    
}
</style>
