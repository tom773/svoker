<script lang="ts">
    import { PlayerType } from '$lib/utils/players'; 
    export let you: PlayerType;

</script>


<div>

    <h1>{you.name}: ${you.chips}</h1>

</div>
