import React from 'react';

export default function Card({ title, description, price, image }) {
  return (
    <div class="bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300">
      <img 
        src={image} 
        alt={title}
        class="w-full h-48 object-cover"
      />
      <div class="p-6">
        <h3 class="text-xl font-bold text-gray-900 mb-2">{title}</h3>
        <p class="text-gray-600 text-sm mb-4 leading-relaxed">{description}</p>
        <div class="flex items-center justify-between">
          <span class="text-2xl font-bold text-blue-600">${price}</span>
          <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-colors">
            Add to Cart
          </button>
        </div>
      </div>
    </div>
  );
} 