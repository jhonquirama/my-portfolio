"use client";

import React from "react";
import { useLanguage } from "./LanguageContext";
import { ArrowRight, Download, Terminal, CheckCircle2, ShieldCheck, Zap } from "lucide-react";

export default function Hero() {
  const { t, lang } = useLanguage();

  const techBadges = [
    "Golang",
    "Apache Kafka",
    "AWS & Kubernetes",
    "Hexagonal Arch",
    "PostgreSQL / PostGIS",
    "Redis Geospatial",
    "Next.js / TypeScript",
    "AI Workflows",
  ];

  return (
    <section className="relative pt-32 pb-20 sm:pt-40 sm:pb-28 overflow-hidden">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 relative z-10">
        {/* Top Tag */}
        <div className="inline-flex items-center gap-2 px-3 py-1.5 rounded-full border border-cyan-500/30 bg-cyan-950/30 backdrop-blur-md mb-6">
          <span className="w-2 h-2 rounded-full bg-cyan-400 animate-ping" />
          <span className="text-xs font-mono font-semibold tracking-wider text-cyan-300 uppercase">
            {t.hero.badge}
          </span>
        </div>

        {/* Main Headline */}
        <div className="max-w-4xl">
          <h1 className="text-4xl sm:text-6xl lg:text-7xl font-extrabold text-white tracking-tight leading-[1.1] mb-6">
            {t.hero.titleStart}{" "}
            <span className="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 via-sky-300 to-indigo-400">
              {t.hero.titleHighlight}
            </span>
          </h1>

          <p className="text-base sm:text-xl text-slate-300 leading-relaxed max-w-3xl mb-10 font-normal">
            {t.hero.description}
          </p>

          {/* Action CTAs */}
          <div className="flex flex-wrap items-center gap-4 mb-16">
            <a
              href="#contact"
              className="inline-flex items-center gap-2 px-6 py-3.5 rounded-xl text-sm sm:text-base font-bold text-black bg-cyan-400 hover:bg-cyan-300 shadow-xl shadow-cyan-500/25 transition-all hover:scale-105 active:scale-95"
            >
              <span>{t.hero.ctaPrimary}</span>
              <ArrowRight className="w-4 h-4" />
            </a>

            <a
              href="#experience"
              className="inline-flex items-center gap-2 px-6 py-3.5 rounded-xl text-sm sm:text-base font-semibold text-slate-200 bg-slate-900/80 hover:bg-slate-800 border border-slate-700 hover:border-slate-600 transition-all"
            >
              <span>{t.hero.ctaSecondary}</span>
            </a>

            <a
              href={lang === "es" ? "/cv-jhon-quirama-es.pdf" : "/cv-jhon-quirama-en.pdf"}
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center gap-2 px-5 py-3.5 rounded-xl text-sm font-mono text-slate-400 hover:text-cyan-400 border border-transparent hover:border-cyan-500/30 transition-all"
            >
              <Download className="w-4 h-4" />
              <span>{t.hero.downloadCv}</span>
            </a>
          </div>
        </div>

        {/* Floating Metrics Grid */}
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4 sm:gap-6 p-6 rounded-2xl bg-slate-900/60 backdrop-blur-xl border border-slate-800/80 shadow-2xl">
          {t.hero.metrics.map((metric, idx) => (
            <div key={idx} className="flex flex-col border-l-2 border-cyan-500/40 pl-4 py-1">
              <span className="text-3xl sm:text-4xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-white to-slate-200 font-mono tracking-tight">
                {metric.value}
              </span>
              <span className="text-xs sm:text-sm text-slate-400 mt-1 font-medium">
                {metric.label}
              </span>
            </div>
          ))}
        </div>

        {/* Tech Stack Ribbon */}
        <div className="mt-8 flex flex-wrap items-center gap-2 sm:gap-3">
          <span className="text-xs font-mono text-slate-500 uppercase tracking-wider mr-2">
            Core Stack:
          </span>
          {techBadges.map((badge) => (
            <span
              key={badge}
              className="px-3 py-1 rounded-lg text-xs font-mono font-medium text-slate-300 bg-slate-800/50 border border-slate-700/50 hover:border-cyan-500/40 hover:text-cyan-300 transition-colors"
            >
              {badge}
            </span>
          ))}
        </div>
      </div>
    </section>
  );
}
