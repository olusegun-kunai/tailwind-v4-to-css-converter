import { component$ } from '@builder.io/qwik';

export const Accordion = component$(() => {
  return (
    <div class="w-full max-w-md mx-auto">
      <div class="bg-white rounded-lg shadow-md overflow-hidden">
        <button class="w-full px-6 py-4 text-left bg-gray-50 hover:bg-gray-100 focus:outline-none">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">What is Qwik?</h3>
            <svg class="w-5 h-5 text-gray-500 transform transition-transform" viewBox="0 0 20 20">
              <path d="M6 8l4 4 4-4" stroke="currentColor" strokeWidth="2" />
            </svg>
          </div>
        </button>
        <div class="px-6 py-4 text-gray-700 bg-white">
          <p class="text-sm">Qwik is a new kind of web framework that can deliver instant loading web applications.</p>
        </div>
      </div>
    </div>
  );
}); 