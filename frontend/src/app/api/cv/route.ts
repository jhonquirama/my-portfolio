import { NextRequest } from "next/server";
import { z } from "zod";
import { checkRateLimit } from "@/lib/rate-limiter";
import { cvErrorResponse, serveCvDownload } from "@/lib/cv-access";

const bodySchema = z.object({
  token: z.string().trim().min(10).max(200),
  lang: z.enum(["es", "en"]),
});

function clientIp(req: NextRequest) {
  return (
    req.headers.get("x-forwarded-for")?.split(",")[0]?.trim() ||
    req.headers.get("x-real-ip") ||
    "127.0.0.1"
  );
}

export async function POST(req: NextRequest) {
  const rateLimit = checkRateLimit(`cv:${clientIp(req)}`, 20, 10 * 60_000);
  if (!rateLimit.allowed) return cvErrorResponse("rate_limited", 429);

  let json: unknown;
  try {
    json = await req.json();
  } catch {
    return cvErrorResponse("invalid_token", 401);
  }

  const parsed = bodySchema.safeParse(json);
  if (!parsed.success) return cvErrorResponse("invalid_token", 401);

  return serveCvDownload(parsed.data.token, parsed.data.lang);
}
