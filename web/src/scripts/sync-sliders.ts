/**
 * Synchronizes the values between number inputs and slider inputs.
 *
 * @param id - The ID of the number input.
 * @param isSlider - Whether the function is called from slider input.
 */
function syncValues(id: string, isSlider: boolean): void {
    const numberInput = document.getElementById(id) as HTMLInputElement;
    const sliderInput = document.getElementById(
        `${id}Slider`,
    ) as HTMLInputElement;

    if (isSlider && numberInput) {
        numberInput.value = sliderInput.value;
    } else if (sliderInput) {
        sliderInput.value = numberInput.value;
    }
}

/**
 * Initializes slider and number input synchronization for form fields.
 *
 * @param formFields - An array containing IDs of form fields to synchronize.
 */
export function syncSliders(formFields: string[]): void {
    formFields.forEach(function (field) {
        const numberInput = document.getElementById(field) as HTMLInputElement;
        const sliderInput = document.getElementById(
            `${field}Slider`,
        ) as HTMLInputElement;

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
