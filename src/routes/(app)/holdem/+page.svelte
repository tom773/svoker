<script lang="ts">
    import * as Table from '$lib/components/ui/table';
    import { Button } from '$lib/components/ui/button';
    import { serializeNonPOJOs } from '$lib/utils';
    export let data;    
</script>

<main>
    <div class="games w-1/2 mt-10 rounded-lg bg-black text-white">
        <Table.Root>
            <Table.Caption class="text-white">Public Tables</Table.Caption>
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
                            <Button href="./game" class="bg-green-500 hover:bg-green-700 active:bg-green-800">Join</Button>
                        </Table.Cell>
                    </Table.Row>
                {/each}
            </Table.Body>
        </Table.Root>
    </div>

</main>

<style>
    
    main {
        display: flex;
        justify-content: center;
        align-items: center;
    }

</style>
