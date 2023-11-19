"use strict";

/**
 * Synchronizes the values between number inputs and slider inputs.
 *
 * @param {string} id - The ID of the number input.
 * @param {boolean} isSlider - Whether the function is called from slider input.
 */
function syncValues(id, isSlider) {
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
export function syncSliders(formFields) {
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
