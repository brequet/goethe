<script>
  import Counter from "$lib/Counter.svelte";
  import { onMount } from "svelte";
  import { ApiService } from "../services/example/exampleService";

  let helloMessage = "";
  let goodbyeMessage = "";

  const apiService = new ApiService();

  onMount(async () => {
    try {
      helloMessage = await apiService.getHello();
      goodbyeMessage = await apiService.getGoodbye();
    } catch (error) {
      console.error("Failed to fetch data:", error);
    }
  });
</script>

<div class="flex flex-col items-center">
  <h1 class="text-4xl font-bold mb-8">Hello</h1>
  <p class="text-xl mb-4">TailwindCSS is installed.</p>
  <p class="text-xl mb-4">
    The <Counter /> component comes from shadcn that is already installed in this
    template.
  </p>
  <p class="text-xl">The svelte-spa-router is used for routing.</p>
  <p>Server says:</p>
  <ul>
    <li>Hello Message: {helloMessage}</li>
    <li>Goodbye Message: {goodbyeMessage}</li>
  </ul>
</div>
