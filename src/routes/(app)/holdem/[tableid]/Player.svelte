<script lang="ts">
    import * as Avatar from '$lib/components/ui/avatar';
    export let data: any;
    export let tableid: any;
    import { onMount } from 'svelte';
    import { numberWithCommas, getImageURL } from '$lib/utils';
    import { LoaderCircle } from 'lucide-svelte';
    import { fade } from 'svelte/transition';
    import { userData, getUserObj } from '$lib/utils/userapi';

    $: players = [];
    $: playersData = [];
    $: changed = 0;

    let socket: WebSocket;
    function getPlayers(){
        socket = new WebSocket('ws://localhost:8080/ws/user');
        socket.onopen = () => {
            userData(socket, tableid);
        };
        socket.onmessage = (event: any) => {
            let tmp = JSON.parse(event.data);
            for (let player of tmp){
                players.push(player);
            }
            test(players);
        };
        socket.onerror = (error: any) => {
            console.error('WebSocket error:', error);
        };
    }
    async function test(players: any){
        for (let player of players){
            playersData.push(await getUserObj(player));
            changed += 1;
        }
    }
    onMount(async () => {
        getPlayers();
    });
    
</script>
<div class="player1 flex flex-row">
    {#if playersData.length == 0} 
        <span class="loading" transition:fade={{delay: 0, duration: 300}}>
        &nbsp;<LoaderCircle class="inline-flex w-6 left-0 h-6 absolute animate-spin" />
        </span>
    {/if}
    {#key changed}
        {#each playersData as player}
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
    {/key}
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
