// Synchronize slider values with input fields
function syncValues(id, isSlider) {
  const input = document.getElementById(id);
  const slider = document.getElementById(id + "Slider");
  if (isSlider) {
    input.value = slider.value;
  } else {
    slider.value = input.value;
  }
}
