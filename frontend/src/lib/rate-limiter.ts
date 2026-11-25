interface RateLimitRecord {
  count: number;
  resetAt: number;
}

type HeaderSource = {
  get(name: string): string | null;
};

// Prefer the address set by the platform. The left-most X-Forwarded-For value is client-controlled.
export function clientAddress(headers: HeaderSource, platformIp?: string | null): string {
  const realIp = headers.get("x-real-ip")?.trim();
  if (realIp) return realIp;
  if (platformIp?.trim()) return platformIp.trim();

  const forwarded = headers.get("x-forwarded-for");
  if (forwarded) {
    const hops = forwarded.split(",").map((part) => part.trim()).filter(Boolean);
    const nearest = hops.at(-1);
    if (nearest) return nearest;
  }

  return "127.0.0.1";
}

const memoryStore = new Map<string, RateLimitRecord>();

/**
 * Sliding window rate limiter for API protection
 * @param key Identifier (e.g. client IP + route)
 * @param maxRequests Maximum allowed requests in the time window
 * @param windowMs Window duration in milliseconds (default 60 seconds)
 */
export function checkRateLimit(
  key: string,
  maxRequests = 5,
  windowMs = 60_000
): { allowed: boolean; remaining: number; resetInMs: number } {
  const now = Date.now();
  const record = memoryStore.get(key);

  // Periodically clean expired records if map gets large
  if (memoryStore.size > 5000) {
    memoryStore.forEach((v, k) => {
      if (v.resetAt < now) {
        memoryStore.delete(k);
      }
    });
  }

  if (!record || record.resetAt < now) {
    memoryStore.set(key, {
      count: 1,
      resetAt: now + windowMs,
    });
    return {
      allowed: true,
      remaining: maxRequests - 1,
      resetInMs: windowMs,
    };
  }

  if (record.count >= maxRequests) {
    return {
      allowed: false,
      remaining: 0,
      resetInMs: Math.max(0, record.resetAt - now),
    };
  }

  record.count += 1;
  return {
    allowed: true,
    remaining: maxRequests - record.count,
    resetInMs: Math.max(0, record.resetAt - now),
  };
}
