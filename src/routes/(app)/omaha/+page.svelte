<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    import { onMount } from 'svelte';
    import { dealToTable, resetTable } from '$lib/utils/newapi';

    import PocketBase from 'pocketbase';
    const pb = new PocketBase('http://localhost:8090');
    
    export let data;
    let socket: WebSocket;
    
    $: usercards = [];
    $: communityCards = [];

    onMount(() => {
        // Socket stuff
        socket = new WebSocket('ws://localhost:8080/ws');
        
        socket.onopen = () => {
            console.log('Connected to server');
        };

        socket.onmessage = (event) => {
            console.log(event.data);
        };
        
        socket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
        //Subscribing to PB
        pb.collection('v2gameuser').subscribe('*', function (e) {
            usercards = e.record['hand'];
            console.log('User cards:', usercards);
        });
        pb.collection('v2game').subscribe('*', function (e) {
            let ccards = e.record['deck'];
            communityCards = ccards.slice(0, 5);
        });

    });
    
    function dealDeckToTable() {
        dealToTable(socket, data.user.id);
    }

    function reset() {
        resetTable(socket, data.user.id);
        usercards = [];
        communityCards = [];
    }

</script>   
<div class="text-xl text-white justify-center w-screen h-screen m-5 p-5 items-center">
    <div class="m-auto">
        <h1 class="text-3xl text-white">Test V2 API Endpoints Here:</h1><br><hr><br>
        <Button on:click={dealDeckToTable}>Deal a Game</Button>
        <Button on:click={reset}>Reset Table</Button>
        <div class="text-white">
            <h1>User Cards:</h1>
            <div class="inline-flex">
                {#each usercards as card}
                    <img class="w-16 h-24" src="../{card['Rank']+card['Suit']}.png" alt="f"/>
                {/each}
            </div>
            <h1>Community Cards:</h1>
            <div class="inline-flex">
                {#each communityCards as card}
                    <img class="w-16 h-24" src="../{card['Rank']+card['Suit']}.png" alt="f"/>
                {/each}
            </div>
        </div>
    </div>
</div>
