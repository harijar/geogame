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

async function guess(text){ // should execute on guess button clicked event
    const data = await GuessGameRequest(text);
    if (data === null) { // game hasn't started
        await startGame();
        return;
    } else if (data['ok']) { // country guessed
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
    // TODO: implement, should add new prompt on page
}

function showCountryGuessed(country) {
    // TODO: implement, should show message about successfully country guess
}

function showTriesExceeded(country) {
    // TODO: implement, should show message about tries exceeded and show country
}

function clearPrompts() {
    // TODO: implement, should delete all prompts from page
}