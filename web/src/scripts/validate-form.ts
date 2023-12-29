/**
 * Validates a number that's formatted in the German numbering system, ensuring it's within a given range.
 *
 * @param inputValue - The number to validate, formatted as a German number.
 * @param min - The minimum acceptable value for the number.
 * @param max - The maximum acceptable value for the number.
 * @returns Whether the number is valid and within the specified range.
 */
function validateGermanNumberInRange(
    inputValue: string,
    min: number,
    max: number,
): boolean {
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
 * @param formId - The ID of the form to validate.
 * @param elementId - The ID of the form element to validate.
 * @param min - The minimum acceptable value for the element.
 * @param max - The maximum acceptable value for the element.
 */
function validateForm(
    formId: string,
    elementId: string,
    min: number,
    max: number,
): void {
    const form = document.getElementById(formId) as HTMLFormElement;
    const element = document.getElementById(elementId) as HTMLInputElement;

    form.addEventListener("submit", function (event: Event) {
        const inputValue = element.value;

        if (!validateGermanNumberInRange(inputValue, min, max)) {
            event.preventDefault();
        }
    });
}
