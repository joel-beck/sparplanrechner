"use strict";

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
