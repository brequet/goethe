<script lang="ts">
  import Router from "svelte-spa-router";
  import routes from ".";
</script>

<main class="flex flex-col items-center justify-center min-h-screen">
  <h1 class="text-4xl font-bold mb-8">Brequet Template</h1>

  <div class="flex flex-row items-center space-x-4 mb-16">
    <a href="#/" class="text-xl text-blue-500 hover:text-blue-700">Hello page</a
    >
    <a href="#/about" class="text-xl text-blue-500 hover:text-blue-700"
      >About page</a
    >
    <a href="#/nothing" class="text-xl text-blue-500 hover:text-blue-700"
      >This page does not exist</a
    >
  </div>

  <div class="flex-grow flex">
    <Router {routes} class="flex-grow" />
  </div>
</main>
