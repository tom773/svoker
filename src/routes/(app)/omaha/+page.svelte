<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    import { onMount } from 'svelte';
    export let data;
    
    let socket: WebSocket;
    interface GameState {
        flop: [];
        turn: [];
        river: [];
    };

    let gameState: GameState = {
        flop: [],
        turn: [],
        river: []
    }; 
    $: hand = []; 

    function deal() {
        socket.send(JSON.stringify({ type: 'deal'}));
    }

    function handleClick() {
        hand = [];
    }
    
    onMount(() => {
        socket = new WebSocket('ws://localhost:8090/ws');
        
        socket.onopen = () => {
            console.log('Connected to server');
        };
        
        socket.onmessage = (event) => {
            const d = JSON.parse(event.data);
            if (d.type === 'dealResponse') {
                hand = d.cards;
            }
            console.log(d.cards);
        };
        socket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
    });

</script>   
<div class="text-xl text-white justify-center w-screen h-screen m-5 p-5 items-center">
    <div class="m-auto">
        <Button on:click={handleClick}>Test</Button>
        <div class="text-white">
        </div>
    </div>
    <div class="py-5">
        <Button on:click={deal}>Deal myself cards</Button>
        <div class="py-5 m-auto">
            <div class="inline-flex text-white">
                {#if hand}
                    {#each hand as card}
                        <img class="w-24 h-32 pr-2" src="../{card}.png" alt={card} />
                    {/each}
                {/if}
            </div>
        </div>
    </div>
</div>
