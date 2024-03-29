import daisyui from "daisyui";

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
    "node_modules/react-daisyui/dist/**/*.js"
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['Monserrat', 'Arial', 'sans-serif'],
        "slogan": "Race Sport"
      }
    }
  },
  plugins: [
      daisyui,
  ],
  daisyui: {
    themes: [
      {
        corporate: {
          ...require("daisyui/src/theming/themes")["[data-theme=corporate]"],
          primary: "#F38944",
          secondary: "#476a6f",
          "base-100": "#ffff",
          "base-200": "#eff3f6",
          "--rounded-box": "0.3rem", // border radius rounded-box utility class, used in card and other large boxes
          "--rounded-btn": "0.3rem", // border radius rounded-btn utility class, used in buttons and similar element
          "--rounded-badge": "1rem", // border radius rounded-badge utility class, used in badges and similar
          "--animation-btn": "0.25s", // duration of animation when you click on button
          "--animation-input": "0.2s", // duration of animation for inputs like checkbox, toggle, radio, etc
          "--btn-text-case": "uppercase", // set default text transform for buttons
          "--btn-focus-scale": "0.95", // scale transform of button when you focus on it
          "--border-btn": "1px", // border width of buttons
          "--tab-border": "0px", // border width of tabs
          "--tab-radius": "0.5rem", // border radius of tabs
        },
      },
    ],
  },
};