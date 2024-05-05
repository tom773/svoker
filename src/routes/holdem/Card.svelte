<script lang="ts">
    export let drawn: string;
</script>

<div class="usercard">
    <img src="./{drawn}.png" alt="card">
</div>

<style>
    .usercard{
       width: 100px; 
       height: 150px;
       rotate: var(--rot);
    }
</style>
