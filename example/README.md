# Example Project: Tailwind to CSS Converter Demo

This example demonstrates how the Tailwind v4+ to CSS Converter works on a realistic React project.

## What's Included

**Original Components (with Tailwind classes):**
- `src/Card.jsx` - Product card with hover effects, shadows, and responsive design
- `src/Header.jsx` - Navigation header with responsive menu and cart icon
- `src/ProductGrid.jsx` - Grid layout with responsive breakpoints

## Before Conversion

The components use typical Tailwind classes like:
```jsx
<div class="bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300">
<button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
```

## Running the Conversion

From the root directory:

```bash
# Convert the example project
./tailwind-converter --input ./example/src --output ./example/dist --verbose
```

## After Conversion

The converted components will:
- Use semantic CSS class names (`styles.card_container_1`, `styles.button_primary_2`)
- Have automatically generated CSS modules
- Include proper CSS module imports
- Maintain the exact same visual appearance

## Key Demonstrations

1. **Complex Layout**: Product grid with responsive breakpoints
2. **Interactive Elements**: Hover states, focus rings, transitions
3. **Typography**: Various font sizes, weights, and colors
4. **Spacing**: Margins, padding, gaps in different contexts
5. **Visual Effects**: Shadows, borders, rounded corners

This example shows how real-world Tailwind components convert to maintainable CSS modules. 



## 🎯 **Today's Implementation Plan**

Based on our standup update, we need to **pivot our existing tool** to match your lead's requirements. Here's what we'll implement:

### **1. Modify Component Parser** (30 mins)
**Change from:** Generic class names like `div_layout_1`, `button_1`  
**Change to:** Semantic names like `otp-root`, `otp-input`, `accordion-trigger`

```go
// Instead of: div_layout_1
// Generate: otp-root (based on component context)
```

### **2. Implement @apply Strategy** (1 hour)
**Replace our direct CSS mapping with:**
- Generate intermediate CSS with `@apply` syntax
- Use Tailwind CLI to compile it to real CSS

```go
// Generate this intermediate step:
// .otp-root { @apply flex justify-center gap-2; }
// 
// Then compile with Tailwind to get:
// .otp-root { display: flex; justify-content: center; gap: 0.5rem; }
```

### **3. Create Dual Output System** (45 mins)
**Modify our current single output to generate:**
- `component-tailwind.jsx` (original Tailwind version)
- `component-vanilla.jsx` (semantic class version)  
- `styles.css` (compiled vanilla CSS)

### **4. Test with Qwik Components** (30 mins)
**Update our example to use:**
- Qwik component syntax (`<Otp.Root>` instead of `<div>`)
- Test the @apply → compilation workflow

## 🛠 **Practical Steps**
1. **Update the converter logic** to use @apply strategy
2. **Add Tailwind CLI integration** for CSS compilation
3. **Modify file generation** for dual output
4. **Test with a simple Qwik OTP component**

**Goal:** By end of day, input 1 Tailwind component → get 2 components + 1 CSS file (like Base UI examples)




That's **totally normal and actually smart!** 🧠 Let's make sure you **really understand** this before coding. 

## 🎯 **Practice Explaining: The Elevator Pitch**

**Imagine someone asks you:** *"What are you building?"*

Try explaining it in **one sentence first:**

---

### **Your Turn:** Complete this sentence:
*"I'm building a tool that takes _____ and converts it into _____ so that _____."*

<details>
<summary>Click to see the answer</summary>

*"I'm building a tool that takes **Qwik components with Tailwind classes** and converts it into **both Tailwind and vanilla CSS versions** so that **documentation websites can show developers both options."***

</details>

---

## 🤔 **Test Your Understanding**

**Scenario:** Your friend asks these questions. How would you answer?

### **Q1:** "Why do you need BOTH versions?"
**Your answer:** _(try answering before looking below)_

<details>
<summary>Answer</summary>
Because some developers love Tailwind CSS, but others prefer regular CSS. Instead of forcing everyone to use one approach, the documentation shows both, so everyone's happy.
</details>

### **Q2:** "What's the actual problem you're solving?"
**Your answer:** _(try answering)_

<details>
<summary>Answer</summary>
Right now, if I want to document a component, I have to manually write it twice - once with Tailwind classes and once with regular CSS. That's double work and prone to mistakes. My tool automates this.
</details>

### **Q3:** "Give me a concrete example"
**Your answer:** _(try explaining with an example)_

<details>
<summary>Answer</summary>

**Input:** One file with `<button class="bg-blue-500 px-4 py-2 rounded">`

**Output:** 
- **Tailwind version:** `<button class="bg-blue-500 px-4 py-2 rounded">`
- **CSS version:** `<button class="my-button">` + CSS file with `.my-button { background: blue; padding: 1rem; border-radius: 0.25rem; }`

Same button, two different ways to write it.

</details>

## ✅ **Ready Check**

**Can you now explain this to:**
- [ ] Your mom (in simple terms)
- [ ] A developer friend (with technical details)  
- [ ] Your lead (the business value)

**Once you feel confident explaining it, we'll start coding!** 🚀

**How did the practice go? Which part still feels unclear?**

## 🎉 **AMAZING! It's Working!**

Let's see what we've accomplished:

### ✅ **What We Built:**

1. **Dual Output System** - Creates both Tailwind and vanilla CSS versions
2. **Semantic Class Names** - Generated meaningful names like `button-primary`, `input-field`, `content-container`
3. **@apply Strategy** - CSS uses `@apply` for accurate Tailwind compilation
4. **Perfect for Documentation** - Exactly what qwik.design needs!

### 📁 **Generated Structure:**
```
test-output/
├── tailwind/qwik-otp.tsx    # Original with Tailwind classes
├── vanilla/qwik-otp.tsx     # Semantic classes version  
└── vanilla/qwik-otp.css     # @apply-based CSS
```

### 🎯 **Key Features Working:**

1. **Tailwind Version**: Unchanged original component
2. **Vanilla Version**: Semantic classes like `styles.button-primary`
3. **CSS File**: Uses `@apply` for accurate compilation
4. **Semantic Names**: Meaningful instead of generic

Let's **test with a more complex example** to show the power:

