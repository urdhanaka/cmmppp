import { writable } from "svelte/store";

export type Route = "/" | "/login"

export const route = writable<Route>("/")

export function navigate(to: Route) {
  route.set(to);
  // Optional: sync with history
  if (typeof history !== "undefined") {
    history.pushState({}, "", to === "/" ? "/" : to);
  }
}

if (typeof window !== "undefined") {
  window.addEventListener("popstate", () => {
    const path = window.location.pathname as Route;
    if (path === "/" || path === "/login" || path === "/dashboard") {
      route.set(path);
    } else {
      route.set("/login");
    }
  });
}
