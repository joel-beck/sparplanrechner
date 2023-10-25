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
  const formFields = ["startCapital", "savingsRate", "annualReturn", "years"];

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

let rateFieldCount = 0;
let totalDuration = 0;

function validateDuration() {
  const inputDuration = document.getElementById("years").value;
  const validationMessage = document.getElementById("validationMessage");
  const checkMark = document.getElementById("checkMark");
  if (totalDuration === parseInt(inputDuration)) {
    validationMessage.textContent = "";
    checkMark.style.display = "inline";
  } else {
    checkMark.style.display = "none";
    validationMessage.textContent = `Sum of durations should be ${inputDuration}. Currently, it is ${totalDuration}.`;
  }
}

function updateTotalDuration() {
  const durations = document.querySelectorAll(".durationField");
  totalDuration = 0;
  durations.forEach((durationField) => {
    totalDuration += parseInt(durationField.value);
  });
  validateDuration();
}

function addRateField() {
  rateFieldCount++;
  const dynamicRateFields = document.getElementById("dynamicRateFields");
  const newRateField = document.createElement("div");
  newRateField.className = "col-md-6"; // Bootstrap column class

  newRateField.innerHTML = `
    <div class="mb-3">
      <label for="savingsRate${rateFieldCount}" class="form-label">Savings Rate (€):</label>
      <input type="number" class="form-control" id="savingsRate${rateFieldCount}" name="savingsRate${rateFieldCount}" required>
      <input type="range" class="form-range" id="savingsRateSlider${rateFieldCount}" required>
    </div>
    <div class="mb-3">
      <label for="duration${rateFieldCount}" class="form-label">Duration (Years):</label>
      <input type="number" class="form-control" id="duration${rateFieldCount}" name="duration${rateFieldCount}" oninput="updateTotalDuration()" required>
      <input type="range" class="form-range" id="durationSlider${rateFieldCount}" oninput="syncValues('duration${rateFieldCount}', true)" required>
    </div>
  `;

  dynamicRateFields.appendChild(newRateField);
  updateTotalDuration();
}

function removeRateField(element) {
  const parentDiv = element.parentElement;
  parentDiv.remove();
  updateTotalDuration();
}
