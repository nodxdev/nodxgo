import { toPascalCase } from "jsr:@std/text@1.0.10";

export interface El {
  name: string;
  isVoid: boolean;
  description: string;
}

export interface Attr {
  name: string;
  description: string;
}

export const staticConflicts = ["map"];

export function hasConflict(
  name: string,
  els: El[],
  attrs: Attr[],
): boolean {
  for (const conflict of staticConflicts) {
    if (conflict.toLowerCase() === name.toLowerCase()) return true;
  }

  let matches = 0;

  for (const el of els) {
    if (el.name.toLowerCase() === name.toLowerCase()) matches++;
  }
  for (const attr of attrs) {
    if (attr.name.toLowerCase() === name.toLowerCase()) matches++;
  }

  return matches > 1;
}

export interface FuncName {
  name: string;
  isGlob: boolean;
}

export function createFuncName(
  name: string,
  type: "El" | "Attr",
  hasConflict: boolean,
): FuncName {
  const isGlob = name.endsWith("*");
  if (isGlob) {
    name = name.slice(0, -1);
  }

  const word = name.split("-").join(" ");
  name = toPascalCase(word.toLowerCase());

  if (hasConflict) {
    name = `${name}${type}`;
  }

  return { name, isGlob };
}
