-- ====================================================================
-- SUPABASE SCHEMA FOR JHON QUIRAMA PORTFOLIO
-- Execute this script in your Supabase Project -> SQL Editor
-- ====================================================================

-- 1. Table: Contact Leads (Solicitudes de software a la medida & mensajes)
CREATE TABLE IF NOT EXISTS public.contact_leads (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::text, now()) NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    project_type TEXT NOT NULL DEFAULT 'custom_software',
    budget TEXT,
    timeline TEXT,
    message TEXT NOT NULL,
    ip_hash TEXT,
    status TEXT NOT NULL DEFAULT 'new' CHECK (status IN ('new', 'contacted', 'archived', 'spam'))
);

-- Row Level Security (RLS) on contact_leads
ALTER TABLE public.contact_leads ENABLE ROW LEVEL SECURITY;

-- Allow anonymous users to ONLY insert new valid leads (never read other leads)
DROP POLICY IF EXISTS "Allow anonymous lead submission" ON public.contact_leads;
CREATE POLICY "Allow anonymous lead submission" 
    ON public.contact_leads 
    FOR INSERT 
    TO anon 
    WITH CHECK (
        length(name) > 1 AND 
        length(email) > 4 AND 
        length(message) > 5
    );

-- Allow authenticated users (Jhon logged in to Supabase) full access
DROP POLICY IF EXISTS "Allow owner read and write leads" ON public.contact_leads;
CREATE POLICY "Allow owner read and write leads" 
    ON public.contact_leads 
    FOR ALL 
    TO authenticated 
    USING (true);


-- 2. Table: Visitor Counter & Page Views (Métricas y Contador en Vivo)
CREATE TABLE IF NOT EXISTS public.site_analytics (
    id SERIAL PRIMARY KEY,
    page_path TEXT NOT NULL UNIQUE,
    views_count BIGINT DEFAULT 1 NOT NULL,
    unique_visitors_count BIGINT DEFAULT 1 NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT timezone('utc'::text, now()) NOT NULL
);

-- Table for deduplicating daily unique visitors anonymously (stores salted hash)
CREATE TABLE IF NOT EXISTS public.visitor_daily_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    visitor_hash TEXT NOT NULL,
    visit_date DATE DEFAULT CURRENT_DATE NOT NULL,
    page_path TEXT NOT NULL,
    CONSTRAINT unique_daily_visit UNIQUE (visitor_hash, visit_date, page_path)
);

-- Enable RLS for site_analytics
ALTER TABLE public.site_analytics ENABLE ROW LEVEL SECURITY;
ALTER TABLE public.visitor_daily_logs ENABLE ROW LEVEL SECURITY;

-- Anyone can read view counts
DROP POLICY IF EXISTS "Public can view analytics summary" ON public.site_analytics;
CREATE POLICY "Public can view analytics summary" 
    ON public.site_analytics 
    FOR SELECT 
    TO anon, authenticated 
    USING (true);

-- Atomic RPC function to increment views safely
CREATE OR REPLACE FUNCTION public.record_page_visit(page_target TEXT, visitor_hash_input TEXT)
RETURNS JSONB
LANGUAGE plpgsql
SECURITY DEFINER
AS $$
DECLARE
    is_unique BOOLEAN := FALSE;
    total_views BIGINT;
    unique_visitors BIGINT;
BEGIN
    -- Check if visitor is unique for today
    IF visitor_hash_input IS NOT NULL AND visitor_hash_input <> '' THEN
        BEGIN
            INSERT INTO public.visitor_daily_logs (visitor_hash, visit_date, page_path)
            VALUES (visitor_hash_input, CURRENT_DATE, page_target);
            is_unique := TRUE;
        EXCEPTION WHEN unique_violation THEN
            is_unique := FALSE;
        END;
    END IF;

    -- Upsert counter
    INSERT INTO public.site_analytics (page_path, views_count, unique_visitors_count, updated_at)
    VALUES (page_target, 1, CASE WHEN is_unique THEN 1 ELSE 0 END, now())
    ON CONFLICT (page_path) DO UPDATE SET
        views_count = public.site_analytics.views_count + 1,
        unique_visitors_count = public.site_analytics.unique_visitors_count + (CASE WHEN is_unique THEN 1 ELSE 0 END),
        updated_at = now()
    RETURNING views_count, unique_visitors_count INTO total_views, unique_visitors;

    RETURN jsonb_build_object(
        'page_path', page_target,
        'views', total_views,
        'unique_visitors', unique_visitors,
        'is_new_today', is_unique
    );
END;
$$;
