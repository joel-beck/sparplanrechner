"use strict";

/**
 * Initializes a checkbox to enable or disable associated inputs.
 *
 * @param {string} checkboxId - The ID of the checkbox that will control the state of the fields.
 * @param {Array} controlledIds - An array of IDs for fields that should be controlled by the checkbox.
 */
export function initializeCheckboxToggle(checkboxId, controlledIds) {
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
export function initializeCheckboxBooleanValue(formId, checkboxIds) {
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
