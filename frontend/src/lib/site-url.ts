const CANONICAL_SITE_URL = "https://my-portfolio-jhonqs-projects.vercel.app";

function hostnameOf(value: string): string | null {
  const withProtocol = /^https?:\/\//i.test(value) ? value : `https://${value}`;
  try {
    return new URL(withProtocol).hostname.toLowerCase();
  } catch {
    return null;
  }
}

function isLocalHostname(hostname: string): boolean {
  return hostname === "localhost" || hostname === "127.0.0.1" || hostname.endsWith(".local");
}

/**
 * Deployment hosts such as my-portfolio-fdfm6zckc-jhonqs-projects.vercel.app
 * change on every Vercel build. Google Search Console rejects them when the
 * sitemap belongs to the stable production alias.
 */
function isDeploymentHostname(hostname: string): boolean {
  const deploymentUrl = process.env.VERCEL_URL?.toLowerCase();
  if (deploymentUrl && hostname === deploymentUrl) return true;
  return /^my-portfolio-[a-z0-9]{7,}-jhonqs-projects\.vercel\.app$/.test(hostname);
}

export function getSiteUrl(): string {
  const productionHost = process.env.VERCEL_PROJECT_PRODUCTION_URL
    ? hostnameOf(process.env.VERCEL_PROJECT_PRODUCTION_URL)
    : null;

  if (productionHost && !isLocalHostname(productionHost) && !isDeploymentHostname(productionHost)) {
    return `https://${productionHost}`;
  }

  const configured = process.env.NEXT_PUBLIC_SITE_URL?.trim().replace(/\/+$/, "");
  if (configured) {
    const host = hostnameOf(configured);
    if (host && !isLocalHostname(host) && !isDeploymentHostname(host)) {
      return configured.split("#")[0];
    }
  }

  return CANONICAL_SITE_URL;
}
