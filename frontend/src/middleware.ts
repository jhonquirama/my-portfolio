import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

const RETIRED_CV_PATHS = new Set(["/cv-jhon-quirama-es.pdf", "/cv-jhon-quirama-en.pdf"]);

export function middleware(request: NextRequest) {
  if (!RETIRED_CV_PATHS.has(request.nextUrl.pathname)) return NextResponse.next();

  return new NextResponse("Not found", {
    status: 404,
    headers: {
      "X-Robots-Tag": "noindex, nofollow, noarchive",
      "Cache-Control": "no-store",
    },
  });
}

export const config = {
  matcher: ["/cv-jhon-quirama-es.pdf", "/cv-jhon-quirama-en.pdf"],
};
