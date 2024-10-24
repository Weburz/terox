import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";
import starlightLinksValidatorPlugin from "starlight-links-validator";

export default defineConfig({
  site: "https://weburz.github.io/repoforge",
  base: "repoforge",
  integrations: [
    starlight({
      title: "RepoForge",
      description: "Scaffold your projects through the power of automation!",
      editLink: {
        baseUrl: "https://github.com/Weburz/repoforge/edit/main/docs",
      },
      social: {
        github: "https://github.com/Weburz/repoforge",
        discord: "https://discord.gg/QeYqwyxBhR",
        email: "mailto:contact@weburz.com",
        facebook: "https://www.facebook.com/Weburz",
        instagram: "https://www.instagram.com/weburzit",
        linkedin: "https://www.linkedin.com/company/weburz",
        youtube: "https://www.youtube.com/@Weburz",
        twitter: "https://x.com/weburz",
      },
      lastUpdated: true,
      sidebar: [
        {
          label: "Introduction",
          slug: "introduction",
        },
        {
          label: "Usage Guide",
          items: [
            { label: "Installation Guide", slug: "usage-guide/installation" },
            {
              label: "Command-Line Interface (CLI) Reference",
              slug: "usage-guide/cli-reference",
            },
          ],
        },
        {
          label: "Contribution & Development Guide",
          items: [
            {
              label: "Development Guidelines",
              slug: "development-guide/development",
            },
            {
              label: "Software Requirements Specifications (SRS)",
              slug: "development-guide/spec-sheet",
            },
          ],
        },
      ],
      plugins: [starlightLinksValidatorPlugin()],
    }),
  ],
});
