"use strict";

// import { injectHTMLFiles } from "./compile.js";

/**
 * Validates a number that's formatted in the German numbering system, ensuring it's within a given range.
 *
 * @param {string} inputValue - The number to validate, formatted as a German number.
 * @param {number} min - The minimum acceptable value for the number.
 * @param {number} max - The maximum acceptable value for the number.
 * @returns {boolean} Whether the number is valid and within the specified range.
 */
function validateGermanNumberInRange(inputValue, min, max) {
  const germanNumberRegex = /^(\d+|\d{1,3}(\.\d{3})*),\d{2}$/;
  if (!germanNumberRegex.test(inputValue)) {
    alert(
      'Invalid format. The number must be in German format, e.g., "123.456,78"',
    );
    return false;
  }
  const standardNumber = parseFloat(
    inputValue.replace(/\./g, "").replace(",", "."),
  );
  if (standardNumber < min || standardNumber > max) {
    alert(
      `The number must be between ${min.toLocaleString(
        "de-DE",
      )} and ${max.toLocaleString("de-DE")}`,
    );
    return false;
  }
  return true;
}

/**
 * Validates form fields with German number formats.
 *
 * @param {string} formId - The ID of the form to validate.
 * @param {string} elementId - The ID of the form element to validate.
 * @param {number} min - The minimum acceptable value for the element.
 * @param {number} max - The maximum acceptable value for the element.
 */
function validateForm(formId, elementId, min, max) {
  const form = document.getElementById(formId);
  const element = document.getElementById(elementId);

  form.addEventListener("submit", function (event) {
    const inputValue = element.value;

    if (!validateGermanNumberInRange(inputValue, min, max)) {
      event.preventDefault();
    }
  });
}

/**
 * Synchronizes the values between number inputs and slider inputs.
 *
 * @param {string} id - The ID of the number input.
 * @param {boolean} isSlider - Whether the function is called from slider input.
 */
export function syncValues(id, isSlider) {
  const numberInput = document.getElementById(id);
  const sliderInput = document.getElementById(`${id}Slider`);

  if (isSlider) {
    numberInput.value = sliderInput.value;
  } else {
    sliderInput.value = numberInput.value;
  }
}

/**
 * Initializes slider and number input synchronization for form fields.
 *
 * @param {Array} formFields - An array containing IDs of form fields to synchronize.
 */
function syncSliders(formFields) {
  formFields.forEach(function (field) {
    const numberInput = document.getElementById(field);
    const sliderInput = document.getElementById(`${field}Slider`);

    if (numberInput) {
      numberInput.addEventListener("input", function () {
        syncValues(field, false);
      });
    }

    if (sliderInput) {
      sliderInput.addEventListener("input", function () {
        syncValues(field, true);
      });
    }
  });
}

/**
 * Initializes Bootstrap tooltips for elements with data-bs-toggle="tooltip" attributes.
 */
function initializeTooltips() {
  const tooltipTriggerList = [].slice.call(
    document.querySelectorAll('[data-bs-toggle="tooltip"]'),
  );
  tooltipTriggerList.forEach(function (tooltipTriggerEl) {
    new bootstrap.Tooltip(tooltipTriggerEl);
  });
}

/**
 * Initializes a checkbox to enable or disable associated inputs.
 *
 * @param {string} checkboxId - The ID of the checkbox that will control the state of the fields.
 * @param {Array} controlledIds - An array of IDs for fields that should be controlled by the checkbox.
 */
function initializeCheckboxToggle(checkboxId, controlledIds) {
  const checkbox = document.getElementById(checkboxId);
  const controlledElements = controlledIds.map((id) =>
    document.getElementById(id),
  );

  const toggleDisabledState = () => {
    const isChecked = checkbox.checked;
    controlledElements.forEach((el) => {
      el.disabled = !isChecked;
    });
  };

  toggleDisabledState();
  checkbox.addEventListener("change", toggleDisabledState);
}

/**
 * Sets the value of the checkbox to "true" or "false" based on its checked state.
 *
 * @param {HTMLInputElement} checkbox - The checkbox element.
 */
function setCheckboxBooleanValue(checkbox) {
  checkbox.value = checkbox.checked ? "true" : "false";
}

/**
 * Initializes a checkbox to have its value set to "true" or "false" upon change and
 * form submission.
 *
 * @param {string} formId - The ID of the form containing the checkbox.
 * @param {Array} checkboxIds - An array of IDs for checkboxes whose values should be
 * set to "true" or "false".
 */
function initializeCheckboxBooleanValue(formId, checkboxIds) {
  const form = document.getElementById(formId);
  const checkboxes = checkboxIds.map((id) => document.getElementById(id));

  // Initialize checkbox value based on its default checked state
  checkboxes.forEach(setCheckboxBooleanValue);

  // Update checkbox value on change and form submission
  checkboxes.forEach((checkbox) => {
    checkbox.addEventListener("change", () =>
      setCheckboxBooleanValue(checkbox),
    );
  });

  form.addEventListener("submit", () => {
    checkboxes.forEach(setCheckboxBooleanValue);
  });
}

/**
 * Auto-formats numeric inputs to German locale.
 *
 * @param {Array} formFields - An array containing IDs of form fields to format.
 */
function formatNumericInputs(formFields) {
  formFields.forEach(function (field) {
    const inputElement = document.getElementById(field);
    const sliderElement = document.getElementById(`${field}Slider`);

    const formatValue = (value) => {
      value = value.replace(/[\D\s\._\-]+/g, "");
      value = value ? parseInt(value, 10) : 0;
      return value === 0 ? "" : value.toLocaleString("de-DE");
    };

    if (inputElement) {
      inputElement.addEventListener("input", function () {
        const selection = window.getSelection().toString();
        if (selection !== "") {
          return;
        }
        inputElement.value = formatValue(inputElement.value);
      });
    }

    if (sliderElement) {
      sliderElement.addEventListener("input", function () {
        const displayElement = document.getElementById(`${field}Display`);
        if (displayElement) {
          displayElement.textContent = formatValue(sliderElement.value);
        }
      });
    }
  });
}

function prefersSystemDarkmode() {
  return window.matchMedia("(prefers-color-scheme: dark)").matches;
}

function showLightIcon(themeToggle, lightIcon) {
  themeToggle.classList.add("bg-yellow-500", "-translate-x-2");
  themeToggle.classList.remove("bg-gray-700", "translate-x-2");

  setTimeout(() => {
    themeToggle.innerHTML = lightIcon;
  }, 250);
}

function showDarkIcon(themeToggle, darkIcon) {
  themeToggle.classList.add("bg-gray-700", "translate-x-2");
  themeToggle.classList.remove("bg-yellow-500", "-translate-x-2");

  setTimeout(() => {
    themeToggle.innerHTML = darkIcon;
  }, 250);
}

function setInitialTheme(themeToggle, darkIcon, lightIcon) {
  if (prefersSystemDarkmode()) {
    localStorage.setItem("isDarkmode", true);
    showDarkIcon(themeToggle, darkIcon);
  } else {
    localStorage.setItem("isDarkmode", false);
    showLightIcon(themeToggle, lightIcon);
  }
}

function isDarkmode() {
  return localStorage.getItem("isDarkmode") === "true";
}

function toggleTheme(themeToggle, darkIcon, lightIcon) {
  if (isDarkmode()) {
    localStorage.setItem("isDarkmode", false);
    showLightIcon(themeToggle, lightIcon);
    return;
  }

  localStorage.setItem("isDarkmode", true);
  showDarkIcon(themeToggle, darkIcon);
}

/**
 * Attaches event listeners when the DOM is fully loaded.
 */
function main() {
  document.addEventListener("DOMContentLoaded", async function () {
    const sliderFields = [
      "savingsRate",
      "years",
      "annualReturnRate",
      "inflationRate",
      "takeoutRate",
      "tax",
    ];
    // const formFields = sliderFields.concat(["startCapital"]);

    // Inject HTML content into the page
    // await injectHTMLFiles();

    syncSliders(sliderFields);
    initializeTooltips();

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

    const darkIcon = `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
</svg>`;

    const lightIcon = `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
</svg>`;

    const themeToggle = document.querySelector("#themeToggle");

    setInitialTheme(themeToggle, darkIcon, lightIcon);

    // TODO: Button moves but theme does not actually change on click
    // TODO: Button Movement is not yet symmetric from light to dark mode
    themeToggle.addEventListener("click", function () {
      toggleTheme(themeToggle, darkIcon, lightIcon);
    });

    // formatNumericInputs(formFields);
  });
}

main();
