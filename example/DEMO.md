# 🚀 Live Demo: Tailwind to CSS Converter

## Before & After Comparison

### ❌ **BEFORE** - Messy Tailwind Classes
```jsx
// Original Card.jsx
<div class="bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300">
  <img class="w-full h-48 object-cover" />
  <div class="p-6">
    <h3 class="text-xl font-bold text-gray-900 mb-2">{title}</h3>
    <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
      Add to Cart
    </button>
  </div>
</div>
```

### ✅ **AFTER** - Clean CSS Modules
```jsx
// Converted Card.jsx
import styles from './Card.module.css';

<div class={styles.div_layout_2}>
  <img class={styles.img_layout_4} />
  <div class={styles.div_layout_2}>
    <h3 class={styles.h3_text_3}>{title}</h3>
    <button class={styles.button_1}>
      Add to Cart
    </button>
  </div>
</div>
```

```css
/* Generated Card.module.css */
.div_layout_2 {
  border-radius: 0.75rem;
  background-color: #ffffff;
  padding: 1.5rem;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
}

.button_1 {
  border-radius: 0.5rem;
  color: #ffffff;
  background-color: #2563eb;
}

.button_1:hover {
  background-color: #1d4ed8;
}
```

## 🎯 **Key Benefits Demonstrated**

1. **Cleaner JSX** - No more long class strings
2. **Semantic naming** - `button_1` instead of `bg-blue-600 text-white px-4 py-2...`
3. **Proper CSS structure** - Hover states as CSS pseudo-classes
4. **Maintainable** - Easy to modify colors, spacing, etc. in CSS
5. **No Tailwind dependency** - Pure CSS that works anywhere

## 📊 **Stats from this Demo**

| Metric | Before | After |
|--------|---------|-------|
| JSX Readability | ❌ Long class strings | ✅ Semantic names |
| CSS Dependencies | ❌ Tailwind required | ✅ Pure CSS |
| Hover States | ❌ Inline classes | ✅ Proper CSS |
| Maintainability | ❌ Hard to modify | ✅ Easy to update |

## 🛠 **How to Run This Demo**

```bash
# 1. View original files
ls example/src/

# 2. Run the converter
./tailwind-converter --input ./example/src --output ./example/dist --verbose

# 3. Compare results
diff example/src/Card.jsx example/dist/Card.jsx
```

**Perfect for showing clients, team leads, or anyone curious about migrating from Tailwind!** 