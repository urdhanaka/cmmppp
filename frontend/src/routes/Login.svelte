<script lang="ts">
  import { apiPost } from "$lib/api";

  let email = "";
  let password = "";
  let loading = false;
  let error = "";

  async function onSubmit(e: SubmitEvent) {
    e.preventDefault();
    error = "";
    loading = true;

    try {
      // Backend returns token (see auth handler). Adjust if it returns { token: ... }
      const res = await apiPost<string | { token: string }>("/auth/login", {
        email,
        password,
      });

      const t = typeof res === "string" ? res : res.token;
      if (!t) throw new Error("No token returned");
    } catch (err) {
      error = err instanceof Error ? err.message : "Login failed";
    } finally {
      loading = false;
    }
  }
</script>

<main class="auth-wrap">
  <form class="card" onsubmit={onSubmit}>
    <h1>Login</h1>
    <label>
      Email
      <input type="email" bind:value={email} required autocomplete="email" />
    </label>
    <label>
      Password
      <input
        type="password"
        bind:value={password}
        required
        autocomplete="current-password"
      />
    </label>
    {#if error}
      <p class="error">{error}</p>
    {/if}
    <button disabled={loading}>{loading ? "Signing in..." : "Sign in"}</button>
  </form>
</main>

<style>
  .auth-wrap {
    display: grid;
    place-items: center;
    min-height: 100vh;
    padding: 1rem;
  }
  .card {
    width: min(360px, 100%);
    display: grid;
    gap: 1rem;
    padding: 2rem;
    border-radius: 12px;
    background: #1a1a1a;
    border: 1px solid #333;
  }
  h1 {
    margin: 0;
    font-size: 1.5rem;
  }
  label {
    display: grid;
    gap: 0.5rem;
    font-size: 0.9rem;
  }
  input {
    padding: 0.6rem 0.7rem;
    border-radius: 8px;
    border: 1px solid #444;
    background: #0f0f0f;
    color: inherit;
  }
  button {
    padding: 0.7rem;
    border-radius: 8px;
    border: 0;
    background: #646cff;
    color: white;
    cursor: pointer;
  }
  button:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }
  .error {
    color: #ff6b6b;
    margin: 0;
    font-size: 0.9rem;
  }
</style>
