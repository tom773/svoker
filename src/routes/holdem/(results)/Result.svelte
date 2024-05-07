<script lang="ts">
    
    import { handtype_, ranks_, suits_ } from '$lib/store';
    let ranks: string[] = [];
    ranks_.subscribe((value) => {
        ranks = value;
    });
    
</script>

<div class="test">

    {#if $handtype_ !== ''}
        {#if $handtype_==='HighCard'}<!-- Implemented -->
            <h1 id="result" style="visibility: visible;">{$ranks_[0]} High</h1>
        {:else if $handtype_==='Pair'}<!-- Implemented -->
            <h1 id="result" style="visibility: visible;">Pair of {$ranks_}'s</h1>
        {:else if $handtype_==='TwoPair'}<!-- Implemented -->
            <h1 id="result" style="visibility: visible;">Two Pair: {$ranks_[0]}'s & {$ranks_[1]}'s</h1>
        {:else if $handtype_==='ThreeOfAKind'}<!-- Implemented -->
            <h1 id="result" style="visibility: visible;">Trips: {$ranks_}'s</h1>
        {:else if $handtype_==='Straight'}
            <h1 id="result" style="visibility: visible;">Straight: {$ranks_[0]} High</h1>
        {:else if $handtype_==='Flush'}<!-- Implemented -->
            <h1 id="result" style="visibility: visible;">Flush: {$ranks_} High {$suits_}</h1>
        {:else if $handtype_==='FullHouse'}<!-- Implemented -->
            <h1 id="result" style="visibility: visible;">Full House: {$ranks_[0]}'s over {$ranks_[1]}'s</h1>
        {:else if $handtype_==='Quads'}<!-- Implemented -->
            <h1 id="result" style="visibility: visible;">Quads: {$ranks_}'s</h1>
        {:else if $handtype_==='StraightFlush'}
            <h1 id="result" style="visibility: visible;">Straight Flush: {$ranks_} High</h1>
        {:else if $handtype_==='RoyalFlush'}
            <h1 id="result" style="visibility: visible;">Royal Flush: {$ranks_} High</h1>
        {:else}
            <h1 id="result" style="visibility: hidden;">PH</h1>
        {/if}

    {:else}
        <h1 id="result" style="visibility: hidden;">PH</h1>
    {/if}

</div>
