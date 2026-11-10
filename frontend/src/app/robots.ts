import { MetadataRoute } from "next";
import { getSiteUrl } from "@/lib/site-url";

export const dynamic = "force-dynamic";

export default function robots(): MetadataRoute.Robots {
  if (process.env.VERCEL_ENV === "preview") {
    return {
      rules: {
        userAgent: "*",
        disallow: "/",
      },
    };
  }

  return {
    rules: {
      userAgent: "*",
      allow: "/",
      disallow: ["/api/", "/cv/", "/cv-jhon-quirama-es.pdf", "/cv-jhon-quirama-en.pdf"],
    },
    sitemap: `${getSiteUrl()}/sitemap.xml`,
  };
}
