<script lang="ts">
    import * as Avatar from '$lib/components/ui/avatar';
    export let data: any;
    export let tableid: any;
    import { onMount } from 'svelte';
    import PocketBase from 'pocketbase';
    const pb = new PocketBase("http://localhost:8090");
    import { numberWithCommas, getImageURL } from '$lib/utils';
    import { _players } from '$lib/stores/table'; 
    import { LoaderCircle } from 'lucide-svelte';
    import { fade } from 'svelte/transition';

    let players_: string = data.tables[tableid-1].players;
    
    async function getPlayers(){
        let $_players = [];
        for ( let i = 0; i < players_.length; i++) {
            const record = await pb.collection("users").getOne(players_[i]);
            _players.update(() => {
                $_players.push(record);
            });
        }
        return $_players;
    }
    let promise = getPlayers();
    onMount(()=> {
        promise = getPlayers();
    });
</script>

<div class="player1 flex flex-row">
{#await promise}
    <span class="loading" transition:fade={{delay: 0, duration: 300}}>
        &nbsp;<LoaderCircle class="inline-flex w-6 left-0 h-6 absolute animate-spin" />
    </span>
{:then _players}
    {#each _players as player}
        {#if player.id != data.user.id}
            <div class="flex avatar">
                <Avatar.Root class="w-20 h-20">
                    <Avatar.Image class="bg-white border-4 rounded-full border-gray-600 cursor-pointer hover:border-gray-400" src={getImageURL(player.id, player.avatar)} alt="avatar" />
                </Avatar.Root>
            </div>
            <div class="flex text-left flex-col">
                <p>{player.username}</p>
                <p>${numberWithCommas(player.balance)}</p>
            </div>
        {/if}
    {/each}
{:catch error}
    <p>{error.message}</p>
{/await}
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
