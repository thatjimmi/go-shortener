<script>
  import Router from "./routes/index.svelte";
  import Tailwind from "./lib/Tailwind.svelte";
</script>

<Tailwind />
<main>
  <Router />
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }
</style>
