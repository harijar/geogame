import { StartGameRequest, GuessGameRequest} from "./api.js";
import { GetPrompts, SavePrompts} from "./storage.js";

let prompts = await GetPrompts();

window.onload = async function() {
    await init();
    let guessButton = document.getElementById("guessButton");
    let playAgainButton = document.getElementById("playAgainButton");
    let mainMenuButton = document.getElementById("mainMenuButton");
    guessButton.onclick = async function() {
        if (prompts.length > 0) {
            await guess();
        } else {
            window.location.reload();
        }
    }
    playAgainButton.onclick = async function() {
        window.location.reload();
    }
    mainMenuButton.onclick = async function() {
        window.location.href = 'index.html'
    }
};

async function init() { // should execute on event document loaded
    if (prompts.length < 1) { // if prompts array is empty it means that game hasn't started, otherwise we continue prev game
        prompts = [(await StartGameRequest())['prompt']];
        SavePrompts(prompts);
    }
    prompts.forEach(function (prompt) {
        showPrompt(prompt);
    })
}

async function guess(){
    let guess = document.getElementById("country").value
    const data = await GuessGameRequest(guess);
    if (data === null) { // game hasn't started
        window.location.reload();
        return;
    } else if (data['right']) { // country guessed
        prompts = [];
        gameEnded();
        showCountryGuessed(data['country']);
    } else if (data['country']) { // tries limit exceeded
        prompts = [];
        gameEnded();
        showTriesExceeded(data['country']);
    } else { // wrong guess
        const prompt = data['prompt'];
        prompts.push(prompt);
        showPrompt(prompt);
    }
    await SavePrompts(prompts);
}

function showPrompt(text) {
    let prompt = document.createElement("p");
    prompt.style.textAlign = "center";
    prompt.textContent = text;
    document.getElementById("promptsDiv").append(prompt);
    let country = document.getElementById("country");
    country.value = "";
    country.autofocus = true;
}

function showCountryGuessed(country) {
    let notify = document.createElement("b");
    notify.textContent = `You guessed the country! It was ${country}.`;
    document.getElementById("promptsDiv").append(notify);
}

function showTriesExceeded(country) {
    let notify = document.createElement("b");
    notify.textContent = `The country was ${country}.`;
    document.getElementById("promptsDiv").append(notify);
}

function gameEnded() {
    document.getElementById("guessButton").style.display = 'none';
    document.getElementById("playAgainButton").style.display = 'inline';
    document.getElementById("mainMenuButton").style.display = 'inline';
}