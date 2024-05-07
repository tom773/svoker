<script lang="ts">
    import { PlayerType } from "$lib/utils/players";
    import { Avatar } from "@skeletonlabs/skeleton";
    export let player: PlayerType;
</script>


<div class="flex object-center items-center flex-row" id="chips">
    <div class="flex items-center flex-col object-center avatar">
        <Avatar border="border-4 border-gray-600 hover:!border-surface-800"
	            cursor="cursor-pointer" src={player.avatar} width="w-20" rounded="rounded-full" alt="avatar" />
    </div>
    <div class="flex items-center p-2 flex-col">
        <p class="py-1" style="font-size: 18px;">@{player.name}</p>
        <p class="py-1" style="font-size: 16px;">${player.chips}</p>
    </div>
</div>

<style>
</style>
