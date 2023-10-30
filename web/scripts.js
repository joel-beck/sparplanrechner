"use strict";

// Define a function to sync input and range values
function syncValues(id, isSlider) {
  const numberInput = document.getElementById(id);
  const sliderInput = document.getElementById(`${id}Slider`);

  if (isSlider) {
    numberInput.value = sliderInput.value;
  } else {
    sliderInput.value = numberInput.value;
  }
}

// Attach event listeners when the DOM is fully loaded
document.addEventListener("DOMContentLoaded", function () {
  const formFields = [
    "startCapital",
    "savingsRate",
    "annualReturn",
    "years",
    "inflationRate",
    "takeoutRate",
    "tax",
  ];

  formFields.forEach(function (field) {
    // Event listener for the number input
    const numberInput = document.getElementById(field);
    numberInput.addEventListener("input", function () {
      syncValues(field, false);
    });

    // Event listener for the range slider
    const sliderInput = document.getElementById(`${field}Slider`);
    sliderInput.addEventListener("input", function () {
      syncValues(field, true);
    });
  });
});

// Initialize all tooltips on the page
var tooltipTriggerList = [].slice.call(
  document.querySelectorAll('[data-bs-toggle="tooltip"]')
);
var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
  return new bootstrap.Tooltip(tooltipTriggerEl);
});
