<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    import { onMount } from 'svelte';
    import { dealToTable } from '$lib/utils/newapi';
    import PocketBase from 'pocketbase';

    export let data;


    const pb = new PocketBase('http://localhost:8090');

    let socket: WebSocket;
    $: ({user: data});
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
        // Subscribing to PB
        pb.collection('v2gameuser').subscribe('*', function (e) {
            console.log("Test:", e.record);
        });
    });
    
    function dealDeckToTable() {
        dealToTable(socket, data.user.id);
    }
    // This doesn't work. Why. Why doesn't this work.
    async function showdown() {
         
        await pb.collection('v2gameuser').getList(1, 20).then((res) => {
            console.log(res);
        });
    } 

</script>   
<div class="text-xl text-white justify-center w-screen h-screen m-5 p-5 items-center">
    <div class="m-auto">
        <h1 class="text-3xl text-white">Test V2 API Endpoints Here:</h1><br><hr><br>
        <Button on:click={dealDeckToTable}>Deal</Button>
        <Button on:click={showdown}>Showdown!</Button>
        <div class="inline-flex text-white">
            <h1>User Cards:</h1>
            {#each usercards as card}
                <img class="w-16 h-20" src="../{card['Rank']+card['Suit']}.png" alt="f"/>
            {/each}
        </div>
    </div>
</div>
