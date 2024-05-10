<script lang="ts">
    import ArrowLeft from "lucide-svelte/icons/arrow-left";
    import { Button } from '$lib/components/ui/button/index';
    export let tnum: number;
</script>

<div class="landing">
    <Button href="/holdem" variant="ghost" class="text-3xl"><ArrowLeft /></Button>
    
    <p class="text-3xl">Welcome to Table {tnum}!</p>
</div>

<style>
    .landing{
        z-index: 100;
        display: flex;
        font-size: 2rem;
    }
</style>
