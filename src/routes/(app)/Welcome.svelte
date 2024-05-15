<script lang="ts">
   import { Button } from "$lib/components/ui/button/index"
</script>

<img class="ldgcard" src="cardlanding.png" alt="arrow">
<div class="welcome">
    <div class="flex items-baseline flex-row">
        <p class="bottom-0 text-5xl">Welcome</p>
    </div>
    <p class="font-bold text-6xl">To The Show</p>
    <div class="flex py-5 items-center flex-row">
        <Button href="/holdem" class="mx-2">Holdem'</Button>
        <Button href="/omaha" class="mx-2">Omaha</Button>
        <Button class="mx-2">Blackjack</Button>
        <Button class="mx-2">Roulette</Button>
    </div>
</div>

<style>
    
    .ldgcard {
        margin-bottom: 8%;
        margin-left: 8%;
        position: absolute;
        width: 300px;
        height: 400px;
        opacity: 0.3;
        rotate: 22deg;
        object-fit: default;
        border-radius: 10px;
    }
</style>
