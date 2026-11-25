import { NextRequest } from "next/server";
import { checkRateLimit, clientAddress } from "@/lib/rate-limiter";
import { cvErrorResponse, serveCvDownload, type CvLang } from "@/lib/cv-access";

export const dynamic = "force-dynamic";

function clientIp(req: NextRequest) {
  return clientAddress(req.headers, req.ip);
}

export async function GET(req: NextRequest, { params }: { params: { token: string } }) {
  const rateLimit = checkRateLimit(`cv:${clientIp(req)}`, 20, 10 * 60_000);
  if (!rateLimit.allowed) return cvErrorResponse("rate_limited", 429);

  const requested = req.nextUrl.searchParams.get("lang");
  const lang: CvLang = requested === "en" ? "en" : "es";
  let token = params.token || "";
  try {
    token = decodeURIComponent(token);
  } catch {
    return cvErrorResponse("invalid_token", 401);
  }

  return serveCvDownload(token, lang);
}
