<script lang="ts">
    import { Button } from '$lib/components/ui/button';
    import { onMount } from 'svelte';
    import { deal, doGame } from '$lib/utils/newapi';
    //export let data;
    
    let socket: WebSocket;
    $: cards = [];
    onMount(() => {
        socket = new WebSocket('ws://localhost:8080/ws');
        
        socket.onopen = () => {
            console.log('Connected to server');
        };

        socket.onmessage = (event) => {
            const dealtCards = doGame(event.data);
            cards = dealtCards;
            console.log(cards);
        };
        
        socket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
    });

    function handleClick() {
        deal(socket);
    }
    
     

</script>   
<div class="text-xl text-white justify-center w-screen h-screen m-5 p-5 items-center">
    <div class="m-auto">
        <Button on:click={handleClick}>Deal</Button>
        <div class="text-white">
        </div>
    </div>
    <div class="py-5">
        <div class="py-5 m-auto">
            <div class="inline-flex text-white">
                <img class="w-24 h-32 pr-2" src="../{cards[0]}.png" alt={"memes"} />
                <img class="w-24 h-32 pr-2" src="../{cards[1]}.png" alt={"memes"} />
            </div>
        </div>
        <div class="py-5 m-auto">
            <div class="inline-flex text-white">
            </div>
        </div>
    </div>
    <div class="py-5">
        <div class="py-5 m-auto">
            <div class="inline-flex text-white">
                <img class="w-24 h-32 pr-2" src="../{cards[2]}.png" alt={"memes"} />
                <img class="w-24 h-32 pr-2" src="../{cards[3]}.png" alt={"memes"} />
                <img class="w-24 h-32 pr-2" src="../{cards[4]}.png" alt={"memes"} />
                <img class="w-24 h-32 pr-2" src="../{cards[5]}.png" alt={"memes"} />
                <img class="w-24 h-32 pr-2" src="../{cards[6]}.png" alt={"memes"} />
            </div>
        </div>
    </div>

</div>
