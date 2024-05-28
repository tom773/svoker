<script lang="ts">
    import { Button } from "$lib/components/ui/button/index";
    import * as Card from "$lib/components/ui/card/index";
    import { Input } from "$lib/components/ui/input/index";
    export let data: any;
</script>
<div class="grid grid-cols-2 gap-4">
    <Card.Root>
      <Card.Header>
        <Card.Title>Get Help</Card.Title>
        <Card.Description>
            Contact our unfreindly CS team if you're struggling with any part of the site!
        </Card.Description>
      </Card.Header>
      <form method="" action="" class="items-center">
          <Card.Content class="grid-rows-2 border-t px-6 py-5 items-center">
                <div class="grid grid-cols-2 gap-4">
                    <div class="flex flex-col w-full">
                        <label for="name" class="font-bold text-xl mb-1">Name</label>
                        <Input id="name" name="name" type="text" placeholder="John Doe" required/>
                    </div>
                    <div class="flex flex-col w-full">
                        <label for="email" class="font-bold text-xl mb-1">Email</label>
                        <Input id="email" name="email" type="email" placeholder="assman91@gmail.com" />
                    </div>
                </div>
                <div class="flex flex-col w-full">
                    <label for="message" class="font-bold text-xl mb-1">Message</label>
                    <textarea id="message" name="message" placeholder="Your message here..." class="w-full p-2 border rounded-md" rows="4" required></textarea>
                </div>
            </Card.Content>
          <Card.Footer class="border-t px-6 py-5">
            <Button type="submit">Submit</Button>
          </Card.Footer>
      </form>
    </Card.Root>
</div>


