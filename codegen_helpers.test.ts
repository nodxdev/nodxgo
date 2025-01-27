import { assertEquals } from "jsr:@std/assert@1.0.11";
import {
  createFuncName,
  hasConflict,
  staticConflicts,
} from "./codegen_helpers.ts";
import type { Attr, El } from "./codegen_helpers.ts";

Deno.test("test hasConflict", async (t) => {
  await t.step("no conflict", () => {
    const els: El[] = [
      { name: "a", isVoid: false, description: "" },
      { name: "b", isVoid: false, description: "" },
    ];
    const attrs: Attr[] = [
      { name: "c", description: "" },
      { name: "d", description: "" },
    ];

    assertEquals(hasConflict("z", els, attrs), false);
  });

  await t.step("one conflict", () => {
    const els: El[] = [
      { name: "a", isVoid: false, description: "" },
      { name: "b", isVoid: false, description: "" },
    ];
    const attrs: Attr[] = [
      { name: "c", description: "" },
      { name: "d", description: "" },
    ];

    assertEquals(hasConflict("a", els, attrs), false);
  });

  await t.step("two conflicts", () => {
    const els: El[] = [
      { name: "a", isVoid: false, description: "" },
      { name: "b", isVoid: false, description: "" },
    ];
    const attrs: Attr[] = [
      { name: "a", description: "" },
      { name: "c", description: "" },
    ];

    assertEquals(hasConflict("a", els, attrs), true);
  });

  await t.step("three conflicts", () => {
    const els: El[] = [
      { name: "a", isVoid: false, description: "" },
      { name: "b", isVoid: false, description: "" },
    ];
    const attrs: Attr[] = [
      { name: "a", description: "" },
      { name: "b", description: "" },
    ];

    assertEquals(hasConflict("a", els, attrs), true);
  });

  await t.step("static conflict", () => {
    for (const conflict of staticConflicts) {
      assertEquals(hasConflict(conflict, [], []), true);
    }
  });
});

Deno.test("test createFuncName", async (t) => {
  await t.step("normal element name", () => {
    assertEquals(createFuncName("div", "El", false), {
      name: "Div",
      isGlob: false,
    });

    assertEquals(createFuncName("DiV", "El", false), {
      name: "Div",
      isGlob: false,
    });

    assertEquals(createFuncName("dIv", "El", false), {
      name: "Div",
      isGlob: false,
    });
  });

  await t.step("normal attribute name", () => {
    assertEquals(createFuncName("class", "Attr", false), {
      name: "Class",
      isGlob: false,
    });

    assertEquals(createFuncName("ClAsS", "Attr", false), {
      name: "Class",
      isGlob: false,
    });

    assertEquals(createFuncName("cLaSs", "Attr", false), {
      name: "Class",
      isGlob: false,
    });
  });

  await t.step("element with conflict", () => {
    assertEquals(createFuncName("div", "El", true), {
      name: "DivEl",
      isGlob: false,
    });

    assertEquals(createFuncName("DiV", "El", true), {
      name: "DivEl",
      isGlob: false,
    });

    assertEquals(createFuncName("dIv", "El", true), {
      name: "DivEl",
      isGlob: false,
    });
  });

  await t.step("attribute with conflict", () => {
    assertEquals(createFuncName("class", "Attr", true), {
      name: "ClassAttr",
      isGlob: false,
    });

    assertEquals(createFuncName("ClAsS", "Attr", true), {
      name: "ClassAttr",
      isGlob: false,
    });

    assertEquals(createFuncName("cLaSs", "Attr", true), {
      name: "ClassAttr",
      isGlob: false,
    });
  });

  await t.step("glob attribute", () => {
    assertEquals(createFuncName("data-*", "Attr", false), {
      name: "Data",
      isGlob: true,
    });

    assertEquals(createFuncName("data-*", "Attr", true), {
      name: "DataAttr",
      isGlob: true,
    });
  });
});
