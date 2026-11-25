"use client";

import { useEffect } from "react";

export default function VisitTracker() {
  useEffect(() => {
    const controller = new AbortController();
    fetch("/api/counter", { method: "POST", signal: controller.signal }).catch(() => {
      // A failed count must not affect the page.
    });
    return () => controller.abort();
  }, []);

  return null;
}
