import {
    initializeCheckboxBooleanValue,
    initializeCheckboxToggle,
} from "../src/scripts/checkboxes";
import { setInitialTheme, toggleTheme } from "../src/scripts/dark-mode";
import { syncSliders } from "../src/scripts/sync-sliders";
import { collapseTable } from "../src/scripts/toggle-dropdown";
import { hideTooltip, showTooltip } from "../src/scripts/tooltips";

/**
 * Initializes the theme of the application upon startup.
 * Sets the initial theme based on the user's preference or default settings.
 *
 * @param {string} darkIcon - SVG markup for the dark mode icon.
 * @param {string} lightIcon - SVG markup for the light mode icon.
 * @param {HTMLElement} themeToggleIcon - The HTML element representing the theme toggle icon.
 */
function runStartupContent(
    darkIcon: string,
    lightIcon: string,
    themeToggleIcon: HTMLElement,
): void {
    setInitialTheme(themeToggleIcon, darkIcon, lightIcon);
}

/**
 * Initializes dynamic content and functionality after an htmx swap.
 * Sets up sliders, checkbox toggles, and dropdowns for the application.
 */
function runAfterSwap(): void {
    const sliderFields: string[] = [
        "savingsRate",
        "years",
        "annualReturnRate",
        "inflationRate",
        "takeoutRate",
        "tax",
    ];
    syncSliders(sliderFields);

    initializeCheckboxToggle("inflationRateCheckbox", [
        "inflationRate",
        "inflationRateSlider",
    ]);
    initializeCheckboxToggle("taxCheckbox", ["tax", "taxSlider"]);
    initializeCheckboxToggle("takeoutRateCheckbox", [
        "takeoutRate",
        "takeoutRateSlider",
    ]);

    initializeCheckboxBooleanValue("form", [
        "inflationRateCheckbox",
        "takeoutRateCheckbox",
        "taxCheckbox",
    ]);

    const dropdowns = document.querySelectorAll(".dropdown");
    dropdowns.forEach(function (dropdown) {
        const button = dropdown.querySelector(".button") as HTMLElement;
        const icon = button.querySelector(".dropdown-icon") as HTMLElement;
        const tableContainer = dropdown.querySelector(
            ".table-container",
        ) as HTMLElement;

        button.addEventListener("click", () => {
            collapseTable(tableContainer, icon);
        });
    });

    const tooltips = document.querySelectorAll(".tooltip");
    tooltips.forEach(function (tooltip) {
        tooltip.addEventListener("mouseenter", () => {
            showTooltip(tooltip as HTMLElement);
        });

        tooltip.addEventListener("mouseleave", () => {
            hideTooltip(tooltip as HTMLElement);
        });
    });
}

/**
 * Sets up theme toggle functionality after page load or htmx reload.
 * Attaches an event listener to the theme toggle button for toggling the theme.
 *
 * @param {string} darkIcon - SVG markup for the dark mode icon.
 * @param {string} lightIcon - SVG markup for the light mode icon.
 * @param {HTMLElement} themeToggleIcon - The HTML element representing the theme toggle icon.
 */
function runAfterOnLoad(
    darkIcon: string,
    lightIcon: string,
    themeToggleIcon: HTMLElement,
): void {
    const themeToggleButton = document.querySelector(
        "#themeToggleButton",
    ) as HTMLElement;
    themeToggleButton.addEventListener("click", () => {
        toggleTheme(themeToggleIcon, darkIcon, lightIcon);
    });
}

/**
 * Main function to initialize the application.
 * Attaches event listeners for page load and htmx page reload to ensure that the dynamic content is executed at the appropriate times.
 */
function main(): void {
    const darkIcon: string = `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
    </svg>`;

    const lightIcon: string = `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
    </svg>`;

    const themeToggleIcon = document.querySelector(
        "#themeToggleIcon",
    ) as HTMLElement;

    document.addEventListener("DOMContentLoaded", () => {
        runStartupContent(darkIcon, lightIcon, themeToggleIcon);
        runAfterSwap();
        runAfterOnLoad(darkIcon, lightIcon, themeToggleIcon);
    });

    document.addEventListener("htmx:AfterOnLoad", () => {
        runAfterOnLoad(darkIcon, lightIcon, themeToggleIcon);
    });

    document.addEventListener("htmx:afterSwap", () => {
        runAfterSwap();
    });
}

main();
