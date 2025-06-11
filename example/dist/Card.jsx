import React from 'react';
import styles from './Card.module.css';

export default function Card({ title, description, price, image }) {
  return (
    <div class={["shadow-lg overflow-hidden transition-shadow duration-300", styles.div_layout_2].join(' ')}>
      <img 
        src={image} 
        alt={title}
        class={["object-cover", styles.img_layout_4].join(' ')}
      />
      <div class={styles.div_layout_2}>
        <h3 class={["mb-2", styles.h3_text_3].join(' ')}>{title}</h3>
        <p class={["mb-4", styles.p_text_5].join(' ')}>{description}</p>
        <div class={styles.div_layout_2}>
          <span class={styles.span_text_6}>${price}</span>
          <button class={["px-4 py-2 transition-colors", styles.button_1].join(' ')}>
            Add to Cart
          </button>
        </div>
      </div>
    </div>
  );
} 