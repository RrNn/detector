import { createWebHistory, createRouter } from "vue-router";
import Home from "../components/home";
import Url from "../components/url";
import About from "../components/about";
import Landing from "../components/landing";
import NotFound from "../components/notfound";

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
  {
    path: "/urls/:id",
    name: "url",
    component: Url,
  },
  {
    path: "/about",
    name: "about",
    component: About,
  },
  {
    path: "/logout",
    name: "logout",
    component: Landing,
  },
  {
    path: "/:catchAll(.*)",
    component: NotFound,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
