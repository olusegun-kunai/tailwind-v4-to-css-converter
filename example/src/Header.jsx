import React from 'react';

export default function Header() {
  return (
    <header class="bg-white shadow-md border-b border-gray-200">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-4">
            <img 
              src="/logo.png" 
              alt="Logo" 
              class="w-8 h-8 rounded-full"
            />
            <h1 class="text-2xl font-bold text-gray-900">ShopApp</h1>
          </div>
          
          <nav class="hidden md:flex items-center space-x-8">
            <a href="/" class="text-gray-700 hover:text-blue-600 font-medium">Home</a>
            <a href="/products" class="text-gray-700 hover:text-blue-600 font-medium">Products</a>
            <a href="/about" class="text-gray-700 hover:text-blue-600 font-medium">About</a>
            <a href="/contact" class="text-gray-700 hover:text-blue-600 font-medium">Contact</a>
          </nav>
          
          <div class="flex items-center space-x-4">
            <button class="relative p-2 text-gray-700 hover:text-blue-600">
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M7 4V2C7 1.45 7.45 1 8 1H16C16.55 1 17 1.45 17 2V4H20C20.55 4 21 4.45 21 5S20.55 6 20 6H19V19C19 20.1 18.1 21 17 21H7C5.9 21 5 20.1 5 19V6H4C3.45 6 3 5.55 3 5S3.45 4 4 4H7Z"/>
              </svg>
              <span class="absolute -top-1 -right-1 bg-red-500 text-white text-xs rounded-full w-5 h-5 flex items-center justify-center">3</span>
            </button>
            <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 font-medium">
              Sign In
            </button>
          </div>
        </div>
      </div>
    </header>
  );
} 