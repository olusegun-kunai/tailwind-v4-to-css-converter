# 🔄 **New Requirements: UnoCSS-Based Tailwind to CSS Modules Converter**

## **Project Pivot: From @apply Strategy to UnoCSS Generator**

### **Context**
The lead has identified limitations with our current @apply + Tailwind CLI compilation approach and is pivoting to a more robust solution using UnoCSS's `createGenerator` function with Tailwind 4 preset.

---

## 🎯 **New Objective**

Create a **TypeScript-based converter** that transforms Qwik components with Tailwind classes into **CSS Modules** using UnoCSS's programmatic API.

### **Technology Stack Change**
- **FROM**: Go-based CLI with @apply strategy
- **TO**: TypeScript-based converter with UnoCSS generator

---

## 🔧 **Core Technical Approach**

### **1. UnoCSS Generator Setup**
```typescript
import { createGenerator } from '@unocss/core'
import presetWind from '@unocss/preset-wind'

const generator = createGenerator({
  presets: [presetWind()] // Tailwind 4 preset
})
```

### **2. Node-by-Node Processing**
- Parse JSX string template literal line by line
- Identify JSX nodes/elements
- Extract Tailwind classes from each node individually
- Generate CSS using UnoCSS generator for each node
- Assign semantic class names based on component type

---

## 📝 **Processing Flow**

### **Input Example**
```tsx
import { component$, useStyles$ } from "@builder.io/qwik";
import { Checkbox } from "@kunai-consulting/qwik";

export default component$(() => {
  return (
    <Checkbox.Root>
      <Checkbox.HiddenInput />
      <div class="flex items-center gap-2">
        <Checkbox.Trigger
          class="size-[25px] rounded-lg relative bg-gray-500 
                 focus-visible:outline focus-visible:outline-1 focus-visible:outline-white
                 disabled:opacity-50 bg-qwik-neutral-200 data-[checked]:bg-qwik-blue-800 focus-visible:ring-[3px] ring-qwik-blue-600"
        >
          <Checkbox.Indicator
            class="data-[checked]:flex justify-center items-center absolute inset-0"
          >
            <LuCheck />
          </Checkbox.Indicator>
        </Checkbox.Trigger>
        <Checkbox.Label class="text-sm">
          This is a trusted device, don't ask again
        </Checkbox.Label>
      </div>
    </Checkbox.Root>
  );
});
```

### **Processing Steps**

#### **Step 1: JSX Node Identification**
For each line, determine if it contains a JSX node:
- `<Checkbox.Indicator class="data-[checked]:flex justify-center items-center absolute inset-0">`
- Extract classes: `"data-[checked]:flex justify-center items-center absolute inset-0"`

#### **Step 2: UnoCSS Generation**
```typescript
const classes = "data-[checked]:flex justify-center items-center absolute inset-0";
const result = await generator.generate(classes);
// Returns CSS for each class
```

#### **Step 3: Class Name Assignment**
- **Component nodes**: Semantic names based on component type
  - `<Checkbox.Indicator />` → `styles.indicator`
- **HTML nodes**: Generic indexed names
  - `<div />` → `styles.node0`
  - `<span />` → `styles.node1`

---

## 🎨 **CSS Generation Rules**

### **Basic Class Conversion**
```tsx
// Input
<Checkbox.Indicator class="flex justify-center items-center" />

// Generated CSS
.indicator {
  display: flex;
  justify-content: center;
  align-items: center;
}
```

### **Modifier Handling**
Modifiers should create **separate CSS declarations**:

```tsx
// Input
<Checkbox.Indicator class="bg-red-100 focus:bg-red-500" />

// Generated CSS
.indicator {
  background-color: var(--color-bg-red-100);
}

.indicator:focus {
  background-color: var(--color-bg-red-500);
}
```

### **Generic HTML Elements**
```tsx
// Input
<div class="flex items-center gap-2">

// Generated CSS
.node0 {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}
```

---

## 📁 **Expected Output Structure**

### **CSS Modules File** (`checkbox.module.css`)
```css
.indicator {
  display: flex;
  justify-content: center;
  align-items: center;
}

.indicator:focus {
  background-color: var(--color-bg-red-500);
}

.trigger {
  width: 25px;
  height: 25px;
  border-radius: 0.5rem;
  position: relative;
  background-color: var(--color-gray-500);
}

.node0 {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}
```

### **Updated Component File**
```tsx
import { component$, useStyles$ } from "@builder.io/qwik";
import { Checkbox } from "@kunai-consulting/qwik";
import styles from './checkbox.module.css';

export default component$(() => {
  return (
    <Checkbox.Root>
      <Checkbox.HiddenInput />
      <div class={styles.node0}>
        <Checkbox.Trigger class={styles.trigger}>
          <Checkbox.Indicator class={styles.indicator}>
            <LuCheck />
          </Checkbox.Indicator>
        </Checkbox.Trigger>
        <Checkbox.Label class={styles.label}>
          This is a trusted device, don't ask again
        </Checkbox.Label>
      </div>
    </Checkbox.Root>
  );
});
```

---

## 🔧 **Implementation Requirements**

### **1. JSX Parser**
- Parse TypeScript/JSX files
- Extract class attributes from each JSX node
- Identify component vs HTML element types
- Handle multi-line class attributes

### **2. UnoCSS Integration**
- Set up UnoCSS generator with Tailwind 4 preset
- Process classes node by node
- Handle modifiers and variants properly
- Generate consolidated CSS per node

### **3. Semantic Naming Logic**
- **Component Detection**: `<Checkbox.Indicator />` → `indicator`
- **HTML Elements**: `<div />`, `<span />` → `node0`, `node1`, etc.
- **Future Enhancement**: AI-powered semantic naming for HTML elements

### **4. CSS Modules Output**
- Generate `.module.css` file with unique class names
- Update JSX to import and use CSS modules
- Handle CSS custom properties (CSS variables)

---

## 📚 **Key Resources**

### **UnoCSS Documentation**
- [UnoCSS Core Tools](https://unocss.dev/tools/core)
- [createGenerator Function](https://github.com/unocss/unocss/blob/1031312057a3bea1082b7d938eb2ad640f57613a/test/preset-wind4.test.ts#L12)

### **Implementation Priorities**
1. **Phase 1**: Set up UnoCSS generator with hardcoded string input
2. **Phase 2**: Build JSX parser to extract class attributes
3. **Phase 3**: Implement semantic naming logic
4. **Phase 4**: Generate CSS modules output
5. **Phase 5**: Handle complex scenarios (modifiers, variants, etc.)

---

## ✅ **Success Criteria**

- [ ] UnoCSS generator working with Tailwind 4 preset
- [ ] Node-by-node class extraction from JSX
- [ ] Semantic naming for component elements
- [ ] Generic naming for HTML elements
- [ ] CSS modules generation with proper imports
- [ ] Modifier handling (focus:, hover:, etc.)
- [ ] Updated component files with CSS modules imports

---

## 🚀 **Next Steps**

1. **Create TypeScript project structure**
2. **Install UnoCSS dependencies**
3. **Set up basic generator with hardcoded test**
4. **Build JSX parsing logic**
5. **Implement semantic naming rules**
6. **Generate CSS modules output**

---

## 📝 **Notes**

- This approach eliminates the brittle @apply compilation issues
- UnoCSS's programmatic API provides reliable CSS generation
- CSS modules ensure scoped styles with unique hashes
- Future AI enhancement can improve HTML element naming
- TypeScript provides better maintainability vs Go for this use case 