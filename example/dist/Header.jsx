import React from 'react';
import styles from './Header.module.css';

export default function Header() {
  return (
    <header class={["shadow-md", styles.header_visual_7].join(' ')}>
      <div class="container mx-auto px-4 py-4">
        <div class={styles.div_layout_8}>
          <div class={styles.div_layout_8}>
            <img 
              src="/logo.png" 
              alt="Logo" 
              class={styles.img_layout_5}
            />
            <h1 class={styles.h1_text_9}>ShopApp</h1>
          </div>
          
          <nav class={styles.nav_layout_6}>
            <a href="/" class={styles.a_text_3}>Home</a>
            <a href="/products" class={styles.a_text_3}>Products</a>
            <a href="/about" class={styles.a_text_3}>About</a>
            <a href="/contact" class={styles.a_text_3}>Contact</a>
          </nav>
          
          <div class={styles.div_layout_8}>
            <button class={["relative", styles.a_text_3].join(' ')}>
              <svg class={styles.svg_layout_4} fill="currentColor" viewBox="0 0 24 24">
                <path d="M7 4V2C7 1.45 7.45 1 8 1H16C16.55 1 17 1.45 17 2V4H20C20.55 4 21 4.45 21 5S20.55 6 20 6H19V19C19 20.1 18.1 21 17 21H7C5.9 21 5 20.1 5 19V6H4C3.45 6 3 5.55 3 5S3.45 4 4 4H7Z"/>
              </svg>
              <span class={["absolute -top-1 -right-1", styles.span_container_2].join(' ')}>3</span>
            </button>
            <button class={["px-4 py-2", styles.button_1].join(' ')}>
              Sign In
            </button>
          </div>
        </div>
      </div>
    </header>
  );
} 