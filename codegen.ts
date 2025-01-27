#!/usr/bin/env -S deno run -A

const elements = await getElements();
const attributes = await getAttributes();
const keywords = await getKeywords();

generateElements(elements, keywords);
generateAttributes(attributes, keywords);
runFormatter();

///////////////
// Functions //
///////////////

interface El {
  name: string;
  isVoid: boolean;
  description: string;
}

interface Attr {
  name: string;
  description: string;
}

async function getJson<T>(path: string): Promise<T> {
  const response = await fetch(path);
  return response.json();
}

async function getElements(): Promise<El[]> {
  return await getJson(
    "https://raw.githubusercontent.com/nodxdev/nodx/refs/heads/main/data/elements.json",
  );
}

async function getAttributes(): Promise<Attr[]> {
  return await getJson(
    "https://raw.githubusercontent.com/nodxdev/nodx/refs/heads/main/data/attributes.json",
  );
}

async function getKeywords(): Promise<string[]> {
  const langs = await getJson<{ [lang: string]: string[] }>(
    "https://raw.githubusercontent.com/nodxdev/nodx/refs/heads/main/data/keywords.json",
  );

  const keywords: string[] = [];
  for (const lang in langs) {
    keywords.push(...langs[lang]);
  }

  return [...new Set(keywords)];
}

function runFormatter() {
  const cmd = new Deno.Command("task", { args: ["fmt"] });
  const out = cmd.outputSync();
  if (!out.success) {
    console.error(new TextDecoder().decode(out.stderr));
    Deno.exit(1);
  }
}

function generateElements(elements: El[], keywords: string[]) {
  const fileContent: string[] = [`package nodx`, ``];

  for (const element of elements) {
    const funcName = resolveNameConflict(element.name, "El", keywords);

    if (element.isVoid) {
      fileContent.push(`// ${funcName} ${element.description}`);
      fileContent.push(`func ${funcName}(children ...Node) Node {`);
      fileContent.push(`\treturn ElVoid("${element.name}", children...)`);
      fileContent.push(`}`);
      fileContent.push("");
    }

    if (!element.isVoid) {
      fileContent.push(`// ${funcName} ${element.description}`);
      fileContent.push(`func ${funcName}(children ...Node) Node {`);
      fileContent.push(`\treturn El("${element.name}", children...)`);
      fileContent.push(`}`);
      fileContent.push("");
    }
  }

  Deno.removeSync("gen_elements.go");
  Deno.writeTextFileSync("gen_elements.go", fileContent.join("\n"));
  console.log("Generated gen_elements.go");
}

function generateAttributes(attrs: Attr[], keywords: string[]) {
  const fileContent: string[] = [`package nodx`, ``];

  for (const attr of attrs) {
    const funcName = resolveNameConflict(attr.name, "Attr", keywords);
    fileContent.push(`// ${funcName} ${attr.description}`);
    fileContent.push(`func ${funcName}(value string) Node {`);
    fileContent.push(`\treturn Attr("${attr.name}", value)`);
    fileContent.push(`}`);
    fileContent.push("");
  }

  Deno.removeSync("gen_attributes.go");
  Deno.writeTextFileSync("gen_attributes.go", fileContent.join("\n"));
  console.log("Generated gen_attributes.go");
}

function resolveNameConflict(
  name: string,
  suffix: string,
  keywords: string[],
): string {
  const camelCaseName = capitalize(name);
  if (keywords.includes(name.toLowerCase())) {
    return `${camelCaseName}${suffix}`;
  }
  return camelCaseName;
}

function capitalize(str: string): string {
  return str.charAt(0).toUpperCase() + str.slice(1);
}
