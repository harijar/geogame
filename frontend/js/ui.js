let promptsDiv;
let guessInput;

window.onload = async function() {
    promptsDiv = document.getElementById("promptsDiv");
    guessInput = document.getElementById("guessInput");
}

export function ShowPrompt(text) {
    let prompt = document.createElement("p");
    prompt.style.textAlign = "center";
    prompt.textContent = text;
    promptsDiv.append(prompt);

    guessInput.value = "";
}

export function ShowCountryGuessed(country) {
    let notify = document.createElement("b");
    notify.textContent = `You guessed the country! It was ${country}.`;
    promptsDiv.append(notify);
}

export function ShowTriesExceeded(country) {
    let notify = document.createElement("b");
    notify.textContent = `The country was ${country}.`;
    promptsDiv.append(notify);
}

export function GameEnded() {
    document.getElementById("guessButton").style.display = 'none';
    document.getElementById("promptsDiv").style.display = 'inline';
}