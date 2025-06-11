import { component$ } from '@builder.io/qwik';

export const OtpInput = component$(() => {
  return (
    <div className={["min-h-screen", styles.content-container].join(' ')}>
      <div className={["max-w-md shadow-md", styles.content-container].join(' ')}>
        <div className={styles.content-container}>
          <h2 className={styles.h2-text}>Enter OTP</h2>
          <p className={["mt-2", styles.p-text].join(' ')}>
            We've sent a verification code to your email
          </p>
        </div>
        
        <div className={styles.content-container}>
          <input
            type="text"
            maxLength={1}
            className={styles.input-field}
          />
          <input
            type="text"
            maxLength={1}
            className={styles.input-field}
          />
          <input
            type="text"
            maxLength={1}
            className={styles.input-field}
          />
          <input
            type="text"
            maxLength={1}
            className={styles.input-field}
          />
        </div>
        
        <button
          type="submit"
          className={["px-4 py-2", styles.button-primary].join(' ')}
        >
          Verify OTP
        </button>
        
        <div className={styles.content-container}>
          <button
            type="button"
            className={["underline", styles.button-primary].join(' ')}
          >
            Resend code
          </button>
        </div>
      </div>
    </div>
  );
}); 