function formatValueGerman(value: string): string {
    // Remove all non-numeric characters
    const numericString = value.replace(/[\D\s\._\-]+/g, "");
    const integerValue = numericString ? parseInt(numericString, 10) : 0;
    return integerValue === 0 ? "" : integerValue.toLocaleString();
}

/**
 * Auto-formats numeric inputs to German locale.
 *
 * @param formFields - An array containing IDs of form fields to format.
 */
function formatNumericInputs(formFields: string[]): void {
    formFields.forEach(function (field) {
        const inputElement = document.getElementById(field) as HTMLInputElement;
        const sliderElement = document.getElementById(
            `${field}Slider`,
        ) as HTMLInputElement;

        if (inputElement) {
            inputElement.addEventListener("input", function () {
                const selection = window.getSelection()?.toString();
                if (selection !== "") {
                    return;
                }
                inputElement.value = formatValueGerman(inputElement.value);
            });
        }

        if (sliderElement) {
            sliderElement.addEventListener("input", function () {
                const displayElement = document.getElementById(
                    `${field}Display`,
                ) as HTMLElement;
                if (displayElement) {
                    displayElement.textContent = formatValueGerman(
                        sliderElement.value,
                    );
                }
            });
        }
    });
}
