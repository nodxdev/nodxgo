#!/usr/bin/env -S deno run -A
import {
  createFuncName,
  decapitalize,
  hasConflict,
} from "./codegen_helpers.ts";
import type { Attr, El } from "./codegen_helpers.ts";

const elements = await getElements();
const attributes = await getAttributes();

generateElements(elements, attributes);
generateAttributes(elements, attributes);
runFormatter();

///////////////
// Functions //
///////////////

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

function runFormatter() {
  const cmd = new Deno.Command("task", { args: ["fmt"] });
  const out = cmd.outputSync();
  if (!out.success) {
    console.error(new TextDecoder().decode(out.stderr));
    Deno.exit(1);
  }
}

function generateElements(els: El[], attrs: Attr[]) {
  const fileContent: string[] = [
    `package nodx`,
    ``,
    `// Code generated by NodX. DO NOT EDIT.`,
    ``,
  ];

  for (const el of els) {
    const isConflict = hasConflict(el.name, els, attrs);
    const funcName = createFuncName(el.name, "El", isConflict);
    el.description = decapitalize(el.description);

    if (el.isVoid) {
      fileContent.push(`// ${funcName.name} ${el.description}`);
      fileContent.push(`//`);
      fileContent.push(`// Output: <${el.name} ...>`);
      fileContent.push(`func ${funcName.name}(children ...Node) Node {`);
      fileContent.push(`\treturn ElVoid("${el.name}", children...)`);
      fileContent.push(`}`);
      fileContent.push("");
    }

    if (!el.isVoid) {
      fileContent.push(`// ${funcName.name} ${el.description}`);
      fileContent.push(`//`);
      fileContent.push(`// Output: <${el.name} ...>...</${el.name}>`);
      fileContent.push(`func ${funcName.name}(children ...Node) Node {`);
      fileContent.push(`\treturn El("${el.name}", children...)`);
      fileContent.push(`}`);
      fileContent.push("");
    }
  }

  const elNames = els.map((el) => el.name).join("\n");
  writeFile("generated_elements.go", fileContent.join("\n"));
  writeFile("generated_elements.txt", elNames);
  console.log("Generated generated_elements.go");
}

function generateAttributes(els: El[], attrs: Attr[]) {
  const fileContent: string[] = [
    `package nodx`,
    ``,
    `// Code generated by NodX. DO NOT EDIT.`,
    ``,
  ];

  for (const attr of attrs) {
    const isConflict = hasConflict(attr.name, els, attrs);
    const funcName = createFuncName(attr.name, "Attr", isConflict);
    attr.description = decapitalize(attr.description);

    if (funcName.isGlob) {
      attr.name = attr.name.replaceAll("*", "");
      fileContent.push(`// ${funcName.name} ${attr.description.toLowerCase()}`);
      fileContent.push(`//`);
      fileContent.push(`// Output: ${attr.name}{key}="{value}"`);
      fileContent.push(`func ${funcName.name}(key string,value string) Node {`);
      fileContent.push(`\treturn Attr("${attr.name}"+key, value)`);
      fileContent.push(`}`);
      fileContent.push("");
    }

    if (!funcName.isGlob) {
      fileContent.push(`// ${funcName.name} ${attr.description}`);
      fileContent.push(`//`);
      fileContent.push(`// Output: ${attr.name}="{value}"`);
      fileContent.push(`func ${funcName.name}(value string) Node {`);
      fileContent.push(`\treturn Attr("${attr.name}", value)`);
      fileContent.push(`}`);
      fileContent.push("");
    }
  }

  const attrNames = attrs.map((attr) => attr.name).join("\n");
  writeFile("generated_attributes.txt", attrNames);
  writeFile("generated_attributes.go", fileContent.join("\n"));
  console.log("Generated generated_attributes.go");
}

function writeFile(path: string, content: string) {
  try {
    Deno.removeSync(path);
  } catch {
    // no-op
  }
  Deno.writeTextFileSync(path, content);
}
