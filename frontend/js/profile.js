import {ProfileRequest} from "./api.js";

window.onload = async function() {
    document.getElementById("mainMenuButton").onclick = async function() {
        window.location.href = '/index.html';
    }
    await loadProfile()
}

async function loadProfile() {
    const data = await ProfileRequest()
    document.getElementById("title").textContent = data['name']
    document.getElementById("greetings").textContent = `Greetings, ${data['name']}!`
    document.getElementById("totalGames").textContent = data['total_games']
    document.getElementById("gamesWon").textContent = data['games_won']
}