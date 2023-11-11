import { StartGameRequest, GuessGameRequest} from "./api.js";
import { GetPrompts, SavePrompts} from "./storage.js";

let prompts = await GetPrompts();

window.onload = async function() {
    await init();
};

async function init() { // should execute on event document loaded
    if (prompts.length < 1) { // if prompts array is empty it means that game hasn't started, otherwise we continue prev game
        await startGame();
    }
    prompts.forEach(function (prompt) {
        showPrompt(prompt);
    })
}

async function startGame() {
    prompts = [(await StartGameRequest())['prompt']];
    SavePrompts(prompts);
    clearPrompts();
}

async function guess(){
    let guess = document.getElementById("country").value
    const data = await GuessGameRequest(guess);
    if (data === null) { // game hasn't started
        await startGame();
        return;
    } else if (data['right']) { // country guessed
        prompts = [];
        showCountryGuessed(data['country']);
    } else if (data['country']) { // tries limit exceeded
        prompts = [];
        showTriesExceeded(data['country']);
    } else { // wrong guess
        const prompt = data['prompt'];
        prompts += prompt;
        showPrompt(prompt);
    }
    await SavePrompts(prompts);
}

function showPrompt(text) {
    document.append(`<p style="text-align: center">${text}</p>`)
}

function showCountryGuessed(country) {
    document.append(`<p style="text-align: center">You guessed the country! It was ${country}.</p>
        <button>Play again</button>
        <button>To the main page</button>`)
}

function showTriesExceeded(country) {
    document.append(`<p style="text-align: center">The country was ${country}.</p>
        <button>Play again</button>
        <button>To the main page</button>`)
}

function clearPrompts() {
    const fs = require('fs')
    fs.readFile('game.html', 'utf-8', (err, data) => {
        if (err) throw err;
        document.write(data);
    })
}