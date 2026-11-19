<script lang="ts">
  import { apiPost } from "$lib/api";
  import { login } from "$lib/auth";
  import { Alert, Button, Card, Field } from "$lib/ui";

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
      const res = await apiPost<{
        status: boolean;
        message: string;
        data: { token: string };
      }>("/api/v1/login", {
        email,
        password,
      });

      const t = res.status ? res.data.token : "";
      if (!t || t === "") throw new Error("No token returned");
      login(t);
    } catch (err) {
      error = err instanceof Error ? err.message : "Login failed";
    } finally {
      loading = false;
    }
  }
</script>

<main class="auth-wrap">
  <div class="auth-box">
    <Card>
      <form class="auth-form" onsubmit={onSubmit}>
        <h1>Login</h1>
        <Field
          label="Email"
          type="email"
          bind:value={email}
          required
          autocomplete="email"
        />
        <Field
          label="Password"
          type="password"
          bind:value={password}
          required
          autocomplete="current-password"
        />
        {#if error}
          <Alert variant="danger">{error}</Alert>
        {/if}
        <Button type="submit" disabled={loading}>
          {loading ? "Signing in..." : "Sign in"}
        </Button>
      </form>
    </Card>
  </div>
</main>

<style>
  .auth-wrap {
    display: grid;
    place-items: center;
    min-height: 100svh;
    padding: var(--space-4);
  }

  .auth-box {
    width: min(360px, 100%);
  }

  .auth-form {
    display: grid;
    gap: var(--space-4);
  }
</style>
