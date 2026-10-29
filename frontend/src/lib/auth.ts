import { writable } from "svelte/store";

// key of JWT token that will be saved on localStorage
const TOKEN_KEY = "cmp_token";

export const token = writable<string | null>(getStoredToken())
export const isAuthed = writable<boolean>(!!getStoredToken())

function getStoredToken() {
  if (typeof localStorage === "undefined")
    return null;

  return localStorage.getItem(TOKEN_KEY);
}

token.subscribe((t) => {
  if (typeof localStorage === "undefined") return;

  if (t) {
    localStorage.setItem(TOKEN_KEY, t)
    isAuthed.set(true);
  } else {
    localStorage.removeItem(TOKEN_KEY)
    isAuthed.set(false);
  }
})

export function logout() {
  token.set(null)
}

export function login(t: string) {
  token.set(t)
}
