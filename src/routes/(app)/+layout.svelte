<script lang="ts">
    import { Button } from "$lib/components/ui/button/index";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index";
    import * as Avatar from "$lib/components/ui/avatar/index";
    import { onMount } from "svelte";
    export let data;
    function getAvatar(avatar: any, id: string, size='0x0') {

        return `http://localhost:8090/api/files/users/${id}/${avatar}?thumb=${size}`;
    }
    let active = false;
    onMount(() => {
        document.onmousemove = () => {
            active = true;
            setTimeout(() => {
                active = false;
            }, 1200000);
        };
    });
</script>

<svelte:head>
    <link rel="preconnect" href="https://rsms.me/">
    <link rel="stylesheet" href="https://rsms.me/inter/inter.css">
</svelte:head>

<div class="auth">
    {#if !data.user}
        <Button href="/signin">Sign In</Button>
    {:else }
        <DropdownMenu.Root>
            <DropdownMenu.Trigger>
                <Avatar.Root class="w-20 h-20">
                    <Avatar.Image src={getAvatar(data.user.avatar, data.user.id)} class="border-4 rounded-full border-black cursor-pointer hover:border-gray-600" alt="lol" />
                    <Avatar.Fallback></Avatar.Fallback>
                </Avatar.Root>
                {#if active}
                    <svg viewBox="0 0 20 20" style="height: 15px; z-index: 102; margin-left: 8px; margin-top: -15px;" class="border border-green rounded-full shadow" xmlns="http://www.w3.org/2000/svg">
                        <circle fill="#2bff44" cx="10" cy="10" r="10" />
                    </svg>
                {/if}
            </DropdownMenu.Trigger>
            <DropdownMenu.Content>
                <DropdownMenu.Group>
                    <DropdownMenu.Label>My Account</DropdownMenu.Label>
                    <DropdownMenu.Separator />
                    <DropdownMenu.Item class="cursor-pointer" href="/settings">Profile</DropdownMenu.Item>
                    <DropdownMenu.Item class="cursor-pointer">Billing</DropdownMenu.Item>
                    <DropdownMenu.Item class="cursor-pointer" href="/settings">Settings</DropdownMenu.Item>
                    <DropdownMenu.Item>
                        <form action="/signout" method="POST">
                            <Button type="submit">Sign Out</Button>
                        </form>
                    </DropdownMenu.Item>
                </DropdownMenu.Group>
            </DropdownMenu.Content>
        </DropdownMenu.Root>
    {/if}
</div>
<slot />
<style>
    .auth{
        text-align: right;
        right: 0;
        width: 10%;
        padding: 1rem;
        color: #f3f3f3;
        z-index: 1000;
        position: absolute;
    }
</style>
