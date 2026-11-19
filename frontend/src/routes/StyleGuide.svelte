<script lang="ts">
  import { Alert, Button, Card, Field, Modal } from "$lib/ui";

  const colors = [
    "bg",
    "surface",
    "border",
    "input-bg",
    "text",
    "text-muted",
    "focus",
    "primary",
    "primary-hover",
    "primary-text",
    "danger",
    "success",
    "info",
  ];
  const spaces = [1, 2, 3, 4, 6, 8];
  const radii = ["sm", "md", "lg"];

  let modalOpen = $state(false);
  let sample = $state("");
  let sampleError = $derived(
    sample.length > 0 && sample.length < 3 ? "Minimum 3 characters" : undefined,
  );
</script>

<main class="sg">
  <header>
    <h1>Style guide</h1>
    <p class="muted">Reference for tokens and ui/ components. Dev only.</p>
  </header>

  <section>
    <h2>Colors</h2>
    <div class="grid">
      {#each colors as name}
        <div class="swatch">
          <div class="chip" style="background: var(--color-{name})"></div>
          <code>--color-{name}</code>
        </div>
      {/each}
    </div>
  </section>

  <section>
    <h2>Typography</h2>
    <h1>Heading 1</h1>
    <h2>Heading 2</h2>
    <p>Body text. The quick brown fox jumps over the lazy dog.</p>
    <p class="muted">Muted text, used for secondary information.</p>
    <p><code>inline code</code></p>
  </section>

  <section>
    <h2>Spacing</h2>
    {#each spaces as n}
      <div class="row">
        <code>--space-{n}</code>
        <div class="bar" style="width: var(--space-{n})"></div>
      </div>
    {/each}
  </section>

  <section>
    <h2>Radius</h2>
    <div class="row wrap">
      {#each radii as r}
        <div class="radius" style="border-radius: var(--radius-{r})">
          <code>--radius-{r}</code>
        </div>
      {/each}
    </div>
  </section>

  <section>
    <h2>Button</h2>
    <div class="row wrap">
      <Button>Primary</Button>
      <Button variant="ghost">Ghost</Button>
      <Button variant="danger">Danger</Button>
      <Button disabled>Disabled</Button>
    </div>
  </section>

  <section>
    <h2>Field</h2>
    <Card>
      <Field label="Default" placeholder="Type here" />
      <Field label="With error" value="ab" error="Minimum 3 characters" />
      <Field label="Disabled" value="Not editable" disabled />
      <Field
        label="Bound (validates on input)"
        bind:value={sample}
        error={sampleError}
      />
    </Card>
  </section>

  <section>
    <h2>Alert</h2>
    <div class="stack">
      <Alert variant="info">Info: something to know.</Alert>
      <Alert variant="success">Success: action completed.</Alert>
      <Alert variant="danger">Danger: something failed.</Alert>
    </div>
  </section>

  <section>
    <h2>Card</h2>
    <Card>
      <h2>Card title</h2>
      <p class="muted">Card content goes here.</p>
    </Card>
  </section>

  <section>
    <h2>Modal</h2>
    <Button onclick={() => (modalOpen = true)}>Open modal</Button>
    <Modal bind:open={modalOpen} title="Example modal">
      <p>Press Escape or the × button to close.</p>
    </Modal>
  </section>
</main>

<style>
  .sg {
    max-width: 960px;
    margin: 0 auto;
    padding: var(--space-8) var(--space-4);
    display: grid;
    gap: var(--space-8);
  }

  section {
    display: grid;
    gap: var(--space-4);
    align-items: start;
  }

  .muted {
    color: var(--color-text-muted);
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
    gap: var(--space-4);
  }

  .swatch {
    display: grid;
    gap: var(--space-2);
    font-size: var(--font-sm);
  }

  .chip {
    height: 48px;
    border: 1px solid var(--color-border);
    border-radius: var(--radius-md);
  }

  .row {
    display: flex;
    align-items: center;
    gap: var(--space-4);
  }

  .wrap {
    flex-wrap: wrap;
  }

  .bar {
    height: 12px;
    background: var(--color-primary);
    border-radius: var(--radius-sm);
  }

  .radius {
    width: 96px;
    height: 96px;
    display: grid;
    place-items: center;
    border: 1px solid var(--color-border);
    background: var(--color-surface);
    font-size: var(--font-sm);
  }

  .stack {
    display: grid;
    gap: var(--space-3);
  }
</style>
