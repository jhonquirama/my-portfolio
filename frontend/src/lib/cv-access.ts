import { createHash } from "crypto";
import { NextResponse } from "next/server";
import { supabaseService } from "@/lib/supabase";

export const CV_BUCKET = "cv";
const TOKEN_OBJECT = "tokens.json";

const PDF_OBJECTS = {
  es: "cv-jhon-quirama-es.pdf",
  en: "cv-jhon-quirama-en.pdf",
} as const;

export type CvLang = keyof typeof PDF_OBJECTS;
export type CvErrorCode = "invalid_token" | "rate_limited" | "unavailable";

type TokenRecord = {
  id: string;
  label: string;
  hash: string;
  expiresAt: string | null;
  maxDownloads: number | null;
  downloadCount: number;
  revokedAt: string | null;
  createdAt: string;
  lastUsedAt: string | null;
};

type TokenFile = {
  version: 1;
  tokens: TokenRecord[];
};

let bucketPrivacyChecked = false;

export function hashCvToken(token: string): string {
  return createHash("sha256").update(token).digest("hex");
}

export function cvErrorResponse(code: CvErrorCode, status: number) {
  return NextResponse.json(
    { error: code },
    {
      status,
      headers: {
        "Cache-Control": "no-store",
        "X-Robots-Tag": "noindex, nofollow, noarchive",
      },
    }
  );
}

async function ensurePrivateBucket() {
  if (!supabaseService || bucketPrivacyChecked) return;
  const { data } = await supabaseService.storage.getBucket(CV_BUCKET);
  if (data?.public) {
    await supabaseService.storage.updateBucket(CV_BUCKET, { public: false });
  }
  bucketPrivacyChecked = true;
}

async function readTokenFile(): Promise<TokenFile | null> {
  if (!supabaseService) return null;
  await ensurePrivateBucket();

  const { data, error } = await supabaseService.storage.from(CV_BUCKET).download(TOKEN_OBJECT);
  if (error || !data) return null;

  try {
    const parsed = JSON.parse(await data.text()) as TokenFile;
    if (!parsed || !Array.isArray(parsed.tokens)) return null;
    return parsed;
  } catch {
    return null;
  }
}

async function writeTokenFile(file: TokenFile) {
  if (!supabaseService) return false;
  const { error } = await supabaseService.storage.from(CV_BUCKET).upload(
    TOKEN_OBJECT,
    Buffer.from(JSON.stringify(file)),
    {
      upsert: true,
      contentType: "application/json",
      cacheControl: "0",
    }
  );
  return !error;
}

function tokenStatus(record: TokenRecord, now: number): "ok" | "invalid" {
  if (record.revokedAt) return "invalid";
  if (record.expiresAt && Date.parse(record.expiresAt) <= now) return "invalid";
  if (record.maxDownloads !== null && record.downloadCount >= record.maxDownloads) return "invalid";
  return "ok";
}

export async function serveCvDownload(token: string, lang: CvLang) {
  if (!supabaseService) return cvErrorResponse("unavailable", 503);

  const normalized = token.trim();
  if (!/^cv_[A-Za-z0-9_-]{40,80}$/.test(normalized)) {
    return cvErrorResponse("invalid_token", 401);
  }

  const file = await readTokenFile();
  if (!file) return cvErrorResponse("unavailable", 503);

  const hash = hashCvToken(normalized);
  const record = file.tokens.find((item) => item.hash === hash);
  if (!record || tokenStatus(record, Date.now()) !== "ok") {
    return cvErrorResponse("invalid_token", 401);
  }

  const pdfName = PDF_OBJECTS[lang];
  const { data, error } = await supabaseService.storage.from(CV_BUCKET).download(pdfName);
  if (error || !data) return cvErrorResponse("unavailable", 503);

  record.downloadCount += 1;
  record.lastUsedAt = new Date().toISOString();
  const saved = await writeTokenFile(file);
  if (!saved) return cvErrorResponse("unavailable", 503);

  const bytes = Buffer.from(await data.arrayBuffer());
  return new NextResponse(bytes, {
    status: 200,
    headers: {
      "Content-Type": "application/pdf",
      "Content-Disposition": `attachment; filename="${pdfName}"`,
      "Content-Length": String(bytes.byteLength),
      "Cache-Control": "private, no-store",
      "Referrer-Policy": "no-referrer",
      "X-Robots-Tag": "noindex, nofollow, noarchive",
    },
  });
}
