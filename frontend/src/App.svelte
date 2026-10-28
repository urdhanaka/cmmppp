<script>
  import { route } from "$lib/router";
  import Login from "./routes/Login.svelte";
</script>

{#if $route === "/login" || $route === "/"}
  <Login />
{/if}
