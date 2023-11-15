/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/src/**/*.{html,js}"],
  theme: {
    extend: {},
    fontFamily: {
      sans: ["Lato", "Montserrat", "Roboto", "sans-serif"],
      serif: ["Merriweather", "serif"],
      mono: ['"Roboto Mono"', "monospace"],
    },
  },
  plugins: [
    require("@tailwindcss/typography"),
    require("@tailwindcss/forms"),
    require("@tailwindcss/aspect-ratio"),
    require("@tailwindcss/container-queries"),
  ],
};
