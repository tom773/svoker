<script lang="ts">
	import { Button } from "$lib/components/ui/button/index";
	import { Input } from "$lib/components/ui/input/index";
	import { Label } from "$lib/components/ui/label/index";
	import { cn } from "$lib/utils.js";
    import "../../../../main.css";

	let className: string | undefined | null = undefined;
	export { className as class };
    
</script>

<div class={cn("grid gap-6", className)} {...$$restProps}>
	<form action="?/login" method="POST">
		<div class="grid gap-2">
			<div class="grid gap-3">
				<Label class="text-gray-200" for="email">Email</Label>
				<Input
					id="email"
                    name="email"
					placeholder="bruce@fishaddictsanonymous.com"
					type="email"
					autocapitalize="none"
					autocomplete="email"
					autocorrect="off"
				/>
				<Label class="text-gray-200" for="password">Password</Label>
				<Input
					id="password"
                    name="password"
					placeholder="password"
					type="password"
					autocapitalize="none"
					autocomplete="email"
					autocorrect="off"
				/>
			</div>
			<Button type="submit">
				Sign Up 
			</Button>
		</div>
	</form>
</div>
