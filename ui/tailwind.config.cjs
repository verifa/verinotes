/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/routes/**/*.{svelte,js,ts}'],
  theme: { // in one example theme is not present, in another it is
    extend: {},
  },
  plugins: [
    require("@tailwindcss/typography"),
    require('daisyui')
  ],
};
