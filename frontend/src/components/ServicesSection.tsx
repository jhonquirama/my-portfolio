"use client";

import React from "react";
import { useLanguage } from "./LanguageContext";
import { Code2, Server, Database, CloudCog, Check, ArrowRight } from "lucide-react";

export default function ServicesSection() {
  const { t } = useLanguage();

  const iconMap: Record<string, React.ReactNode> = {
    "custom-dev": <Code2 className="w-6 h-6 text-cyan-400" />,
    "backend-distributed": <Server className="w-6 h-6 text-sky-400" />,
    "db-optimization": <Database className="w-6 h-6 text-indigo-400" />,
    "cloud-devops": <CloudCog className="w-6 h-6 text-emerald-400" />,
  };

  return (
    <section id="services" className="py-24 relative z-10">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Section Header */}
        <div className="mb-16">
          <div className="inline-block text-xs font-mono font-semibold tracking-wider text-cyan-400 uppercase mb-2">
            {t.services.tag}
          </div>
          <h2 className="text-3xl sm:text-4xl lg:text-5xl font-extrabold text-white tracking-tight">
            {t.services.title}
          </h2>
          <p className="text-base sm:text-lg text-slate-400 mt-3 max-w-2xl">
            {t.services.subtitle}
          </p>
        </div>

        {/* Services Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          {t.services.items.map((service) => (
            <div
              key={service.id}
              className="group relative p-8 rounded-3xl bg-gradient-to-b from-slate-900/90 to-slate-950/80 border border-slate-800 hover:border-cyan-500/50 backdrop-blur-xl transition-all duration-300 hover:shadow-2xl hover:shadow-cyan-950/30 flex flex-col justify-between"
            >
              {/* Glow accent */}
              <div className="absolute top-0 right-0 w-32 h-32 bg-cyan-500/5 rounded-full blur-3xl group-hover:bg-cyan-500/15 transition-all" />

              <div>
                <div className="flex items-center justify-between mb-6">
                  <div className="w-12 h-12 rounded-2xl bg-slate-800/80 border border-slate-700/80 flex items-center justify-center shadow-inner">
                    {iconMap[service.id]}
                  </div>
                  <span className="px-3 py-1 rounded-full text-[11px] font-mono font-semibold text-cyan-300 bg-cyan-950/60 border border-cyan-800/50">
                    {service.badge}
                  </span>
                </div>

                <h3 className="text-xl sm:text-2xl font-bold text-white mb-3 group-hover:text-cyan-300 transition-colors">
                  {service.title}
                </h3>

                <p className="text-sm sm:text-base text-slate-300 leading-relaxed mb-6 font-normal">
                  {service.description}
                </p>

                <div className="space-y-2.5 mb-8">
                  {service.features.map((feature, idx) => (
                    <div key={idx} className="flex items-start gap-2.5 text-xs sm:text-sm text-slate-400">
                      <Check className="w-4 h-4 text-emerald-400 shrink-0 mt-0.5" />
                      <span>{feature}</span>
                    </div>
                  ))}
                </div>
              </div>

              <a
                href="#contact"
                className="inline-flex items-center gap-2 text-xs sm:text-sm font-semibold text-cyan-400 hover:text-cyan-300 group/link pt-4 border-t border-slate-800/80 transition-colors"
              >
                <span>Solicitar propuesta técnica</span>
                <ArrowRight className="w-4 h-4 group-hover/link:translate-x-1 transition-transform" />
              </a>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
