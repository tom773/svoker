<script lang="ts">

    import { handtype_, ranks_ } from '$lib/store';

</script>

<div class="test">

    {#if $handtype_ !== ''}
        <h1 id="result" style="visibility: visible;">{$handtype_} of {$ranks_}'s</h1>
    {:else}
        <h1 id="result" style="visibility: hidden;">PH</h1>
    {/if}

</div>
