"use client";

import React, { useState } from "react";
import { useLanguage } from "./LanguageContext";
import { Terminal, Globe, Menu, X, ArrowUpRight } from "lucide-react";

export default function Navbar() {
  const { lang, setLang, t } = useLanguage();
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  const toggleLang = () => {
    setLang(lang === "es" ? "en" : "es");
  };

  const navLinks = [
    { href: "#services", label: t.nav.services },
    { href: "#experience", label: t.nav.experience },
    { href: "#architecture", label: t.nav.architecture },
    { href: "#terminal", label: t.nav.terminal },
    { href: "#contact", label: t.nav.contact },
  ];

  return (
    <header className="fixed top-0 left-0 right-0 z-50 px-4 sm:px-6 lg:px-8 pt-4">
      <div className="max-w-6xl mx-auto backdrop-blur-xl bg-[#0d121f]/80 border border-slate-800/80 rounded-2xl px-5 py-3 shadow-2xl shadow-cyan-950/20 flex items-center justify-between transition-all">
        {/* Brand / Logo */}
        <a href="#" className="flex items-center space-x-2.5 group">
          <div className="w-8 h-8 rounded-lg bg-gradient-to-tr from-cyan-500 to-blue-600 flex items-center justify-center text-black font-black text-sm shadow-md shadow-cyan-500/20 group-hover:scale-105 transition-transform">
            JQ
          </div>
          <div className="flex flex-col">
            <span className="font-bold tracking-tight text-white group-hover:text-cyan-400 transition-colors text-sm sm:text-base">
              Jhon Quirama
            </span>
            <span className="text-[10px] text-slate-400 font-mono flex items-center gap-1">
              <span className="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse" />
              {t.nav.status}
            </span>
          </div>
        </a>

        {/* Desktop Nav */}
        <nav className="hidden md:flex items-center space-x-1 lg:space-x-2 text-sm text-slate-300 font-medium">
          {navLinks.map((link) => (
            <a
              key={link.href}
              href={link.href}
              className="px-3 py-1.5 rounded-lg hover:text-cyan-400 hover:bg-slate-800/50 transition-colors"
            >
              {link.label}
            </a>
          ))}
        </nav>

        {/* Action Controls */}
        <div className="hidden sm:flex items-center space-x-3">
          {/* Language Switcher */}
          <button
            onClick={toggleLang}
            className="flex items-center space-x-1.5 px-3 py-1.5 rounded-lg border border-slate-700 hover:border-cyan-500/50 bg-slate-900/60 text-xs font-mono text-slate-300 hover:text-cyan-400 transition-all shadow-sm"
            title="Toggle Language / Cambiar Idioma"
          >
            <Globe className="w-3.5 h-3.5 text-cyan-400" />
            <span className="font-bold">{lang.toUpperCase()}</span>
          </button>

          {/* CTA Quote Button */}
          <a
            href="#contact"
            className="group relative inline-flex items-center gap-1.5 px-4 py-1.5 rounded-xl text-xs sm:text-sm font-semibold text-black bg-cyan-400 hover:bg-cyan-300 shadow-lg shadow-cyan-500/25 transition-all hover:scale-[1.02] active:scale-[0.98]"
          >
            <span>{t.nav.hireMe}</span>
            <ArrowUpRight className="w-3.5 h-3.5 group-hover:translate-x-0.5 group-hover:-translate-y-0.5 transition-transform" />
          </a>
        </div>

        {/* Mobile Menu Button */}
        <div className="flex sm:hidden items-center space-x-2">
          <button
            onClick={toggleLang}
            className="p-1.5 rounded-lg border border-slate-700 bg-slate-900/80 text-xs font-mono text-cyan-400"
          >
            {lang.toUpperCase()}
          </button>
          <button
            onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
            className="p-2 text-slate-300 hover:text-white"
            aria-label="Toggle menu"
          >
            {mobileMenuOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
          </button>
        </div>
      </div>

      {/* Mobile Menu Dropdown */}
      {mobileMenuOpen && (
        <div className="md:hidden mt-2 max-w-6xl mx-auto backdrop-blur-2xl bg-[#0d121f]/95 border border-slate-800 rounded-2xl p-4 shadow-2xl flex flex-col space-y-3">
          {navLinks.map((link) => (
            <a
              key={link.href}
              href={link.href}
              onClick={() => setMobileMenuOpen(false)}
              className="px-3 py-2 rounded-lg text-slate-200 hover:text-cyan-400 hover:bg-slate-800/60 font-medium text-sm"
            >
              {link.label}
            </a>
          ))}
          <a
            href="#contact"
            onClick={() => setMobileMenuOpen(false)}
            className="w-full text-center py-2.5 rounded-xl text-sm font-semibold text-black bg-cyan-400 hover:bg-cyan-300"
          >
            {t.nav.hireMe}
          </a>
        </div>
      )}
    </header>
  );
}
