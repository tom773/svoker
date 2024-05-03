<script lang="ts">
    import {enhance} from "$app/forms";
    import '../../main.css';
</script>
<div class="flex flex-col justify-center items-center h-full">
    <div class="formLog rounded-lg variant-glass-surface w-1/3">
        <form method="post" use:enhance>
            <label class="label">
                <span>Username</span>
                <input name="username" id="username" class="input" type="text" placeholder="Username" />
            </label>
            <label class="label">
                <span>Password</span>
                <input name="password" id="password"  class="input" type="password" placeholder="password" />
            </label>
            <button class="btn variant-filled" type="submit">Sign Up</button>
        </form>
        <div class="flex mt-10 content-baseline">
            <p class="flex text-white text-l">Already have an account?<a href="/signup">&#160Log In!</a></p>
        </div>
    </div>
</div>
<style>
    
    .formLog{
        margin: auto;
        width: 40%;
        padding: 5rem;
        color: #333333;
    }

</style>
