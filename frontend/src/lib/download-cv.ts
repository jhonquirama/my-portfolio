export type CvDownloadCode = "invalid_token" | "rate_limited" | "unavailable" | "network";

export class CvDownloadError extends Error {
  code: CvDownloadCode;

  constructor(code: CvDownloadCode) {
    super(code);
    this.code = code;
  }
}

export async function downloadCvFile(token: string, lang: "es" | "en") {
  let response: Response;
  try {
    response = await fetch("/api/cv", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ token: token.trim(), lang }),
    });
  } catch {
    throw new CvDownloadError("network");
  }

  const contentType = response.headers.get("content-type") || "";
  if (!response.ok || !contentType.includes("application/pdf")) {
    let code: CvDownloadCode = "unavailable";
    try {
      const body = (await response.json()) as { error?: string };
      if (body.error === "invalid_token" || body.error === "rate_limited" || body.error === "unavailable") {
        code = body.error;
      }
    } catch {
      code = response.status === 401 ? "invalid_token" : "unavailable";
    }
    throw new CvDownloadError(code);
  }

  const blob = await response.blob();
  const filename = lang === "en" ? "cv-jhon-quirama-en.pdf" : "cv-jhon-quirama-es.pdf";
  const url = URL.createObjectURL(blob);
  const anchor = document.createElement("a");
  anchor.href = url;
  anchor.download = filename;
  document.body.appendChild(anchor);
  anchor.click();
  anchor.remove();
  URL.revokeObjectURL(url);
}
