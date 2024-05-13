<script lang="ts">
    import { fade } from 'svelte/transition';
    import Stack from './Stack.svelte';
    import Flop from './Flop.svelte';
    import Turn from './Turn.svelte';
    import River from './River.svelte';
    import Player from './Player.svelte';
    import ActionBar from './ActionBar.svelte';
    import { onMount } from 'svelte';
    import { gameStore } from '$lib/stores/game';
    import { loopStore } from '$lib/stores/loop';
    export let data: any;
    export let tableid: any;
     
    $: cards = [];
    $: currentPhase_ = 0;
    onMount(async () => {
        let store = await gameStore(tableid);
        store.subscribe(value =>  {
            cards = value;
        });
        let loopstore = await loopStore(tableid);
        loopstore.subscribe((value) => {
            let currentPhase = value;
            currentPhase_ = currentPhase as number;
        });
    });

</script>
<div class="flex w-full h-full items-center justify-center flex-col">
    
    <Player data={data} tableid={tableid}/>
    <div class="flex flex-col items-center justify-center">

        <div class="table flex flex-col items-center justify-center">
            <div class="flex tabcards flex-row">
                <Stack />
                    {#if currentPhase_ >= 1}
                        <div in:fade={{delay: 200, duration: 200}} id="flop" style="">
                            <Flop flop={cards} />
                        </div>
                    {:else}
                        <div in:fade={{delay: 200, duration: 200}} id="flop">
                            <Flop flop={[]} />
                        </div>
                    {/if}
                    {#if currentPhase_ >= 2}
                        <div in:fade={{delay: 200, duration: 200}} id="flop" style="">
                            <Turn turn={[cards[3]]} />
                        </div>
                    {:else}
                        <div in:fade={{delay: 200, duration: 200}} id="flop">
                            <Turn turn={[]} />
                        </div>
                    {/if}
                    {#if currentPhase_ >= 3}
                        <div in:fade={{delay: 200, duration: 200}} id="river" style="">
                            <River river={[cards[4]]}/>
                        </div>
                    {:else}
                        <div in:fade={{delay: 200, duration: 200}} id="river">
                            <River river={[]} />
                        </div>
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
