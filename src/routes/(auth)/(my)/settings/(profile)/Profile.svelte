<script lang="ts">
    import { getImageURL } from "$lib/utils";
    import { Button } from "$lib/components/ui/button/index";
    import * as Card from "$lib/components/ui/card/index";
    import { Input } from "$lib/components/ui/input/index";
    export let data: any;
    import { invalidateAll } from "$app/navigation";
    import * as Avatar from "$lib/components/ui/avatar/index";
    import { Pencil } from "lucide-svelte";
    import { onMount } from "svelte";
    import { LoaderCircle } from "lucide-svelte";
    import { applyAction, enhance } from "$app/forms";
    import { fade } from "svelte/transition";
    import { CircleCheckBig } from 'lucide-svelte';
    
    export let avatarUp = false;
    export let avatarInput: any;
    let loading: boolean;

    $: loading = false;
    
    onMount(() => {
        avatarInput = document.getElementById("avatar");
        avatarInput.addEventListener("change", () => {
            avatarUp = true;
        });
    });

    const preview = (event: any) => {
        const target = event.target;
        const files = target.files;

        if (files.length > 0) {
            const src = URL.createObjectURL(files[0]);
            const preview = document.getElementById('avatar-preview') as HTMLImageElement;
            preview.src = src;
        }
    };

    const subby = () => {
        loading = true;
        return async ({ result }) => {
            switch (result.type){
                case "success":
                    invalidateAll();
                    break;
                default:
                    await applyAction(result);
            }

            setTimeout(() => {
                loading = false;
            }, 2000);
            setTimeout(() => {
                document.getElementById('tick').classList.remove("hidden");
            }, 2300);
            setTimeout(() => {
                document.getElementById('tick').classList.add("hidden");
                avatarUp = false;
            }, 5300);
        };
        
    };
    

</script>
<div class="grid grid-cols-2 gap-4">
    <Card.Root>
      <Card.Header>
        <Card.Title>Display Name</Card.Title>
        <Card.Description>
          Used to identify you by an alias such as "TheAssMan".<br><strong>Current:</strong> {data?.user?.username}
        </Card.Description>
      </Card.Header>
      <form method="POST" action="?/username">
          <Card.Content>
              <Input name="username" id="username" type="username" placeholder="username" />
          </Card.Content>
          <Card.Footer class="border-t px-6 py-4">
            <Button type="submit">Save</Button>
          </Card.Footer>
      </form>
    </Card.Root>
    <Card.Root>
      <Card.Header>
        <Card.Title>Update Email</Card.Title>
        <Card.Description>
          Update your email address.<br><strong>Current:</strong> {data?.user?.email}
        </Card.Description>
      </Card.Header>
    <form method="POST" action="?/email">
      <Card.Content>
          <Input name="email" id="email" type="email" placeholder="assman@kramerica.com" />
      </Card.Content>
      <Card.Footer class="border-t px-6 py-4">
        <Button type="submit">Send Confirmation Email</Button>
      </Card.Footer>
    </form>
    </Card.Root>
    <Card.Root>
      <Card.Header>
        <Card.Title>Update Avatar</Card.Title>
      </Card.Header>
    <form method="POST" use:enhance={subby} enctype="multipart/form-data" action="?/avatar">
      <Card.Content>
            <label for="avatar" class="text-sm items-center justify-center text-center ">
              <Avatar.Root class="w-full h-24 items-center justify-center hover:opacity-75 hover:cursor-pointer hover:brightness-75">
                    <Pencil style="z-index:100" fill="#000000" class="absolute w-8 h-8 hidden hover:block" />
                    <Avatar.Image id="avatar-preview" class="w-24 h-24 border-2 border-black hover:border-gray-600 rounded-full" src={getImageURL(data?.user.id, data?.user?.avatar)} />
              </Avatar.Root>
              <p class="hover:cursor-pointer">Click to change</p>
            </label>
          <Input on:change={preview} style="display: none" name="avatar" accept="image/*" id="avatar" type="file" placeholder="assman@kramerica.com" />
      </Card.Content>
      <Card.Footer class="border-t px-6 py-4">
            {#if avatarUp}
                <Button type="submit" id="upload" class="inline-flex items-center justify-center">Upload
                    {#if loading}
                        <span class="loading" transition:fade={{delay: 0, duration: 300}}>
                            &nbsp;<LoaderCircle class="inline-flex w-6 h-6 relative animate-spin" />
                        </span>
                    {/if}
                    <span class="success" transition:fade={{delay: 0, duration: 500}}>
                        &nbsp;<CircleCheckBig id="tick" fill="green" class="inline-flex relative w-6 h-6 hidden" />
                    </span>
                </Button>
            {:else}
                <Button variant="disabled">Upload</Button>
            {/if}
      </Card.Footer>
    </form>
    </Card.Root>
</div>

