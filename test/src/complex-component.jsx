import React from 'react';

export default function ComplexComponent() {
  return (
    <div class="container mx-auto p-6 bg-white shadow-lg rounded-lg">
      <header class="border-b pb-4 mb-6">
        <h1 class="text-3xl font-bold text-gray-900">Dashboard</h1>
        <p class="text-gray-600 mt-2">Welcome to your admin panel</p>
      </header>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div class="bg-blue-50 p-4 rounded-md border border-blue-200">
          <h2 class="text-lg font-semibold text-blue-900">Statistics</h2>
          <span class="text-2xl font-bold text-blue-600">1,234</span>
        </div>
        
        <div class="bg-green-50 p-4 rounded-md border border-green-200">
          <h2 class="text-lg font-semibold text-green-900">Revenue</h2>
          <span class="text-2xl font-bold text-green-600">$12,345</span>
        </div>
        
        <div class="bg-purple-50 p-4 rounded-md border border-purple-200 hover:shadow-md transition-shadow">
          <h2 class="text-lg font-semibold text-purple-900">Users</h2>
          <span class="text-2xl font-bold text-purple-600">567</span>
        </div>
      </div>
      
      <button class="mt-6 px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
        View Details
      </button>
    </div>
  );
} 