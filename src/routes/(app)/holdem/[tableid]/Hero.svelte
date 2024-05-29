<script lang="ts">
    import ArrowLeft from "lucide-svelte/icons/arrow-left";
    import { Button } from '$lib/components/ui/button/index';
    export let tableid: string;
    export let data: any;
    let table = data.tables.find((table: any) => table.id === tableid);
    let tnum = table.tnum;
</script>

<div class="landing items-center">
    <Button href="/holdem" variant="ghost" class="text-3xl"><ArrowLeft /></Button>
    
    <p class="text-2xl font-bold">Table {tnum}</p>
</div>

<style>
    .landing{
        z-index: 100;
        display: flex;
        font-size: 2rem;
    }
</style>
