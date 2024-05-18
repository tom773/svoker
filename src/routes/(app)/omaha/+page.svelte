<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    import { onMount } from 'svelte';
    export let data;
    
    let socket: WebSocket;
    let messages = [];
    let tableID = ['71wh85i204kt8zx'];

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
    $: com = [];
    $: hand = []; 

    function deal() {
        sendMessage({ type: 'deal', tableID: tableID[0] });
    }
    function comfunc() {
        sendMessage({ type: 'com', tableID: tableID[0] });
    }
    function handleClick() {
        sendMessage({ type: 'reset', tableID: tableID[0]});
    }
    
    onMount(() => {
        socket = new WebSocket('ws://localhost:8090/ws');
        
        socket.onopen = () => {
            console.log('Connected to server');
        };
        
        socket.onmessage = (event) => {

            const d = JSON.parse(event.data);
            handleMessage(d);

        };
        socket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
    });
    function sendMessage(message: any) {
        if (socket && socket.readyState === WebSocket.OPEN) {
          socket.send(JSON.stringify(message));
        } else {
          console.error('WebSocket is not open');
        }
    }
    function handleMessage(message: any) {
        switch (message.type) {
          case 'dealResponse':
            messages.push(`Dealt cards: ${message.cards.join(', ')}`);
            break;
          case 'comResponse':
            messages.push(`Community cards: ${message.cards.join(', ')}`);
            break;
          case 'resetResponse':
            messages.push(message.msg);
            break;
          case 'message':
            messages.push(message.message);
            break;
          default:
            console.error('Unknown message type:', message);
        }
      }
    

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
    <div class="py-5">
        <Button on:click={comfunc}>Deal myself cards</Button>
        <div class="py-5 m-auto">
            <div class="inline-flex text-white">
                {#if com}
                    {#each com as comcard}
                        <img class="w-24 h-32 pr-2" src="../{comcard}.png" alt={comcard} />
                    {/each}
                {/if}
            </div>
        </div>
    </div>

</div>
