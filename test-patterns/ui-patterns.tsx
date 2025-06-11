import { component$ } from '@builder.io/qwik';
import styles from './ui-patterns.module.css';

export const UIPatterns = component$(() => {
  return (
    <div>
      {/* Hero Section Pattern */}
      <div className={["min-h-screen from-blue-500 to-purple-600", styles.grid-container].join(' ')}>
        <h1 className={styles.h1-text}>Hero Section</h1>
      </div>

      {/* Card Pattern */}
      <div className={["shadow-lg max-w-sm", styles.grid-container].join(' ')}>
        <h3 className={styles.h3-text}>Card Title</h3>
        <p className={styles.p-text}>Card content</p>
      </div>

      {/* Modal Overlay Pattern */}
      <div className={["fixed inset-0", styles.grid-container].join(' ')}>
        <div className={styles.grid-container}>Modal Content</div>
      </div>

      {/* Dropdown Menu Pattern */}
      <div className={["absolute top-10 right-0 shadow-lg py-2", styles.grid-container].join(' ')}>
        <a href="#" className={["px-4 py-2", styles.a-interactive].join(' ')}>Menu Item</a>
      </div>

      {/* Badge Pattern */}
      <span className={["px-2 py-1", styles.span-text].join(' ')}>Badge</span>

      {/* Grid Pattern */}
      <div className={styles.grid-container}>
        <div>Item 1</div>
        <div>Item 2</div>
        <div>Item 3</div>
      </div>

      {/* Accordion Trigger Pattern */}
      <button className={["px-4 py-2", styles.accordion-trigger].join(' ')}>
        Accordion Button
      </button>

      {/* Navigation Pattern */}
      <nav className={styles.grid-container}>
        <a href="#" className={styles.a-interactive}>Home</a>
        <a href="#" className={styles.a-interactive}>About</a>
      </nav>

      {/* Toast Notification Pattern */}
      <div className={["fixed top-4 right-4 px-4 py-2", styles.grid-container].join(' ')}>
        Success message
      </div>
    </div>
  );
}); 