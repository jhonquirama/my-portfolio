"use client";

import React, { createContext, useContext, useEffect, useState } from "react";
import { translations, Language } from "@/content/translations";

interface LanguageContextType {
  lang: Language;
  setLang: (lang: Language) => void;
  t: typeof translations.es;
}

const LanguageContext = createContext<LanguageContextType | undefined>(undefined);

export function LanguageProvider({ children }: { children: React.ReactNode }) {
  const [lang, setLangState] = useState<Language>("es");

  useEffect(() => {
    // Check saved preference
    const saved = localStorage.getItem("portfolio_lang") as Language | null;
    if (saved && (saved === "es" || saved === "en")) {
      setLangState(saved);
      document.documentElement.lang = saved;
    } else {
      // Auto-detect browser language
      const browserLang = navigator.language?.toLowerCase() || "";
      const initialLang: Language = browserLang.startsWith("es") ? "es" : "en";
      setLangState(initialLang);
      document.documentElement.lang = initialLang;
    }
  }, []);

  const setLang = (newLang: Language) => {
    setLangState(newLang);
    localStorage.setItem("portfolio_lang", newLang);
    document.documentElement.lang = newLang;
  };

  const t = translations[lang];

  return (
    <LanguageContext.Provider value={{ lang, setLang, t }}>
      {children}
    </LanguageContext.Provider>
  );
}

export function useLanguage() {
  const context = useContext(LanguageContext);
  if (!context) {
    throw new Error("useLanguage must be used within a LanguageProvider");
  }
  return context;
}
