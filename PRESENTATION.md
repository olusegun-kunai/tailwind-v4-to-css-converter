# 🎯 **Tailwind v4+ to CSS Converter**
*Transforming Tailwind Components into Dual Documentation*

---

## 📋 **Summary**

**Problem**: Manual creation of dual documentation (Tailwind + Vanilla CSS) is time-consuming and error-prone

**Solution**: Automated tool that converts Tailwind components into perfect dual documentation like base-ui.com

**Result**: One component input → Complete dual documentation output

---

## 🎯 **The Challenge this project Solved**

### **Before: Manual Dual Documentation**
```tsx
// 😓 Developer has to write BOTH versions manually

// Tailwind Version (Manual)
<div class="flex justify-center items-center min-h-screen bg-gray-50">
  <button class="bg-blue-600 hover:bg-blue-700 px-4 py-2 text-white rounded">
    Click me
  </button>
</div>

// Vanilla Version (Manual)
<div class="hero-container">
  <button class="primary-button">Click me</button>
</div>

/* Manual CSS */
.hero-container { display: flex; justify-content: center; /* ... */ }
.primary-button { background-color: #2563eb; /* ... */ }
```

### **After: One Component → Dual Documentation**
```bash
# ✨ Write once, get both versions automatically
./tailwind-converter -i qwik-otp.tsx -o output/
```

---


### **🔥 Core Features**
- ✅ **Semantic Class Generation**: `otp-root`, `otp-button`, `otp-input`
- ✅ **@apply Strategy**: Bridge between Tailwind utilities and vanilla CSS
- ✅ **Pure Vanilla CSS**: No Tailwind dependencies 
- ✅ **Theme Variables**: CSS custom properties extraction
- ✅ **Dual File Structure**: Perfect base-ui.com mirroring
- ✅ **Qwik Optimization**: Framework-specific handling

### **🎨 The @apply Strategy Explained**
```tsx
// 1. INPUT: Original Tailwind classes
<div className="flex flex-col items-center bg-gray-50">

// 2. INTERMEDIATE: Generated @apply CSS
.otp-root {
  @apply flex flex-col items-center bg-gray-50;
}

// 3. OUTPUT: Compiled vanilla CSS
.otp-root {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #f9fafb;
}
```

### **🎨 Core Flows**
```
Input: Tailwind Component
       ⬇️
Parse Classes → Generate Semantic Names → Create @apply CSS
       ⬇️
Compile @apply → Extract Vanilla CSS → Generate Theme Variables
       ⬇️
Output: Dual Documentation + Theme Files
```

---

## 📁 **Output Structure**
```
output/
├── tailwind/
│   ├── qwik-otp.tsx        # 📄 Original Tailwind version
│   └── theme.css           # 🎨 Shared theme variables
└── vanilla/
    ├── qwik-otp.tsx        # 📄 Semantic CSS version  
    ├── qwik-otp.css        # 🎨 Pure vanilla CSS
    └── theme.css           # 🎨 Shared theme variables
```

---

## 🎮 **Quick Start**
```bash
# Build the tool
go build -o tailwind-converter

# Convert your component
./tailwind-converter -i qwik-component.tsx -o output/ -v

# ✅ Get both Tailwind and vanilla versions automatically!
```

---
✅ **Perfect dual documentation** like base-ui.com  
✅ **Automated workflow** - no manual conversion needed  
✅ **Semantic class naming** - developer-friendly output  
✅ **Pure vanilla CSS** - zero Tailwind dependencies  
✅ **Theme consistency** - shared design tokens  

