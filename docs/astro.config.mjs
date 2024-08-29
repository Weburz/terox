import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

export default defineConfig({
  integrations: [
    starlight({
      title: "RepoForge",
      social: {
        github: "https://github.com/Weburz/repoforge",
      },
      sidebar: [
        { label: "Introduction", slug: "guides/introduction" },
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
