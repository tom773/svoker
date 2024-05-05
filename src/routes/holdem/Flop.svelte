<script lang="ts">
    
    export let flop: string[] = [];

</script>

<div class="flex flex-row">
    <img src="{flop[0]}.png" alt="River Card" class="card" />
    <img src="{flop[1]}.png" alt="River Card" class="card" />
    <img src="{flop[2]}.png" alt="River Card" class="card" />
</div>

<style>
    .card {
        width: 80px;
        height: 120px;
        margin: 10px;
        border-radius: 0;
    }
</style>
