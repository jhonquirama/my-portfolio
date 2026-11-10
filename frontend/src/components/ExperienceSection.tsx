"use client";

import React from "react";
import { useLanguage } from "./LanguageContext";
import { Briefcase, Calendar, MapPin, CheckCircle, Award } from "lucide-react";

export default function ExperienceSection() {
  const { t } = useLanguage();

  return (
    <section id="experience" className="py-24 relative z-10 border-t border-slate-800/60">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Section Header */}
        <div className="mb-16">
          <div className="inline-block text-xs font-mono font-semibold tracking-wider text-cyan-400 uppercase mb-2">
            {t.experience.tag}
          </div>
          <h2 className="text-3xl sm:text-4xl lg:text-5xl font-extrabold text-white tracking-tight">
            {t.experience.title}
          </h2>
          <p className="text-base sm:text-lg text-slate-400 mt-3 max-w-2xl">
            {t.experience.subtitle}
          </p>
        </div>

        {/* Roles Timeline */}
        <div className="space-y-12">
          {t.experience.roles.map((role, idx) => (
            <div
              key={idx}
              className="relative pl-6 sm:pl-10 border-l-2 border-slate-800 hover:border-cyan-500/60 transition-colors group"
            >
              {/* Dot */}
              <div className="absolute -left-[9px] top-1.5 w-4 h-4 rounded-full bg-slate-900 border-2 border-slate-700 group-hover:border-cyan-400 group-hover:bg-cyan-950 transition-all" />

              <div className="p-6 sm:p-8 rounded-3xl bg-slate-900/50 border border-slate-800/80 backdrop-blur-md group-hover:border-slate-700 transition-all">
                {/* Meta details */}
                <div className="flex flex-wrap items-center justify-between gap-4 mb-4">
                  <div>
                    <h3 className="text-xl sm:text-2xl font-bold text-white flex items-center gap-2">
                      <span>{role.role}</span>
                      <span className="text-cyan-400">@ {role.company}</span>
                    </h3>
                  </div>

                  <div className="flex flex-wrap items-center gap-3 text-xs font-mono text-slate-400">
                    <span className="flex items-center gap-1.5 px-3 py-1 rounded-full bg-slate-800/80 border border-slate-700/60">
                      <Calendar className="w-3.5 h-3.5 text-cyan-400" />
                      {role.period}
                    </span>
                    <span className="flex items-center gap-1.5 px-3 py-1 rounded-full bg-slate-800/80 border border-slate-700/60">
                      <MapPin className="w-3.5 h-3.5 text-slate-400" />
                      {role.location}
                    </span>
                  </div>
                </div>

                <p className="text-sm sm:text-base text-slate-300 mb-6 leading-relaxed">
                  {role.description}
                </p>

                {/* Highlights */}
                <div className="space-y-3 mb-6">
                  {role.highlights.map((highlight, hIdx) => (
                    <div key={hIdx} className="flex items-start gap-3 text-xs sm:text-sm text-slate-300">
                      <CheckCircle className="w-4 h-4 text-cyan-400 shrink-0 mt-0.5" />
                      <span className="leading-relaxed">{highlight}</span>
                    </div>
                  ))}
                </div>

                {/* Tech Chips */}
                <div className="flex flex-wrap items-center gap-2 pt-4 border-t border-slate-800/80">
                  {role.tech.map((techItem) => (
                    <span
                      key={techItem}
                      className="px-2.5 py-1 rounded-lg text-xs font-mono font-medium text-cyan-300 bg-cyan-950/40 border border-cyan-900/60"
                    >
                      {techItem}
                    </span>
                  ))}
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
