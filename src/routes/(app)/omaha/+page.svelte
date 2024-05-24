<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    import { onMount } from 'svelte';
    import { dealToTable, gameLoop } from '$lib/utils/newapi';
    export let data;

    let socket: WebSocket;
    
    $: cards = [];
    onMount(() => {

        socket = new WebSocket('ws://localhost:8080/ws');
        
        socket.onopen = () => {
            console.log('Connected to server');
        };

        socket.onmessage = (event) => {
            const dealtCards = gameLoop(event.data);
            cards = dealtCards;
            console.log(cards);
        };
        
        socket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
    });

    function dealDeckToTable() {
        dealToTable(socket);
    }
    
    function tnum() {
        if (data.tables[0].players.includes(data.user.id)){
            return 1;
        }
    } 

</script>   
<div class="text-xl text-white justify-center w-screen h-screen m-5 p-5 items-center">
    <div class="m-auto">
        <h1 class="text-3xl text-white">Test V2 API Endpoints Here:</h1><br><hr><br>
        <Button on:click={dealDeckToTable}>Deal</Button>
        <Button on:click={tnum}>Deal for Table 2</Button>
        <div class="text-white">
        </div>
    </div>
</div>
