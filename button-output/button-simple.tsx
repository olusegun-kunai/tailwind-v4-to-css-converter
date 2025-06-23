// Simple Button with static Tailwind classes
import React from 'react';
import styles from './button-simple.module.css';


interface ButtonProps {
  children: React.ReactNode;
  onClick?: () => void;
  disabled?: boolean;
}

const Button: React.FC<ButtonProps> = ({ children, onClick, disabled = false }) => {
  return (
    <button
      className={styles.node0}
      onClick={onClick}
      disabled={disabled}
    >
      {children}
    </button>
  );
};

export default Button; 