<script lang="ts">
    import * as Avatar from '$lib/components/ui/avatar';
    export let data: any;
    export let tableid: any;
    import { onMount } from 'svelte';
    import { numberWithCommas, getImageURL } from '$lib/utils';
    import { LoaderCircle } from 'lucide-svelte';
    import { fade } from 'svelte/transition';
    import { playerStore } from '$lib/stores/table';
        
    $: players = [];
    $: playerobj = [];
    onMount(async () => {
        let store = await playerStore(tableid);
        store.subscribe(value =>  {
            players = value;
            test();
        });
    });
    async function test(){
        const res = await fetch(`/api/utils`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({players: players})
        });
        playerobj = await res.json();
    }
</script>
<div class="player1 flex flex-row">
{#if playerobj.length > 0}
    {#each playerobj as player}
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
{:else}    
    <span class="loading" transition:fade={{delay: 0, duration: 300}}>
        &nbsp;<LoaderCircle class="inline-flex w-6 left-0 h-6 absolute animate-spin" />
    </span>
{/if}
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
