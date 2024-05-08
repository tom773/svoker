<script lang="ts">
    import { fade } from 'svelte/transition';
    import Stack from './Stack.svelte';
    import Flop from './Flop.svelte';
    import Turn from './Turn.svelte';
    import River from './River.svelte';
    import Player from './Player.svelte';
    import { players } from '$lib/utils/game';
    import { flop_, turn_, river_ } from '$lib/store';
    import ActionBar from './ActionBar.svelte';

</script>
<div class="flex w-full h-full items-center justify-center flex-col">
    <div class="flex flex-col items-center justify-center">

        <div class="table flex flex-col items-center justify-center">
            <div class="topplayers absolute">
               {#each players as player}
                    <Player name={player.name} chips={player.chips} avatar={player.avatar}/>
                {/each} 
            </div>
            <div class="flex tabcards flex-row">
                <Stack />
                {#key $flop_}
                    <div in:fade={{delay: 200, duration: 200}} id="flop" style="">
                        <Flop flop={$flop_} />
                    </div>
                {/key}
                {#key $turn_}
                    <div in:fade={{delay: 200, duration: 200}} id="turn" style="">
                        <Turn turn={$turn_} />
                    </div>
                {/key}
                {#key $river_}
                    <div in:fade={{delay: 200, duration: 200}} id="river" style="">
                        <River river={$river_} />
                    </div>
                {/key}
            </div>
        </div>

    </div>
    <ActionBar /> 
</div>
<style>
.tabcards{
    padding: 40px;
    border-radius: 500px / 400px;
    border: 4px solid yellow;
}

.table {
    margin-bottom: 10%;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 70vw;
    height: 70vh;
    background: url('./velv.jpg');
    border-radius: 500px / 400px;
    border: 10px solid #481E14;
    box-shadow: 0 0 90px rgba(0, 0, 0, 0.9) inset, 0 0 90px rgba(0, 0, 0, 1); 
    
}

.topplayers {
    display: flex;
    top: 5px;
}
</style>
