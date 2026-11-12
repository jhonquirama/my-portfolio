// Vercel adds `X-Robots-Tag: noindex` on my-portfolio-jhonqs-projects.vercel.app.
// The production host it leaves indexable is this one. A custom domain still wins.
const INDEXABLE_VERCEL_HOST = "my-portfolio-gamma-teal-3v1i358tc4.vercel.app";
const FALLBACK_SITE_URL = `https://${INDEXABLE_VERCEL_HOST}`;

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

function normalizedUrl(value: string | undefined): { url: string; host: string } | null {
  if (!value) return null;
  const trimmed = value.trim().replace(/\/+$/, "").split("#")[0];
  const host = hostnameOf(trimmed);
  if (!host || isLocalHostname(host)) return null;
  const url = /^https?:\/\//i.test(trimmed) ? trimmed : `https://${host}`;
  return { url, host };
}

export function getSiteUrl(): string {
  const configured = normalizedUrl(process.env.NEXT_PUBLIC_SITE_URL);
  if (configured && !configured.host.endsWith(".vercel.app")) return configured.url;

  const production = normalizedUrl(process.env.VERCEL_PROJECT_PRODUCTION_URL);
  if (production) return production.url;

  if (configured?.host === INDEXABLE_VERCEL_HOST) return configured.url;

  return FALLBACK_SITE_URL;
}
