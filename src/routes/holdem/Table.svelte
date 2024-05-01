<script lang="ts">
    import Card from './Card.svelte';
    
    let cards: string[] = ['aceh', 'kingh'];
    let drawn: string[] = [];

    function dealCards() {
        drawn = [];
        for (let i = 0; i < 5; i++) {
            drawn.push(draw());
        }
    }

    function draw() {
        return cards[Math.floor(Math.random() * cards.length)];
    }

</script>
<div class="flex flex-col items-center justify-center">

    <div class="table flex flex-col items-center justify-center">
        <div class="cardset items-center justify-center flex flex-row">
            {#each drawn as card}
                <Card drawn={card} />
            {/each}
        </div>
    </div>

    <div class="deal">
        <button on:click={dealCards} type="button" class="btn my-5 variant-filled">Deal</button>
    </div>
</div>
<style>
   
.cardset{
    backdrop-filter: blur(10px);
    height: 100%;
    width: 100%;
}

.table {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 1200px;
    height: 600px;
    background: url('./velv.jpg');
    border-radius: 600px / 300px;
    border: 10px solid #996515;
    box-shadow: 0 0 90px rgba(0, 0, 0, 0.8) inset, 0 0 45px rgba(0, 0, 0, 0.8); 
    
}

</style>
