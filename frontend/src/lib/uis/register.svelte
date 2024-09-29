<script lang="ts">
    import { Button } from "$lib/components/ui/button";
	import Input from "$lib/components/ui/input/input.svelte";
  </script>
   
<div class="flex items-center justify-center min-h-screen w-screen flex-col">
    <h1 class="font-semibold text-2xl ">Welcome!</h1>
    <h2 class="text-lg">Please register an account</h2>
    <form class="space-y-2 mt-4">
    <Input placeholder="Username" /> 
    <Input placeholder="Password" /> 
    <Button class="w-full" variant="outline">Register</Button>
</form>
</div>