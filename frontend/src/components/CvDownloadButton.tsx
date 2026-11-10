"use client";

import React, { useEffect, useId, useRef, useState } from "react";
import { Download, X } from "lucide-react";
import { useLanguage } from "./LanguageContext";
import { CvDownloadError, downloadCvFile } from "@/lib/download-cv";

export default function CvDownloadButton() {
  const { t, lang } = useLanguage();
  const titleId = useId();
  const inputRef = useRef<HTMLInputElement>(null);
  const [open, setOpen] = useState(false);
  const [code, setCode] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (!open) return;
    const frame = window.requestAnimationFrame(() => inputRef.current?.focus());
    const onKey = (event: KeyboardEvent) => {
      if (event.key === "Escape" && !loading) setOpen(false);
    };
    window.addEventListener("keydown", onKey);
    return () => {
      window.cancelAnimationFrame(frame);
      window.removeEventListener("keydown", onKey);
    };
  }, [open, loading]);

  const close = () => {
    if (loading) return;
    setOpen(false);
    setError("");
  };

  const messageFor = (codeName: CvDownloadError["code"]) => {
    if (codeName === "invalid_token") return t.hero.cvInvalid;
    if (codeName === "rate_limited") return t.hero.cvRate;
    return t.hero.cvUnavailable;
  };

  const onSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (!code.trim() || loading) return;
    setError("");
    setLoading(true);
    try {
      await downloadCvFile(code, lang);
      setCode("");
      setOpen(false);
    } catch (err) {
      const codeName = err instanceof CvDownloadError ? err.code : "unavailable";
      setError(messageFor(codeName));
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <button
        type="button"
        onClick={() => {
          setError("");
          setOpen(true);
        }}
        className="inline-flex items-center gap-2 px-5 py-3.5 rounded-xl text-sm font-mono text-slate-200 hover:text-cyan-300 border border-slate-700 hover:border-cyan-400/40 bg-slate-950/40 transition-all"
      >
        <Download className="w-4 h-4" />
        <span>{t.hero.downloadCv}</span>
      </button>

      {open && (
        <div className="fixed inset-0 z-[80] flex items-end sm:items-center justify-center p-4">
          <button
            type="button"
            aria-label={t.hero.cvClose}
            className="absolute inset-0 bg-[#07090e]/80"
            onClick={close}
          />
          <form
            role="dialog"
            aria-modal="true"
            aria-labelledby={titleId}
            onSubmit={onSubmit}
            className="relative w-full max-w-md rounded-2xl border border-slate-700 bg-[#0d121c] p-6 shadow-2xl shadow-black/40"
          >
            <div className="flex items-start justify-between gap-4 mb-3">
              <h2 id={titleId} className="text-xl font-bold text-white tracking-tight">
                {t.hero.cvTitle}
              </h2>
              <button
                type="button"
                onClick={close}
                className="rounded-lg p-1.5 text-slate-300 hover:text-white hover:bg-slate-800"
                aria-label={t.hero.cvClose}
              >
                <X className="w-4 h-4" />
              </button>
            </div>
            <p className="text-sm leading-relaxed text-slate-200 mb-5">{t.hero.cvBody}</p>
            <label className="block text-xs font-semibold uppercase tracking-wide text-cyan-200 mb-2" htmlFor={`${titleId}-code`}>
              {t.hero.cvPlaceholder}
            </label>
            <input
              id={`${titleId}-code`}
              ref={inputRef}
              value={code}
              onChange={(event) => setCode(event.target.value)}
              autoComplete="off"
              autoCapitalize="off"
              autoCorrect="off"
              spellCheck={false}
              disabled={loading}
              placeholder={t.hero.cvPlaceholder}
              className="w-full rounded-xl border border-slate-600 bg-[#07090e] px-4 py-3 text-sm font-mono text-white placeholder:text-slate-400 outline-none focus:border-cyan-300 disabled:opacity-60"
            />
            {error && (
              <p role="alert" className="mt-3 text-sm text-rose-300">
                {error}
              </p>
            )}
            <div className="mt-5 flex items-center justify-end gap-3">
              <button
                type="button"
                onClick={close}
                disabled={loading}
                className="px-4 py-2.5 rounded-xl text-sm font-semibold text-slate-200 hover:text-white disabled:opacity-60"
              >
                {t.hero.cvClose}
              </button>
              <button
                type="submit"
                disabled={loading || code.trim().length < 10}
                className="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl text-sm font-bold text-black bg-cyan-400 hover:bg-cyan-300 disabled:opacity-50 disabled:hover:bg-cyan-400"
              >
                <Download className="w-4 h-4" />
                <span>{loading ? t.hero.cvWorking : t.hero.cvAction}</span>
              </button>
            </div>
          </form>
        </div>
      )}
    </>
  );
}
