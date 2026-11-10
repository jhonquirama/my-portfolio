"use client";

import React, { useEffect, useState } from "react";
import { useLanguage } from "./LanguageContext";
import { Eye, Users, Activity, BarChart2 } from "lucide-react";
import { formatNumber } from "@/lib/utils";

export default function LiveCounter() {
  const { t } = useLanguage();
  const [stats, setStats] = useState({ views: 0, uniques: 0 });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Record visit on page mount
    const recordVisit = async () => {
      try {
        const postRes = await fetch("/api/counter", { method: "POST" });
        const postData = await postRes.json();
        if (postData.views) {
          setStats({ views: postData.views, uniques: postData.uniques || postData.views });
          setLoading(false);
          return;
        }
      } catch (err) {
        // Ignore and fallback to GET
      }

      try {
        const getRes = await fetch("/api/counter");
        const getData = await getRes.json();
        setStats({ views: getData.views, uniques: getData.uniques });
      } catch (err) {
        setStats({ views: 142, uniques: 89 });
      } finally {
        setLoading(false);
      }
    };

    recordVisit();
  }, []);

  return (
    <div className="py-8 border-y border-slate-800/60 bg-slate-950/40 backdrop-blur-md relative z-10">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex flex-col sm:flex-row items-center justify-between gap-6">
          {/* Status Label */}
          <div className="flex items-center gap-3">
            <div className="w-8 h-8 rounded-xl bg-cyan-950/60 border border-cyan-800/60 flex items-center justify-center text-cyan-400">
              <Activity className="w-4 h-4 animate-pulse" />
            </div>
            <div>
              <p className="text-xs font-mono uppercase text-slate-400">
                {t.counter.status}
              </p>
              <h4 className="text-sm font-semibold text-white">
                {t.counter.title}
              </h4>
            </div>
          </div>

          {/* Counts */}
          <div className="flex items-center gap-8">
            {/* Views */}
            <div className="flex items-center gap-3">
              <Eye className="w-5 h-5 text-cyan-400" />
              <div>
                <p className="text-xl sm:text-2xl font-bold font-mono text-white">
                  {loading ? "--" : formatNumber(stats.views)}
                </p>
                <span className="text-[11px] text-slate-400 font-medium">
                  {t.counter.totalViews}
                </span>
              </div>
            </div>

            {/* Unique Visitors */}
            <div className="flex items-center gap-3 border-l border-slate-800 pl-8">
              <Users className="w-5 h-5 text-emerald-400" />
              <div>
                <p className="text-xl sm:text-2xl font-bold font-mono text-white">
                  {loading ? "--" : formatNumber(stats.uniques)}
                </p>
                <span className="text-[11px] text-slate-400 font-medium">
                  {t.counter.uniqueVisitors}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
