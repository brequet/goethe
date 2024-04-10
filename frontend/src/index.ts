import About from "./routes/About.svelte";
import Home from "./routes/Home.svelte";
import NotFound from "./routes/NotFound.svelte";

export default {
  "/": Home,
  "/about": About,
  "*": NotFound,
};
