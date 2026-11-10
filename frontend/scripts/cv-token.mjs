import { createHash, randomBytes, randomUUID } from "crypto";
import { existsSync, readFileSync } from "fs";
import path from "path";
import { fileURLToPath } from "url";
import { createClient } from "@supabase/supabase-js";

const BUCKET = "cv";
const TOKEN_OBJECT = "tokens.json";
const CANONICAL_SITE_URL = "https://my-portfolio-jhonqs-projects.vercel.app";
const PDFS = {
  es: "cv-jhon-quirama-es.pdf",
  en: "cv-jhon-quirama-en.pdf",
};

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");

function loadEnv() {
  for (const name of [".env.local", ".env"]) {
    const file = path.join(root, name);
    if (!existsSync(file)) continue;
    for (const line of readFileSync(file, "utf8").split("\n")) {
      const trimmed = line.trim();
      if (!trimmed || trimmed.startsWith("#") || !trimmed.includes("=")) continue;
      const index = trimmed.indexOf("=");
      const key = trimmed.slice(0, index).trim();
      let value = trimmed.slice(index + 1).trim();
      if (
        (value.startsWith('"') && value.endsWith('"')) ||
        (value.startsWith("'") && value.endsWith("'"))
      ) {
        value = value.slice(1, -1);
      }
      if (!process.env[key]) process.env[key] = value;
    }
  }
}

function siteUrl() {
  const configured = (process.env.NEXT_PUBLIC_SITE_URL || "").trim().replace(/\/+$/, "");
  const deploymentHost = /my-portfolio-[a-z0-9]{7,}-jhonqs-projects\.vercel\.app/;
  if (
    configured.startsWith("https://") &&
    !configured.includes("localhost") &&
    !deploymentHost.test(configured)
  ) {
    return configured.split("#")[0];
  }
  return CANONICAL_SITE_URL;
}

function client() {
  const url = process.env.NEXT_PUBLIC_SUPABASE_URL;
  const key = process.env.SUPABASE_SERVICE_ROLE_KEY;
  if (!url || !key) {
    console.error("Faltan NEXT_PUBLIC_SUPABASE_URL o SUPABASE_SERVICE_ROLE_KEY en frontend/.env.local");
    process.exit(1);
  }
  return createClient(url, key, { auth: { autoRefreshToken: false, persistSession: false } });
}

function hashToken(token) {
  return createHash("sha256").update(token).digest("hex");
}

async function readTokens(supabase) {
  const { data, error } = await supabase.storage.from(BUCKET).download(TOKEN_OBJECT);
  if (error || !data) return { version: 1, tokens: [] };
  const parsed = JSON.parse(await data.text());
  if (!parsed || !Array.isArray(parsed.tokens)) return { version: 1, tokens: [] };
  return parsed;
}

async function writeTokens(supabase, file) {
  const { error } = await supabase.storage.from(BUCKET).upload(TOKEN_OBJECT, Buffer.from(JSON.stringify(file)), {
    upsert: true,
    contentType: "application/json",
    cacheControl: "0",
  });
  if (error) {
    console.error("No se pudo guardar la lista de códigos.");
    process.exit(1);
  }
}

function flag(args, name) {
  const index = args.indexOf(`--${name}`);
  if (index === -1 || !args[index + 1] || args[index + 1].startsWith("--")) return undefined;
  return args[index + 1];
}

function statusOf(token) {
  if (token.revokedAt) return "revocado";
  if (token.expiresAt && Date.parse(token.expiresAt) <= Date.now()) return "vencido";
  if (token.maxDownloads !== null && token.downloadCount >= token.maxDownloads) return "agotado";
  return "activo";
}

async function setup(supabase) {
  const { error: createError } = await supabase.storage.createBucket(BUCKET, { public: false });
  if (createError && !/already exists|duplicate/i.test(createError.message)) {
    console.error("No se pudo crear el bucket privado del CV.");
    process.exit(1);
  }

  const { error: privacyError } = await supabase.storage.updateBucket(BUCKET, { public: false });
  if (privacyError) {
    console.error("No se pudo dejar el bucket del CV como privado.");
    process.exit(1);
  }

  const { data: bucket, error: bucketError } = await supabase.storage.getBucket(BUCKET);
  if (bucketError || !bucket || bucket.public) {
    console.error("El bucket del CV sigue público. No subí los archivos.");
    process.exit(1);
  }

  for (const filename of Object.values(PDFS)) {
    const filePath = path.join(root, "private", "cv", filename);
    if (!existsSync(filePath)) {
      console.error(`Falta el archivo local ${path.relative(root, filePath)}`);
      process.exit(1);
    }
    const { error } = await supabase.storage.from(BUCKET).upload(filename, readFileSync(filePath), {
      upsert: true,
      contentType: "application/pdf",
      cacheControl: "0",
    });
    if (error) {
      console.error(`No se pudo subir ${filename}.`);
      process.exit(1);
    }
  }

  const existing = await supabase.storage.from(BUCKET).download(TOKEN_OBJECT);
  if (existing.error || !existing.data) {
    await writeTokens(supabase, { version: 1, tokens: [] });
  }

  console.log("Listo. Los PDF quedaron en un bucket privado de Supabase y los códigos siguen guardados.");
}

async function create(supabase, args) {
  const label = flag(args, "label");
  if (!label) {
    console.error('Indica a quién se lo das: npm run cv:token -- create --label "Nombre o empresa"');
    process.exit(1);
  }

  const days = flag(args, "days");
  const max = flag(args, "max");
  const expiresAt = days ? new Date(Date.now() + Number(days) * 24 * 60 * 60 * 1000).toISOString() : null;
  const maxDownloads = max ? Number(max) : null;
  if (days && !Number.isFinite(Number(days))) {
    console.error("--days tiene que ser un número.");
    process.exit(1);
  }
  if (max && (!Number.isInteger(maxDownloads) || maxDownloads < 1)) {
    console.error("--max tiene que ser un entero mayor que 0.");
    process.exit(1);
  }

  const token = `cv_${randomBytes(32).toString("base64url")}`;
  const file = await readTokens(supabase);
  file.tokens.push({
    id: randomUUID(),
    label,
    hash: hashToken(token),
    expiresAt,
    maxDownloads,
    downloadCount: 0,
    revokedAt: null,
    createdAt: new Date().toISOString(),
    lastUsedAt: null,
  });
  await writeTokens(supabase, file);

  const base = siteUrl();
  console.log("");
  console.log(`Código creado para: ${label}`);
  console.log(`Vence: ${expiresAt ? expiresAt.slice(0, 10) : "nunca"}`);
  console.log(`Descargas máximas: ${maxDownloads ?? "sin límite"}`);
  console.log("");
  console.log("Compártelo una sola vez. Quien tenga el código o el enlace puede descargar el CV.");
  console.log(`Código: ${token}`);
  console.log(`Español: ${base}/cv/${token}`);
  console.log(`Inglés:  ${base}/cv/${token}?lang=en`);
  console.log("");
}

async function list(supabase) {
  const file = await readTokens(supabase);
  if (file.tokens.length === 0) {
    console.log("No hay códigos creados.");
    return;
  }
  for (const token of file.tokens) {
    const uses = token.maxDownloads === null ? String(token.downloadCount) : `${token.downloadCount}/${token.maxDownloads}`;
    const expires = token.expiresAt ? token.expiresAt.slice(0, 10) : "nunca";
    console.log(`${token.id}  ${statusOf(token).padEnd(9)}  ${uses.padEnd(8)}  vence ${expires}  ${token.label}`);
  }
}

async function revoke(supabase, args) {
  const id = flag(args, "id");
  if (!id) {
    console.error("Indica el id que muestra `npm run cv:token -- list`.");
    process.exit(1);
  }
  const file = await readTokens(supabase);
  const matches = file.tokens.filter((token) => token.id === id || token.id.startsWith(id));
  if (matches.length !== 1) {
    console.error(matches.length === 0 ? "No encontré ese código." : "Ese id es ambiguo. Pega el id completo.");
    process.exit(1);
  }
  matches[0].revokedAt = new Date().toISOString();
  await writeTokens(supabase, file);
  console.log(`Código revocado: ${matches[0].label}`);
}

loadEnv();
const [command, ...args] = process.argv.slice(2);
const supabase = client();

if (command === "setup") {
  await setup(supabase);
} else if (command === "create") {
  await create(supabase, args);
} else if (command === "list") {
  await list(supabase);
} else if (command === "revoke") {
  await revoke(supabase, args);
} else {
  console.log(`Uso:
  npm run cv:token -- setup
  npm run cv:token -- create --label "Empresa" --days 30 --max 3
  npm run cv:token -- list
  npm run cv:token -- revoke --id <id>`);
}
