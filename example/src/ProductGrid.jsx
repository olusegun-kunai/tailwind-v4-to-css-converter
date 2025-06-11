import React from 'react';
import Card from './Card';

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
    <section class="py-12 bg-gray-50">
      <div class="container mx-auto px-4">
        <div class="text-center mb-12">
          <h2 class="text-3xl font-bold text-gray-900 mb-4">Featured Products</h2>
          <p class="text-lg text-gray-600 max-w-2xl mx-auto">
            Discover our handpicked selection of premium products designed to enhance your lifestyle
          </p>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
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
        
        <div class="text-center mt-12">
          <button class="bg-gray-900 text-white px-8 py-3 rounded-lg hover:bg-gray-800 font-medium text-lg">
            View All Products
          </button>
        </div>
      </div>
    </section>
  );
} 