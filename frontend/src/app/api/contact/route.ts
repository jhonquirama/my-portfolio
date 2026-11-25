import { createHash } from "crypto";
import { NextRequest, NextResponse } from "next/server";
import { z } from "zod";
import { checkRateLimit, clientAddress } from "@/lib/rate-limiter";
import { supabaseAdmin } from "@/lib/supabase";
import { Resend } from "resend";

const projectTypes = ["custom_software", "backend_api", "cloud_arch", "consulting", "full_time"] as const;
const budgets = ["under_2k", "2k_5k", "5k_15k", "15k_plus", "discuss"] as const;

const contactSchema = z.object({
  name: z.string().min(2, "El nombre debe tener al menos 2 caracteres").max(100),
  email: z.string().email("Correo electrónico inválido"),
  projectType: z.enum(projectTypes).default("custom_software"),
  budget: z.enum(budgets).default("discuss"),
  message: z.string().min(5, "El mensaje debe tener al menos 5 caracteres").max(3000),
  honeypot: z.string().optional(),
});

function escapeHtml(value: string): string {
  return value
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#39;");
}

function singleLine(value: string): string {
  return value.replace(/[\r\n\t]+/g, " ").trim();
}

function hashIp(ip: string): string {
  return createHash("sha256").update(`${ip}-portfolio-contact`).digest("hex").slice(0, 16);
}

export async function POST(req: NextRequest) {
  try {
    // 1. IP extraction & Rate limiting
    const ip = clientAddress(req.headers, req.ip);

    const rateLimit = checkRateLimit(`contact:${ip}`, 5, 60_000);
    if (!rateLimit.allowed) {
      return NextResponse.json(
        {
          error: "Has alcanzado el límite de intentos. Por favor espera un minuto antes de volver a intentar.",
        },
        { status: 429 }
      );
    }

    // 2. Body Parsing & Validation
    const body = await req.json();
    const parsed = contactSchema.safeParse(body);

    if (!parsed.success) {
      const errorMsg = parsed.error.issues.map((i) => i.message).join(", ");
      return NextResponse.json({ error: errorMsg }, { status: 400 });
    }

    const { name, email, projectType, budget, message, honeypot } = parsed.data;
    const safeName = escapeHtml(name);
    const safeEmail = escapeHtml(email);
    const safeMessage = escapeHtml(message);
    const ipHash = hashIp(ip);

    // 3. Honeypot check (Silent drop for bots)
    if (honeypot && honeypot.trim().length > 0) {
      return NextResponse.json({ success: true, message: "Lead recibido" });
    }

    // 4. Store in Supabase if configured
    let savedInDb = false;
    if (supabaseAdmin) {
      try {
        const { error: dbError } = await supabaseAdmin.from("contact_leads").insert([
          {
            name,
            email,
            project_type: projectType,
            budget,
            message,
            ip_hash: ipHash,
          },
        ]);
        if (dbError) console.error("contact_leads insert failed:", dbError.message);
        else savedInDb = true;
      } catch (dbErr) {
        console.error("Supabase insert error:", dbErr);
      }
    }

    // 5. Send notification email via Resend
    const resendApiKey = process.env.RESEND_API_KEY;
    const receiverEmail = process.env.CONTACT_RECEIVER_EMAIL || "quirama80@gmail.com";
    let emailSent = false;

    if (resendApiKey) {
      try {
        const resend = new Resend(resendApiKey);
        await resend.emails.send({
          from: "Jhon Quirama Portfolio <onboarding@resend.dev>",
          to: receiverEmail,
          replyTo: email,
          subject: `Nuevo lead de software: ${singleLine(name)} (${projectType})`,
          html: `
            <div style="font-family: sans-serif; background: #07090e; color: #f1f5f9; padding: 30px; border-radius: 12px;">
              <h2 style="color: #00f5d4; margin-top: 0;">Nuevo contacto desde tu portafolio</h2>
              <hr style="border: none; border-top: 1px solid #1e293b; margin: 20px 0;" />
              <p><strong>Nombre:</strong> ${safeName}</p>
              <p><strong>Email:</strong> <a href="mailto:${safeEmail}" style="color: #38bdf8;">${safeEmail}</a></p>
              <p><strong>Tipo de Proyecto:</strong> ${escapeHtml(projectType)}</p>
              <p><strong>Presupuesto Estimado:</strong> ${escapeHtml(budget)}</p>
              <div style="background: #0d121f; padding: 20px; border-radius: 8px; border: 1px solid #1e293b; margin-top: 15px;">
                <p style="margin: 0; color: #94a3b8; font-size: 13px;">MENSAJE:</p>
                <p style="margin-top: 8px; white-space: pre-wrap; font-size: 15px;">${safeMessage}</p>
              </div>
              <p style="font-size: 11px; color: #64748b; margin-top: 25px;">
                Enviado desde el portafolio de Jhon Quirama
              </p>
            </div>
          `,
        });
        emailSent = true;
      } catch (mailErr) {
        console.error("Resend error:", mailErr);
      }
    }

    if (supabaseAdmin && !savedInDb) {
      return NextResponse.json(
        { error: "No se pudo guardar el mensaje. Inténtalo de nuevo en un momento." },
        { status: 502 }
      );
    }

    if (!savedInDb && !emailSent) {
      return NextResponse.json(
        { error: "No se pudo enviar el mensaje. Inténtalo de nuevo en un momento." },
        { status: 502 }
      );
    }

    return NextResponse.json({
      success: true,
      message: "Mensaje procesado con éxito",
    });
  } catch (error) {
    console.error("Contact API route error:", error);
    return NextResponse.json(
      { error: "Ocurrió un error interno al procesar el mensaje." },
      { status: 500 }
    );
  }
}
