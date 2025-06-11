# 🎉 Tailwind v4+ to CSS Dual Documentation Generator

## ✅ **COMPLETED: Perfect for qwik.design Documentation!**

This tool transforms **1 Qwik component** into **3 files** for dual documentation (like base-ui.com):

### 📊 **Input → Output Example**

**Input:** `qwik-otp.tsx` (with Tailwind classes)
```tsx
<button className="w-full px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700">
  Verify OTP
</button>
```

**Output:** 3 files generated automatically

#### 1️⃣ **Tailwind Version** (`tailwind/qwik-otp.tsx`)
```tsx
<button className="w-full px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700">
  Verify OTP
</button>
```
*→ Unchanged original for Tailwind users*

#### 2️⃣ **Vanilla CSS Version** (`vanilla/qwik-otp.tsx`)
```tsx
<button className={["px-4 py-2", styles.button-primary].join(' ')}>
  Verify OTP
</button>
```
*→ Semantic classes for vanilla CSS users*

#### 3️⃣ **CSS File** (`vanilla/input.css`)
```css
@tailwind base;
@tailwind components;
@tailwind utilities;

.button-primary {
  @apply bg-blue-600 hover:bg-blue-700 text-white w-full rounded-md;
}
```
*→ Ready for Tailwind compilation!*

## 🚀 **Usage**

```bash
# Generate dual documentation
go run . -i your-component.tsx -o output-dir -v

# Compile the CSS (optional)
npx tailwindcss -i output-dir/vanilla/input.css -o output-dir/vanilla/compiled.css
```

## 🎯 **Advanced Features**

### **Smart Semantic Naming**
- ✅ `accordion-trigger` - Detects accordion buttons
- ✅ `modal-overlay` - Detects modal backgrounds  
- ✅ `navigation-container` - Detects nav elements
- ✅ `grid-container` - Detects grid layouts
- ✅ `card-container` - Detects card patterns
- ✅ `hero-section` - Detects hero sections

### **Perfect for Documentation**
- ✅ **Base UI style** - Shows both Tailwind and vanilla options
- ✅ **Accurate CSS** - Uses @apply for perfect compilation
- ✅ **Semantic names** - Meaningful class names
- ✅ **Ready to compile** - Complete Tailwind workflow

## 📁 **Generated Structure**

```
output/
├── tailwind/
│   └── component.tsx         # Original Tailwind version
├── vanilla/
│   ├── component.tsx         # Semantic CSS version
│   ├── component.css         # @apply styles
│   └── input.css            # Full compilation-ready CSS
└── component.module.css      # Legacy module format
```

## ✨ **Real Example Output**

Generated from OTP component with enhanced patterns:

### Semantic Classes Generated:
- `button-primary` - Submit button with hover states
- `input-field` - OTP input fields  
- `navigation-container` - Layout container
- `h2-text` - Heading typography
- `p-text` - Paragraph text

### Ready for Production:
1. **Documentation Website** ✅
2. **Tailwind Compilation** ✅  
3. **Semantic CSS** ✅
4. **Automated Generation** ✅

**Perfect for qwik.design dual documentation! 🎉** 