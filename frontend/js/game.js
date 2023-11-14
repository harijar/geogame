import { StartGameRequest, GuessGameRequest} from "./api.js";
import { GetPrompts, SavePrompts} from "./storage.js";
import {GameEnded, ShowPrompt, ShowTriesExceeded, ShowCountryGuessed} from "./ui";

let prompts;

window.onload = async function() {
    prompts = await GetPrompts();
    await ContinueOrStartGame();
    prompts.forEach(function (prompt) {
        ShowPrompt(prompt);
    })

    document.getElementById("guessButton").onclick = async function(e) {
        const ok = await Guess(e.target.innerText);
        if (!ok) {
            window.location.reload();
        }
    }
    document.getElementById("playAgainButton").onclick = async function() {
        window.location.reload();
    }
    document.getElementById("mainMenuButton").onclick = async function() {
        window.location.href = 'index.html'
    }
};

async function ContinueOrStartGame() {
    if (prompts.length < 1) { // if prompts array is empty it means that game hasn't started, otherwise we continue prev game
        prompts = [(await StartGameRequest())['prompt']];
        SavePrompts(prompts);
    }
}

async function Guess(prompt){
    const data = await GuessGameRequest(prompt);
    if (data === null) { // game hasn't started
        prompts = [];
        await SavePrompts(prompts);
        return false;
    } else if (data['right']) { // country guessed
        prompts = [];
        GameEnded();
        ShowCountryGuessed(data['country']);
    } else if (data['country']) { // tries limit exceeded
        prompts = [];
        GameEnded();
        ShowTriesExceeded(data['country']);
    } else if (data['prompt']) { // wrong guess
        const prompt = data['prompt'];
        prompts.push(prompt);
        ShowPrompt(prompt);
    }
    await SavePrompts(prompts);
    return true;
}