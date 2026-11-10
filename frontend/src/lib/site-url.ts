const CANONICAL_HOST = "my-portfolio-jhonqs-projects.vercel.app";
const CANONICAL_SITE_URL = `https://${CANONICAL_HOST}`;

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
 * Search Console is registered on the stable alias. Any other *.vercel.app
 * host (deployment URLs like my-portfolio-fdfm6zckc-jhonqs-projects.vercel.app
 * or my-portfolio-gamma-teal-3v1i358tc4.vercel.app) is rejected there.
 * A custom domain is still accepted.
 */
function isUnstableHost(hostname: string): boolean {
  if (isLocalHostname(hostname)) return true;
  return hostname.endsWith(".vercel.app") && hostname !== CANONICAL_HOST;
}

function stableUrl(value: string | undefined): string | null {
  if (!value) return null;
  const trimmed = value.trim().replace(/\/+$/, "").split("#")[0];
  const host = hostnameOf(trimmed);
  if (!host || isUnstableHost(host)) return null;
  return /^https?:\/\//i.test(trimmed) ? trimmed : `https://${host}`;
}

export function getSiteUrl(): string {
  return (
    stableUrl(process.env.NEXT_PUBLIC_SITE_URL) ||
    stableUrl(process.env.VERCEL_PROJECT_PRODUCTION_URL) ||
    CANONICAL_SITE_URL
  );
}
