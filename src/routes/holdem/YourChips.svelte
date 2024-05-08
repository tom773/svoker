<script lang="ts">
    import { PlayerType } from "$lib/utils/players";
    import * as Avatar from "$lib/components/ui/avatar";
    export let player: PlayerType;
    let url = player.avatar;
</script>


<div class="flex object-center items-center flex-row" id="chips">
    <div class="flex items-center flex-col object-center avatar">
        <Avatar.Root class="w-20 h-20">
            <Avatar.Image src={url} class="border-4 rounded-full border-gray-600 cursor-pointer hover:border-gray-400" alt="avatar" />
        </Avatar.Root>
    </div>
    <div class="flex items-center p-2 flex-col">
        <p class="py-1" style="font-size: 18px;">@{player.name}</p>
        <p class="py-1" style="font-size: 16px;">${player.chips}</p>
    </div>
</div>

<style>
</style>
