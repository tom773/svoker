<script lang="ts">
    import * as Avatar from '$lib/components/ui/avatar';
    export let data: any;
    import { onMount } from 'svelte';
    import { numberWithCommas, getImageURL } from '$lib/utils';
    import { LoaderCircle } from 'lucide-svelte';
    import { fade } from 'svelte/transition';

    let socket: WebSocket;
    
    onMount(async () => {
        socket = new WebSocket('ws://localhost:8080/ws/user');

        socket.onopen = () => {
            //Do Nothing
        };
        socket.onmessage = (event: any) => {
            console.log(event.data);
        };
        socket.onerror = (error: any) => {
            console.error('WebSocket error:', error);
        };
    });
    
</script>
<div class="player1 flex flex-row">
    {#if data.tablePlayers.length == 0} 
        <span class="loading" transition:fade={{delay: 0, duration: 300}}>
        &nbsp;<LoaderCircle class="inline-flex w-6 left-0 h-6 absolute animate-spin" />
        </span>
    {/if}
    {#each data.tablePlayers as player}
        {#if player["id"] != data.user.id}
            <div class="flex avatar">
                <Avatar.Root class="w-20 h-20">
                    <Avatar.Image class="bg-white border-4 rounded-full border-gray-600 cursor-pointer hover:border-gray-400" src={getImageURL(player["id"], player["avatar"])} alt="avatar" />
                </Avatar.Root>
            </div>
            <div class="flex p-3 text-left flex-col">
                <p>{player["username"]}</p>
                <p>${numberWithCommas(player["balance"])}</p>
            </div>
        {/if}
    {/each}
    </div>
<style>
    .player1{
        align-items: center;
        justify-content: center;
        margin: 10px;
        margin-left: 5px;
        font-size: 15px;
        text-align: center;
    }

</style>
