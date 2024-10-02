import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

export default defineConfig({
  site: "https://weburz.github.io/repoforge",
  base: "repoforge",
  integrations: [
    starlight({
      title: "RepoForge",
      social: {
        github: "https://github.com/Weburz/repoforge",
      },
      sidebar: [
        { label: "Introduction", slug: "introduction" },
        {
          label: "Guides",
          items: [
            { label: "Developer Guide", slug: "guides/development" },
            { label: "Installation Guide", slug: "guides/installation" },
            { label: "Example Guide", slug: "guides/example" },
          ],
        },
        {
          label: "Reference",
          autogenerate: { directory: "reference" },
        },
      ],
    }),
  ],
});
