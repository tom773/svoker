<script lang="ts">
    import ArrowLeft from "lucide-svelte/icons/arrow-left";
    import { Button } from '$lib/components/ui/button/index';
</script>

<div class="landing">
    <Button href="/" variant="ghost" class="text-3xl"><ArrowLeft /></Button>
</div>

<style>
    .landing{
        z-index: 100;
        display: flex;
        font-size: 2rem;
    }
</style>
