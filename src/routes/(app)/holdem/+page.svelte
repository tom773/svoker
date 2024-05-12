<script lang="ts">
    import * as Table from '$lib/components/ui/table';
    import { Button } from '$lib/components/ui/button';
    import { serializeNonPOJOs } from '$lib/utils';
    import {ArrowLeft} from "lucide-svelte";
    //import { tables } from '$lib/stores/table';
    
    export let data;
    for (let table of data.tables) {
        if (table.currentplayers === table.maxplayers) {
            table.full = true;
        }
        if (table.players.includes(data.user.id)) {
            table.joined = true;
        }
    }
    
</script>

<div class="back m-4">
    <Button href="/" variant="ghost" class="text-3xl"><ArrowLeft /></Button>
</div>
<main>
    <div class="games w-1/2 mt-10 rounded-lg bg-gray-800 text-white">
        <Table.Root>
            <Table.Caption class="text-white my-4">Public Tables</Table.Caption>
                <Table.Header class="py-2">
                    <Table.Row class="py-2">
                        <Table.Head class="text-xl font-bold text-white w-[200px]">Table Number</Table.Head>
                        <Table.Head class="text-xl font-bold text-white text-white">Blinds</Table.Head>
                        <Table.Head class="text-xl font-bold text-white text-white">Active Players</Table.Head>
                        <Table.Head class="text-xl font-bold text-white text-white"></Table.Head>
                    </Table.Row>
                </Table.Header>
            <Table.Body>
                {#each serializeNonPOJOs(data?.tables) as table}
                    <Table.Row>
                        <Table.Cell class="font-medium">{table.tnum}</Table.Cell>
                        <Table.Cell>{table.blinds}</Table.Cell>
                        <Table.Cell>{table.currentplayers}/{table.maxplayers}</Table.Cell>
                        <Table.Cell class="text-right">
                            <div class="w-32 inline-flex justify-end">
                                {#if !table.full && !table.joined}
                                <form method="POST" action="?/addToTable">
                                    <input type="hidden" name="table" value="{table.id}">
                                    <input type="hidden" name="tnum" value="{table.tnum}">
                                    <Button type="submit" class="bg-green-500 hover:bg-green-700 active:bg-green-800">Join</Button>
                                </form>
                                {:else if table.joined}
                                    <Button href="/holdem/{table.tnum}" class="bg-blue-500 hover:bg-blue-700">Return</Button>
                                    <form method="POST" action="?/removeFromTable" class="">
                                        <input type="hidden" name="table_" value="{table.id}">
                                        <Button type="submit" class="bg-red-500 ml-2 hover:bg-red-700">Leave</Button>
                                    </form>
                                {:else}
                                    <Button class="bg-gray-500 cursor-not-allowed">Full</Button>
                                {/if}
                            </div>
                        </Table.Cell>
                    </Table.Row>
                {/each}
            </Table.Body>
        </Table.Root>
    </div>

</main>

<style>

    .back{
        position: absolute;
        z-index: 100;
        display: flex;
        font-size: 2rem;
    }
     
    main {
        display: flex;
        justify-content: center;
        align-items: center;
    }

</style>
