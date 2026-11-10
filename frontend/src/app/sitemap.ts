import { MetadataRoute } from "next";
import { getSiteUrl } from "@/lib/site-url";

export const dynamic = "force-dynamic";

export default function sitemap(): MetadataRoute.Sitemap {
  // One canonical page. Hash anchors (#services, #experience) and per-deploy
  // Vercel hosts are rejected by Google as "URL not allowed". The CV is private.
  return [
    {
      url: getSiteUrl(),
      lastModified: new Date(),
      changeFrequency: "weekly",
      priority: 1,
    },
  ];
}
