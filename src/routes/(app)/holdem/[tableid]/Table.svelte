<script lang="ts">
    import { fade } from 'svelte/transition';
    import Stack from './Stack.svelte';
    import Flop from './Flop.svelte';
    import Turn from './Turn.svelte';
    import River from './River.svelte';
    import Player from './Player.svelte';
    import ActionBar from './ActionBar.svelte';
    import { onMount } from 'svelte';
    export let data: any;
    export let tableid: any;
    import PocketBase from 'pocketbase';
    const pb = new PocketBase("http://localhost:8090");

    $: cards = [];
    $: currentPhase_ = 0;
    $: key_ = 0;

    let socket: WebSocket;
    onMount(() => {
        socket = new WebSocket('ws://localhost:8080/ws');
        socket.onopen = () => {
            console.log('Connected to server');
        };
        socket.onmessage = (event: any) => {
            console.log(event.data);
        };
         
        socket.onerror = (error: any) => {
            console.error('WebSocket error:', error);
        };
        pb.collection('v2game').subscribe('*', function (e) {
            if (e.record['table'] == tableid){
                let cardset = e.record['deck'];
                cards = cardset.slice(0, 5);
                key_++;
            }
        });
    });

</script>
<div class="flex w-full h-full items-center justify-center flex-col">
    
    <Player data={data} tableid={tableid}/>
    <div class="flex flex-col items-center justify-center">

        <div class="table flex flex-col items-center justify-center">
            <div class="flex tabcards flex-row">
                <Stack />
                {#key key_}
                    {#if cards.length != 0}
                        {#each cards as card}
                            <img class="w-24 p-2 h-36" src="../{card['Rank']+card['Suit']}.png" alt="River Card" /> 
                        {/each}
                    {/if}
                {/key}
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
