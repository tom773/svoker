<script lang="ts">
    import { getImageURL } from "$lib/utils";
    import * as Avatar from "$lib/components/ui/avatar";
    export let data: any; 
</script>


<div class="flex object-center items-center flex-row" id="chips">
    <div class="flex items-center flex-col object-center avatar">
        <Avatar.Root class="w-20 h-20">
            <Avatar.Image src={getImageURL(data.user.id, data.user.avatar)} class="border-4 rounded-full border-gray-600 cursor-pointer hover:border-gray-400" alt="avatar" />
        </Avatar.Root>
    </div>
    <div class="flex items-center p-2 flex-col">
        <p class="py-1" style="font-size: 18px;">@{data.user.username}</p>
        <p class="py-1" style="font-size: 16px;">$50,000</p>
    </div>
</div>

<style>
</style>
