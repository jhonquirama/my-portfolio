import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

const PRODUCTION_HOST = "my-portfolio-gamma-teal-3v1i358tc4.vercel.app";

// Vercel adds X-Robots-Tag: noindex on these aliases. Google also rejects
// their sitemap when it points at the production host. Send them there.
const NOINDEX_ALIAS_HOSTS = new Set([
  "my-portfolio-jhonqs-projects.vercel.app",
  "my-portfolio-git-master-jhonqs-projects.vercel.app",
]);

const RETIRED_CV_PATHS = new Set(["/cv-jhon-quirama-es.pdf", "/cv-jhon-quirama-en.pdf"]);

export function middleware(request: NextRequest) {
  const host = request.headers.get("host")?.split(":")[0]?.toLowerCase();

  if (host && NOINDEX_ALIAS_HOSTS.has(host)) {
    const destination = request.nextUrl.clone();
    destination.protocol = "https";
    destination.hostname = PRODUCTION_HOST;
    destination.port = "";
    return NextResponse.redirect(destination, 308);
  }

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
  matcher: ["/((?!_next/static|_next/image|favicon.ico|.*\\.(?:svg|png|jpg|jpeg|gif|webp|ico)$).*)"],
};
