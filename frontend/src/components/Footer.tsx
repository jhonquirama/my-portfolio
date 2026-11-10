"use client";

import React from "react";
import { useLanguage } from "./LanguageContext";
import { Github, Linkedin, Shield, Heart } from "lucide-react";

export default function Footer() {
  const { t } = useLanguage();

  return (
    <footer className="py-12 border-t border-slate-800/80 bg-[#06080d] relative z-10 text-xs text-slate-400 font-mono">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex flex-col md:flex-row items-center justify-between gap-6">
          {/* Brand */}
          <div className="flex items-center space-x-3">
            <div className="w-7 h-7 rounded-lg bg-gradient-to-tr from-cyan-500 to-blue-600 flex items-center justify-center text-black font-black text-xs">
              JQ
            </div>
            <div>
              <p className="font-bold text-slate-200">Jhon Quirama</p>
              <p className="text-[11px] text-slate-500">
                Senior Backend & Distributed Systems Engineer
              </p>
            </div>
          </div>

          {/* Social Links */}
          <div className="flex items-center space-x-4">
            <a
              href="https://www.linkedin.com/in/jhon-quirama"
              target="_blank"
              rel="noopener noreferrer"
              className="p-2 rounded-lg bg-slate-900 border border-slate-800 hover:border-cyan-500/50 hover:text-cyan-400 transition-colors"
              aria-label="LinkedIn"
            >
              <Linkedin className="w-4 h-4" />
            </a>
            <a
              href="https://github.com/jhonquirama"
              target="_blank"
              rel="noopener noreferrer"
              className="p-2 rounded-lg bg-slate-900 border border-slate-800 hover:border-cyan-500/50 hover:text-cyan-400 transition-colors"
              aria-label="GitHub"
            >
              <Github className="w-4 h-4" />
            </a>
          </div>

          {/* Security & Tech notes */}
          <div className="flex flex-col items-center md:items-end text-center md:text-right text-[11px] text-slate-500 space-y-1">
            <p className="flex items-center gap-1.5 text-cyan-400/80">
              <Shield className="w-3.5 h-3.5 text-cyan-400" />
              <span>{t.footer.securityNote}</span>
            </p>
            <p>© {new Date().getFullYear()} Jhon Quirama. {t.footer.rights}</p>
          </div>
        </div>
      </div>
    </footer>
  );
}
