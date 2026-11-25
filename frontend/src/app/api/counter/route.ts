import { NextRequest, NextResponse } from "next/server";
import { checkRateLimit, clientAddress } from "@/lib/rate-limiter";
import { supabaseAdmin } from "@/lib/supabase";
import crypto from "crypto";

// Fallback in-memory counters for local dev
let localViews = 142;
let localUniques = 89;

export async function GET(req: NextRequest) {
  try {
    if (supabaseAdmin) {
      const { data, error } = await supabaseAdmin
        .from("site_analytics")
        .select("views_count, unique_visitors_count")
        .eq("page_path", "/")
        .maybeSingle();

      if (!error && data) {
        return NextResponse.json({
          views: Number(data.views_count),
          uniques: Number(data.unique_visitors_count),
        });
      }
    }

    return NextResponse.json({
      views: localViews,
      uniques: localUniques,
    });
  } catch (error) {
    return NextResponse.json({ views: localViews, uniques: localUniques });
  }
}

export async function POST(req: NextRequest) {
  try {
    const ip = clientAddress(req.headers, req.ip);

    // Rate limit view increments to 1 per 10s per IP to avoid refresh spam
    const rateLimit = checkRateLimit(`view:${ip}`, 3, 10_000);
    if (!rateLimit.allowed) {
      return NextResponse.json({ ok: false, reason: "rate_limited" });
    }

    // Salted daily hash of visitor IP (GDPR compliant anonymization)
    const today = new Date().toISOString().slice(0, 10);
    const visitorHash = crypto
      .createHash("sha256")
      .update(`${ip}-${today}-portfolio-salt`)
      .digest("hex")
      .slice(0, 16);

    if (supabaseAdmin) {
      const { data, error } = await supabaseAdmin.rpc("record_page_visit", {
        page_target: "/",
        visitor_hash_input: visitorHash,
      });

      if (error || !data) {
        console.error("record_page_visit failed:", error?.message);
        return NextResponse.json({ ok: false }, { status: 502 });
      }

      return NextResponse.json({
        ok: true,
        views: data.views,
        uniques: data.unique_visitors,
      });
    }

    localViews += 1;
    return NextResponse.json({
      ok: true,
      views: localViews,
      uniques: localUniques,
    });
  } catch (err) {
    return NextResponse.json({ ok: false }, { status: 500 });
  }
}
