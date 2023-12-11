import {StartGameRequest, GuessGameRequest} from "./api.js";
import {GetPrompts, SavePrompts} from "./storage.js";
import {GameEnded, ShowPrompt, ShowTriesExceeded, ShowCountryGuessed} from "./game_ui.js";
import {Auth, CheckAuth} from "./auth.js";

let prompts;

window.onload = async function() {
    prompts = await GetPrompts();
    await ContinueOrStartGame();
    prompts.forEach(function (prompt) {
        ShowPrompt(prompt);
    })

    document.getElementById("guessForm").onsubmit = async function(event) {
        event.preventDefault();
        const ok = await Guess(document.getElementById("guessInput").value);
        if (!ok) {
            window.location.reload();
        }
    };

    document.getElementById("playAgainButton").onclick = async function() {
        window.location.reload();
    }
    document.getElementById("mainMenuButton").onclick = async function() {
        window.location.href = 'index.html';
    }
    document.getElementById("newGameButton").onclick = async function() {
        SavePrompts([]);
        window.location.reload();
    }
    document.getElementById("stopGameButton").onclick = async function() {
        SavePrompts([]);
        window.location.href = 'index.html';
    }

    window.auth = auth;
    await CheckAuth();
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
        ShowCountryGuessed(data['country']);
        GameEnded();
    } else if (data['country']) { // tries limit exceeded
        prompts = [];
        ShowTriesExceeded(data['country']);
        GameEnded();
    } else if (data['prompt']) { // wrong guess
        const prompt = data['prompt'];
        prompts.push(prompt);
        ShowPrompt(prompt);
    }
    await SavePrompts(prompts);
    return true
}

async function auth(user) {
    await Auth(user);
}