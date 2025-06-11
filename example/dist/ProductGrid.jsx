import React from 'react';
import Card from './Card';
import styles from './ProductGrid.module.css';

export default function ProductGrid() {
  const products = [
    {
      id: 1,
      title: "Wireless Headphones",
      description: "High-quality sound with noise cancellation",
      price: 199,
      image: "/headphones.jpg"
    },
    {
      id: 2,
      title: "Smart Watch",
      description: "Track your fitness and stay connected",
      price: 299,
      image: "/watch.jpg"
    },
    {
      id: 3,
      title: "Laptop Stand",
      description: "Ergonomic design for better posture",
      price: 49,
      image: "/stand.jpg"
    }
  ];

  return (
    <section class={["py-12", styles.section_visual_1].join(' ')}>
      <div class="container mx-auto px-4">
        <div class={["mb-12", styles.div_container_4].join(' ')}>
          <h2 class={["mb-4", styles.h2_text_3].join(' ')}>Featured Products</h2>
          <p class={["max-w-2xl mx-auto", styles.p_text_5].join(' ')}>
            Discover our handpicked selection of premium products designed to enhance your lifestyle
          </p>
        </div>
        
        <div class={styles.div_container_4}>
          {products.map(product => (
            <Card 
              key={product.id}
              title={product.title}
              description={product.description}
              price={product.price}
              image={product.image}
            />
          ))}
        </div>
        
        <div class={["mt-12", styles.div_container_4].join(' ')}>
          <button class={["px-8 py-3", styles.button_2].join(' ')}>
            View All Products
          </button>
        </div>
      </div>
    </section>
  );
} 