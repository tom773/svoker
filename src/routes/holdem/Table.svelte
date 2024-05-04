<script lang="ts">
    import Card from './Card.svelte';
    
    let cards: string[] = ['aceh', 'kingh', 'queenh', 'jackh', 'nineh', 'eighth', 'sevenh', 'sixh', 'fiveh', 'fourh', 'threeh', 'twoh'];
    let drawn: string[] = [];

    function dealCards() {
        drawn = [];
        for (let i = 0; i < 2; i++) {
            drawn.push(draw());
        }
    }

    function draw() {
        return cards[Math.floor(Math.random() * cards.length)];
    }

</script>
<div class="flex w-full h-full items-center justify-center flex-col">
    <div class="flex flex-col items-center justify-center">

        <div class="table flex flex-col items-center justify-center">
        </div>

        <div class="deal">
            <button on:click={dealCards} type="button" class="btn my-5 variant-filled">Deal</button>
        </div>
    </div>

    <div class="cardset w-1/4 items-center justify-start flex flex-row">
        {#each drawn as card}
            <Card drawn={card} />
        {/each}
    </div>
</div>
<style>
   
.cardset{
    position: absolute;
    bottom: 0;
    left: 0;
    padding: 10px;
}

.table {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 70vw;
    height: 80vh;
    background: url('./velv.jpg');
    border-radius: 500px / 400px;
    border: 10px solid #996515;
    box-shadow: 0 0 90px rgba(0, 0, 0, 0.8) inset, 0 0 45px rgba(0, 0, 0, 0.8); 
    
}

</style>
