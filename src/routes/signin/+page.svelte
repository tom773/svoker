<script lang="ts">
    import ArrowLeft from "lucide-svelte/icons/arrow-left";
	import UserAuthForm from "./(components)/user-auth-form.svelte";
	import { Button } from "$lib/components/ui/button/index.js";
    
    function forgot() {
        alert("Sucked in Dory! We haven't set up SMTP yet, so make a new account!");
    }

</script>

<div class="container relative hidden h-full flex-col items-center justify-center md:grid lg:max-w-none lg:grid-cols-2 lg:px-0">
	<Button href="/signup" variant="ghost" style="z-index: 100;" class="absolute text-white right-4 w-20 top-4 md:right-8 md:top-8"> Sign Up </Button>
	<Button href="/" variant="ghost" style="z-index: 100;" class="absolute text-white mx-4 left-1/2 w-20 top-4 md:right-8 md:top-8"> <ArrowLeft/> </Button>
	<div class="relative hidden h-full flex-col bg-muted text-white dark:border-r lg:flex">
		<div class="absolute inset-0 bg-cover" style="background-image: url(singup.webp);"/>
		<div class="relative justify-center h-full w-full backdrop-brightness-50 backdrop-blur z-20 flex items-center text-6xl font-bold">
		    Svoker
            <div class="inline-flex mb-8 mx-5">
                <img src="AH.png" alt="svoker" class="-rotate-12 w-30 h-40" />
		        <img src="KH.png" alt="svoker" class="rotate-12 w-30 h-40" />
            </div>
        </div>
    <!--background: linear-gradient(-65deg, rgb(249, 223, 220), rgb(10, 129, 171), rgb(12, 66, 113), rgb(0, 0, 0));-->
	</div>
    
	<div class="flex lg:p-8 items-center h-full" style="background: linear-gradient(-65deg, rgb(240, 165, 0), rgb(207, 117, 0),  rgb(1, 1, 1));">
		<div class="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
			<div class="flex flex-col space-y-2 text-center">
				<h1 class="text-2xl text-gray-200 font-semibold tracking-tight">Sign In</h1>
				<p class="text-sm text-gray-200">
				    You need to be signed in to play, champ!
                </p>
			</div>
			<UserAuthForm />
            <p class="text-gray-200 text-center">Forgot your password? <a href="#" on:click={forgot} class="text-blue-400" >Click here</a></p>
		</div>
	</div>
</div>
