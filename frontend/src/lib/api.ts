const API_URL = import.meta.env.VITE_API_URL ?? "http://localhost:8000";

export async function apiPost<T>(
  path: string,
  body: unknown,
  token?: string | null
): Promise<T> {
  const headers: Record<string, string> = {
    "Content-Type": "application/json"
  };

  if (token) headers["Authorization"] = `Bearer ${token}`

  const res = await fetch(`${API_URL}${path}`, {
    method: "POST",
    headers: headers,
    body: JSON.stringify(body)
  });

  if (!res.ok) {
    const text = await res.text().catch(() => res.statusText);

    throw new Error(text || `Request failed: ${res.status}`)
  };

  return res.json() as Promise<T>;
}


export function apiGet<T>(path: string, token?: string | null): Promise<T> {
  const headers: Record<string, string> = {};

  if (token) headers["Authorization"] = `Bearer ${token}`

  return fetch(`${API_URL}${path}`, { headers }).then(
    async (res) => {
      if (!res.ok) {
        const text = await res.text().catch(() => res.statusText);

        throw new Error(text || `Request failed: ${res.status}`);
      }

      return res.json() as Promise<T>;
    }
  )
}
