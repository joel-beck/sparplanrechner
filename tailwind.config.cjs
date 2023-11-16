/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/**/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        "color-text-light": "#A9A9A9",
        "color-text-dark": "#D3D3D3",
        "color-bg-light": "#F8FAFC",
        "color-bg-dark": "#203447",
        "color-accent-light": "#F56565",
        "color-accent-dark": "#ED8936",
        "color-form-light": "#A9A9A9",
        "color-form-dark": "#E5E4E2",
      },
    },
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
