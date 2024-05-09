<script lang="ts">
    
    export let turn: string[] = [];

</script>

<div class="flex flex-row">
    <img src="{turn[0]}.png" alt="River Card" class="card" />
</div>

<style>
    .card {
        width: 80px;
        height: 120px;
        margin: 10px;
        border-radius: 0;
    }
</style>
