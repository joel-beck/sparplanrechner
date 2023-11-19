"use strict";

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
