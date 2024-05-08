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
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">{$ranks_[0]} High</p>
        {:else if $handtype_==='Pair'}<!-- Implemented -->
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">Pair of {$ranks_}'s</p>
        {:else if $handtype_==='TwoPair'}<!-- Implemented -->
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">Two Pair: {$ranks_[0]}'s & {$ranks_[1]}'s</p>
        {:else if $handtype_==='ThreeOfAKind'}<!-- Implemented -->
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">Trips: {$ranks_}'s</p>
        {:else if $handtype_==='Straight'}
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">Straight: {$ranks_[0]} High</p>
        {:else if $handtype_==='Flush'}<!-- Implemented -->
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">Flush: {$suits_}</p>
        {:else if $handtype_==='FullHouse'}<!-- Implemented -->
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">Full House: {$ranks_[0]}'s over {$ranks_[1]}'s</p>
        {:else if $handtype_==='Quads'}<!-- Implemented -->
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">Quads: {$ranks_}'s</p>
        {:else if $handtype_==='StraightFlush'}
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">Straight Flush</p>
        {:else if $handtype_==='RoyalFlush'}
            <p class="font-bold" id="result" style="font-size: 18px; visibility: visible;">Royal Flush</p>
        {:else}
            <p class="font-bold" id="result" style="font-size: 18px; visibility: hidden;">PH</p>
        {/if}

    {:else}
        <p class="font-bold" id="result" style="font-size: 18px; visibility: hidden;">PH</p>
    {/if}

</div>
