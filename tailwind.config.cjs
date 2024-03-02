/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: "class",
    content: ["./internal/**/*.templ"],
    theme: {
        extend: {
            colors: {
                "color-text-light": "#313131", // Text color in light mode
                "color-bg-light": "#F8FAFC", // Background color in light mode
                "color-header-light": "#006AE4", // Header + Button background color in light mode
                "color-accent-light": "#3089E4", // Accent color in light mode
                "color-hover-light": "#0056CC", // Hover color in light mode
                "color-form-light": "#EFEFEF", // Form field background in light mode
                "color-text-dark": "#D3D3D3", // Text color in dark mode
                "color-bg-dark": "#112534", // Background color in dark mode
                "color-header-dark": "#DB8910", // Header + Button background color in dark mode
                "color-accent-dark": "#E38F1A", // Accent color in dark mode
                "color-hover-dark": "#C07941", // Hover color in dark mode
                "color-form-dark": "#E5E4E2", // Form field background in dark mode
            },
            // fontFamily: {
            //     body: ["Nunito"],
            // },
        },
    },
    plugins: [
        require("@tailwindcss/typography"),
        require("@tailwindcss/forms"),
        require("@tailwindcss/aspect-ratio"),
        require("@tailwindcss/container-queries"),
    ],
};
