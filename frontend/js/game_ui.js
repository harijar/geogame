export function ShowPrompt(text) {
    let prompt = document.createElement("p")
    prompt.style.textAlign = "center"
    prompt.textContent = text
    document.getElementById("promptsDiv").append(prompt)
    document.getElementById("guessInput").value = ""
}

export function ShowCountryGuessed(country) {
    let notify = document.createElement("b")
    notify.style.textAlign = "center"
    notify.textContent = `You guessed the country! It was ${country}.`
    document.getElementById("promptsDiv").append(notify)
}

export function ShowTriesExceeded(country) {
    let notify = document.createElement("b")
    notify.textContent = `The country was ${country}.`
    notify.style.textAlign = "center"
    document.getElementById("promptsDiv").append(notify)
}

export function GameEnded() {
    document.getElementById("guessButton").style.display = 'none'
    document.getElementById("buttonsGame").style.display = 'none'
    document.getElementById("buttonsGameEnded").style.display = 'inline'
}